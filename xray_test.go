package gobenchtraces

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/aws/aws-xray-sdk-go/xraylog"
	newrelic "github.com/newrelic/go-agent"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/require"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"golang.org/x/sync/errgroup"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type service interface {
	Setup() error
	Start() error
	Stop() error
}

type multiService struct {
	services []service
}

func (m *multiService) Setup() error {
	for _, e := range m.services {
		if err := e.Setup(); err != nil {
			return err
		}
	}
	return nil
}

func (m *multiService) Start() error {
	ctx := context.Background()
	eg, _ := errgroup.WithContext(ctx)
	for _, e := range m.services {
		eg.Go(e.Start)
	}
	return eg.Wait()
}

func (m *multiService) Stop() error {
	for _, e := range m.services {
		if err := e.Stop(); err != nil {
			return err
		}
	}
	return nil
}

type xrayTrial struct {
}

func (m *xrayTrial) Setup() error {
	xray.SetLogger(xraylog.NullLogger)
	return nil
}

func (m *xrayTrial) Start() error {
	return nil
}

func (m *xrayTrial) Stop() error {
	return nil
}

type ddTrial struct {
	server http.Server
}

func (m *ddTrial) Setup() error {
	log.SetOutput(ioutil.Discard) // Yes, the DD logger does this :(
	tracer.Start(tracer.WithServiceName("test"))
	return nil
}

func (m *ddTrial) Start() error {
	m.server.Addr = ":8126"
	m.server.Handler = m
	err := m.server.ListenAndServe()
	if err == nil || err.Error() == http.ErrServerClosed.Error() {
		return nil
	}
	return err
}

func (m *ddTrial) ServeHTTP(rw http.ResponseWriter, _ *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

func (m *ddTrial) Stop() error {
	return m.server.Close()
}

type jaegerTrial struct {
	onClose io.Closer
	tracer  opentracing.Tracer
}

func (m *jaegerTrial) Setup() error {
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		return err
	}
	cfg.ServiceName = "test"
	openTracer, closer, err := cfg.NewTracer(jaegercfg.Sampler(jaeger.NewConstSampler(true)))
	if err != nil {
		return err
	}
	m.tracer = openTracer
	m.onClose = closer
	opentracing.SetGlobalTracer(openTracer)
	return nil
}

func (m *jaegerTrial) Start() error {
	return nil
}

func (m *jaegerTrial) Stop() error {
	if m.onClose != nil {
		return m.onClose.Close()
	}
	return nil
}

type nrTrial struct {
	app newrelic.Application
}

func (m *nrTrial) Setup() error {
	config := newrelic.NewConfig("Your Application Name", "0123456789012345678901234567890123456789")
	config.Logger = newrelic.NewLogger(ioutil.Discard)
	app, err := newrelic.NewApplication(config)
	if err != nil {
		return err
	}
	m.app = app

	return nil
}

func (m *nrTrial) Start() error {
	return nil
}

func (m *nrTrial) Stop() error {
	m.app.Shutdown(time.Second)
	return nil
}

var _ http.HandlerFunc = emptyHandler

func emptyHandler(rw http.ResponseWriter, req *http.Request) {
}

func emptyTestFunc() {

}

type runBag struct {
	nrApp        newrelic.Application
	jaegerTracer opentracing.Tracer
}

type goroutineBag struct {
	rw  *httptest.ResponseRecorder
	req *http.Request
}

func (g *goroutineBag) setup() {
	var err error
	g.req, err = http.NewRequest("POST", "www.example.com", nil)
	if err != nil {
		panic(err)
	}
	g.rw = httptest.NewRecorder()
}

type benchmarkTracesRun struct {
	name string

	atOnce int
	trace  func(b *testing.B, run benchmarkTracesRun)
	toCall func()

	runBag       runBag
	goroutineBag goroutineBag
}

func openjaegerRun(b *testing.B, run benchmarkTracesRun) {
	ctx := context.Background()
	span, _ := opentracing.StartSpanFromContextWithTracer(ctx, run.runBag.jaegerTracer, "operation_name")
	run.toCall()
	span.Finish()
}

func xrayRun(b *testing.B, run benchmarkTracesRun) {
	ctx := context.Background()
	_, s := xray.BeginSegment(ctx, "start")
	run.toCall()
	s.Close(nil)
}

func ddRun(b *testing.B, run benchmarkTracesRun) {
	span := tracer.StartSpan("test")
	run.toCall()
	span.Finish(tracer.WithError(nil))
}

func newRelicRun(b *testing.B, run benchmarkTracesRun) {
	txn := run.runBag.nrApp.StartTransaction("myTxn", run.goroutineBag.rw, run.goroutineBag.req)
	require.NoError(b, txn.End())
}

func BenchmarkTraces(b *testing.B) {
	b.ReportAllocs()
	nrT := nrTrial{}
	jt := &jaegerTrial{}
	ms := multiService{
		services: []service{
			&ddTrial{}, jt, &nrT, &xrayTrial{},
		},
	}
	require.NoError(b, ms.Setup())
	go func() {
		require.NoError(b, ms.Start())
	}()
	const upper = 1000
	runs := []benchmarkTracesRun{
		{
			name:   "x-ray",
			atOnce: 1,
			trace:  xrayRun,
			toCall: emptyTestFunc,
		},
		{
			name:   "x-ray",
			atOnce: upper,
			trace:  xrayRun,
			toCall: emptyTestFunc,
		},
		{
			name:   "datadog",
			atOnce: 1,
			trace:  ddRun,
			toCall: emptyTestFunc,
		},
		{
			name:   "datadog",
			atOnce: upper,
			trace:  ddRun,
			toCall: emptyTestFunc,
		},
		{
			name:   "openjaeger",
			atOnce: 1,
			trace:  openjaegerRun,
			toCall: emptyTestFunc,
		},
		{
			name:   "openjaeger",
			atOnce: upper,
			trace:  openjaegerRun,
			toCall: emptyTestFunc,
		},
		{
			name:   "newrelic",
			atOnce: 1,
			trace:  newRelicRun,
			toCall: emptyTestFunc,
		},
		{
			name:   "newrelic",
			atOnce: upper,
			trace:  newRelicRun,
			toCall: emptyTestFunc,
		},
	}
	for _, run := range runs {
		b.Run(fmt.Sprintf("%s-%d", run.name, run.atOnce), func(b *testing.B) {
			run := run
			run.runBag.nrApp = nrT.app
			run.runBag.jaegerTracer = jt.tracer
			wg := sync.WaitGroup{}
			for ao := 0; ao < run.atOnce; ao++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					run2 := run
					run2.goroutineBag.setup()
					for i := 0; i < b.N/run2.atOnce; i++ {
						run2.trace(b, run2)
					}
				}()
			}
			wg.Wait()
		})
	}
	require.NoError(b, ms.Stop())
}
