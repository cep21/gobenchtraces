```go

(pprof) Showing nodes accounting for 1250382, 69.65% of 1795323 total
Dropped 30 nodes (cum <= 8976)
Showing top 10 nodes out of 60
      flat  flat%   sum%        cum   cum%
    222572 12.40% 12.40%     725024 40.38%  github.com/aws/aws-xray-sdk-go/xray.basicSegment
    199903 11.13% 23.53%     562200 31.31%  github.com/aws/aws-xray-sdk-go/xray.(*DefaultEmitter).Emit
    196611 10.95% 34.48%     327685 18.25%  github.com/aws/aws-xray-sdk-go/xray.NewTraceID
    174770  9.73% 44.22%     174770  9.73%  github.com/aws/aws-xray-sdk-go/internal/logger.Debugf
     98305  5.48% 49.69%      98305  5.48%  fmt.(*pp).printValue
     81921  4.56% 54.26%     362297 20.18%  github.com/aws/aws-xray-sdk-go/xray.packSegments
     78300  4.36% 58.62%     280376 15.62%  encoding/json.Marshal
     66925  3.73% 62.35%      66925  3.73%  runtime.malg
     65538  3.65% 66.00%     202076 11.26%  encoding/json.mapEncoder.encode
     65537  3.65% 69.65%     163842  9.13%  fmt.Sprintf
(pprof) 
(pprof) Showing nodes accounting for 733775, 40.87% of 1795323 total
Dropped 30 nodes (cum <= 8976)
Showing top 10 nodes out of 60
      flat  flat%   sum%        cum   cum%
         0     0%     0%    1655437 92.21%  github.com/cep21/gobenchtraces.BenchmarkTraces.func2.1
         0     0%     0%    1640417 91.37%  github.com/cep21/gobenchtraces.xrayRun
         0     0%     0%     947141 52.76%  github.com/aws/aws-xray-sdk-go/xray.BeginSegment
    222572 12.40% 12.40%     725024 40.38%  github.com/aws/aws-xray-sdk-go/xray.basicSegment
     32768  1.83% 14.22%     693276 38.62%  github.com/aws/aws-xray-sdk-go/xray.(*Segment).Close
    199903 11.13% 25.36%     562200 31.31%  github.com/aws/aws-xray-sdk-go/xray.(*DefaultEmitter).Emit
         0     0% 25.36%     562200 31.31%  github.com/aws/aws-xray-sdk-go/xray.(*Segment).emit
         0     0% 25.36%     562200 31.31%  github.com/aws/aws-xray-sdk-go/xray.(*Segment).flush
     81921  4.56% 29.92%     362297 20.18%  github.com/aws/aws-xray-sdk-go/xray.packSegments
    196611 10.95% 40.87%     327685 18.25%  github.com/aws/aws-xray-sdk-go/xray.NewTraceID
(pprof) 
(pprof) Total: 1795323
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/internal/logger.Debugf in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/internal/logger/logger.go
    174770     174770 (flat, cum)  9.73% of Total
         .          .     19:
         .          .     20:// The Logger instance used by xray to log. Set via xray.SetLogger().
         .          .     21:var Logger xraylog.Logger = xraylog.NewDefaultLogger(os.Stdout, xraylog.LogLevelInfo)
         .          .     22:
         .          .     23:func Debugf(format string, args ...interface{}) {
    174770     174770     24:	Logger.Log(xraylog.LogLevelDebug, printfArgs{format, args})
         .          .     25:}
         .          .     26:
         .          .     27:func Debug(args ...interface{}) {
         .          .     28:	Logger.Log(xraylog.LogLevelDebug, printArgs(args))
         .          .     29:}
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*DefaultEmitter).Emit in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/default_emitter.go
    199903     562200 (flat, cum) 31.31% of Total
         .          .     59:func (de *DefaultEmitter) Emit(seg *Segment) {
         .          .     60:	if seg == nil || !seg.ParentSegment.Sampled {
         .          .     61:		return
         .          .     62:	}
         .          .     63:
     81922     444219     64:	for _, p := range packSegments(seg, nil) {
         .          .     65:		// defer expensive marshal until log message is actually logged
     65537      65537     66:		logger.DebugDeferred(func() string {
         .          .     67:			var b bytes.Buffer
         .          .     68:			json.Indent(&b, p, "", " ")
         .          .     69:			return b.String()
         .          .     70:		})
         .          .     71:		de.Lock()
         .          .     72:
         .          .     73:		if de.conn == nil {
         .          .     74:			if err := de.refresh(de.addr); err != nil {
         .          .     75:				de.Unlock()
         .          .     76:				return
         .          .     77:			}
         .          .     78:		}
         .          .     79:
     52444      52444     80:		_, err := de.conn.Write(append(Header, p...))
         .          .     81:		if err != nil {
         .          .     82:			logger.Error(err)
         .          .     83:		}
         .          .     84:		de.Unlock()
         .          .     85:	}
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).Close in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
     32768     693276 (flat, cum) 38.62% of Total
         .          .    236:	seg.Lock()
         .          .    237:	defer seg.Unlock()
         .          .    238:	if seg.parent != nil {
         .          .    239:		logger.Debugf("Closing subsegment named %s", seg.Name)
         .          .    240:	} else {
     32768     131076    241:		logger.Debugf("Closing segment named %s", seg.Name)
         .          .    242:	}
         .          .    243:	seg.EndTime = float64(time.Now().UnixNano()) / float64(time.Second)
         .          .    244:	seg.InProgress = false
         .          .    245:
         .          .    246:	if err != nil {
         .          .    247:		seg.addError(err)
         .          .    248:	}
         .          .    249:
         .     562200    250:	seg.flush()
         .          .    251:}
         .          .    252:
         .          .    253:// CloseAndStream closes a subsegment and sends it.
         .          .    254:func (subseg *Segment) CloseAndStream(err error) {
         .          .    255:	subseg.Lock()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).GetAWS in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment_model.go
     43692      43692 (flat, cum)  2.43% of Total
         .          .    160:}
         .          .    161:
         .          .    162:// GetAWS returns value of AWS.
         .          .    163:func (s *Segment) GetAWS() map[string]interface{} {
         .          .    164:	if s.AWS == nil {
     43692      43692    165:		s.AWS = make(map[string]interface{})
         .          .    166:	}
         .          .    167:	return s.AWS
         .          .    168:}
         .          .    169:
         .          .    170:// GetService returns value of Service.
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).GetConfiguration in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment_model.go
     40055      40055 (flat, cum)  2.23% of Total
         .          .    200:}
         .          .    201:
         .          .    202:// GetConfiguration returns a value of Config.
         .          .    203:func (s *Segment) GetConfiguration() *Config {
         .          .    204:	if s.Configuration == nil {
     40055      40055    205:		s.Configuration = &Config{}
         .          .    206:	}
         .          .    207:	return s.Configuration
         .          .    208:}
         .          .    209:
         .          .    210:// AddRuleName adds rule name, if present from sampling decision to xray context.
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).GetService in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment_model.go
     54615      54615 (flat, cum)  3.04% of Total
         .          .    168:}
         .          .    169:
         .          .    170:// GetService returns value of Service.
         .          .    171:func (s *Segment) GetService() *ServiceData {
         .          .    172:	if s.Service == nil {
     54615      54615    173:		s.Service = &ServiceData{}
         .          .    174:	}
         .          .    175:	return s.Service
         .          .    176:}
         .          .    177:
         .          .    178:// GetSQL returns value of SQL.
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).addSDKAndServiceInformation in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
     50986     149293 (flat, cum)  8.32% of Total
         .          .    366:		seg.Origin = metadata.Origin
         .          .    367:	}
         .          .    368:}
         .          .    369:
         .          .    370:func (seg *Segment) addSDKAndServiceInformation() {
     50986      94678    371:	seg.GetAWS()["xray"] = SDK{Version: SDKVersion, Type: SDKType}
         .          .    372:
         .      54615    373:	seg.GetService().Compiler = runtime.Compiler
         .          .    374:	seg.GetService().CompilerVersion = runtime.Version()
         .          .    375:}
         .          .    376:
         .          .    377:func (sub *Segment) beforeEmitSubsegment(seg *Segment) {
         .          .    378:	// Only called within a subsegment locked code block
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).assignConfiguration in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0      40055 (flat, cum)  2.23% of Total
         .          .    108:
         .          .    109:// assignConfiguration assigns value to seg.Configuration
         .          .    110:func (seg *Segment) assignConfiguration(cfg *Config) {
         .          .    111:	seg.Lock()
         .          .    112:	if cfg == nil {
         .      40055    113:		seg.GetConfiguration().ContextMissingStrategy = globalCfg.contextMissingStrategy
         .          .    114:		seg.GetConfiguration().ExceptionFormattingStrategy = globalCfg.exceptionFormattingStrategy
         .          .    115:		seg.GetConfiguration().SamplingStrategy = globalCfg.samplingStrategy
         .          .    116:		seg.GetConfiguration().StreamingStrategy = globalCfg.streamingStrategy
         .          .    117:		seg.GetConfiguration().Emitter = globalCfg.emitter
         .          .    118:		seg.GetConfiguration().ServiceVersion = globalCfg.serviceVersion
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).emit in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0     562200 (flat, cum) 31.31% of Total
         .          .    292:	}
         .          .    293:	return false
         .          .    294:}
         .          .    295:
         .          .    296:func (seg *Segment) emit() {
         .     562200    297:	seg.ParentSegment.GetConfiguration().Emitter.Emit(seg)
         .          .    298:}
         .          .    299:
         .          .    300:func (seg *Segment) handleContextDone() {
         .          .    301:	seg.Lock()
         .          .    302:	defer seg.Unlock()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).flush in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0     562200 (flat, cum) 31.31% of Total
         .          .    309:
         .          .    310:func (seg *Segment) flush() {
         .          .    311:	if (seg.openSegments == 0 && seg.EndTime > 0) || seg.ContextDone {
         .          .    312:		if seg.parent == nil {
         .          .    313:			seg.Emitted = true
         .     562200    314:			seg.emit()
         .          .    315:		} else if seg.parent != nil && seg.parent.Facade {
         .          .    316:			seg.Emitted = true
         .          .    317:			seg.beforeEmitSubsegment(seg.parent)
         .          .    318:			logger.Debugf("emit lambda subsegment named: %v", seg.Name)
         .          .    319:			seg.emit()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.BeginSegment in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0     947141 (flat, cum) 52.76% of Total
         .          .     51:	return context.WithValue(ctx, ContextKey, seg), seg
         .          .     52:}
         .          .     53:
         .          .     54:// BeginSegment creates a Segment for a given name and context.
         .          .     55:func BeginSegment(ctx context.Context, name string) (context.Context, *Segment) {
         .     725024     56:	seg := basicSegment(name, nil)
         .          .     57:
         .          .     58:	cfg := GetRecorder(ctx)
         .      40055     59:	seg.assignConfiguration(cfg)
         .          .     60:
         .          .     61:	seg.Lock()
         .          .     62:	defer seg.Unlock()
         .          .     63:
         .          .     64:	seg.addPlugin(plugins.InstancePluginMetadata)
         .     149293     65:	seg.addSDKAndServiceInformation()
         .          .     66:	if seg.ParentSegment.GetConfiguration().ServiceVersion != "" {
         .          .     67:		seg.GetService().Version = seg.ParentSegment.GetConfiguration().ServiceVersion
         .          .     68:	}
         .          .     69:
         .          .     70:	go func() {
         .          .     71:		select {
         .          .     72:		case <-ctx.Done():
         .          .     73:			seg.handleContextDone()
         .          .     74:		}
         .          .     75:	}()
         .          .     76:
         .      32769     77:	return context.WithValue(ctx, ContextKey, seg), seg
         .          .     78:}
         .          .     79:
         .          .     80:func basicSegment(name string, h *header.Header) *Segment {
         .          .     81:	if len(name) > 200 {
         .          .     82:		name = name[:200]
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.NewSegmentID in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
     65537      98305 (flat, cum)  5.48% of Total
         .          .     36:	var r [8]byte
         .          .     37:	_, err := rand.Read(r[:])
         .          .     38:	if err != nil {
         .          .     39:		panic(err)
         .          .     40:	}
     65537      98305     41:	return fmt.Sprintf("%02x", r)
         .          .     42:}
         .          .     43:
         .          .     44:// BeginFacadeSegment creates a Segment for a given name and context.
         .          .     45:func BeginFacadeSegment(ctx context.Context, name string, h *header.Header) (context.Context, *Segment) {
         .          .     46:	seg := basicSegment(name, h)
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.NewTraceID in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
    196611     327685 (flat, cum) 18.25% of Total
         .          .     21:	"github.com/aws/aws-xray-sdk-go/internal/plugins"
         .          .     22:)
         .          .     23:
         .          .     24:// NewTraceID generates a string format of random trace ID.
         .          .     25:func NewTraceID() string {
     65537      65537     26:	var r [12]byte
         .          .     27:	_, err := rand.Read(r[:])
         .          .     28:	if err != nil {
         .          .     29:		panic(err)
         .          .     30:	}
    131074     262148     31:	return fmt.Sprintf("1-%08x-%02x", time.Now().Unix(), r)
         .          .     32:}
         .          .     33:
         .          .     34:// NewSegmentID generates a string format of segment ID.
         .          .     35:func NewSegmentID() string {
         .          .     36:	var r [8]byte
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.basicSegment in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
    222572     725024 (flat, cum) 40.38% of Total
         .          .     79:
         .          .     80:func basicSegment(name string, h *header.Header) *Segment {
         .          .     81:	if len(name) > 200 {
         .          .     82:		name = name[:200]
         .          .     83:	}
     58730      58730     84:	seg := &Segment{parent: nil}
    163842     240304     85:	logger.Debugf("Beginning segment named %s", name)
         .          .     86:	seg.ParentSegment = seg
         .          .     87:
         .          .     88:	seg.Lock()
         .          .     89:	defer seg.Unlock()
         .          .     90:
         .          .     91:	seg.Name = name
         .          .     92:	seg.StartTime = float64(time.Now().UnixNano()) / float64(time.Second)
         .          .     93:	seg.InProgress = true
         .          .     94:
         .          .     95:	if h == nil {
         .     327685     96:		seg.TraceID = NewTraceID()
         .      98305     97:		seg.ID = NewSegmentID()
         .          .     98:		seg.Sampled = true
         .          .     99:	} else {
         .          .    100:		seg.Facade = true
         .          .    101:		seg.ID = h.ParentID
         .          .    102:		seg.TraceID = h.TraceID
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.packSegments in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/default_emitter.go
     81921     362297 (flat, cum) 20.18% of Total
         .          .     83:		}
         .          .     84:		de.Unlock()
         .          .     85:	}
         .          .     86:}
         .          .     87:
     49153      49153     88:func packSegments(seg *Segment, outSegments [][]byte) [][]byte {
     16384      16384     89:	trimSubsegment := func(s *Segment) []byte {
         .          .     90:		ss := globalCfg.StreamingStrategy()
         .          .     91:		if seg.ParentSegment.Configuration != nil && seg.ParentSegment.Configuration.StreamingStrategy != nil {
         .          .     92:			ss = seg.ParentSegment.Configuration.StreamingStrategy
         .          .     93:		}
         .          .     94:		for ss.RequiresStreaming(s) {
         .          .     95:			if len(s.rawSubsegments) == 0 {
         .          .     96:				break
         .          .     97:			}
         .          .     98:			cb := ss.StreamCompletedSubsegments(s)
         .          .     99:			outSegments = append(outSegments, cb...)
         .          .    100:		}
         .          .    101:		b, _ := json.Marshal(s)
         .          .    102:		return b
         .          .    103:	}
         .          .    104:
         .          .    105:	for _, s := range seg.rawSubsegments {
         .          .    106:		outSegments = packSegments(s, outSegments)
         .          .    107:		if b := trimSubsegment(s); b != nil {
         .          .    108:			seg.Subsegments = append(seg.Subsegments, b)
         .          .    109:		}
         .          .    110:	}
         .          .    111:	if seg.parent == nil {
         .     280376    112:		if b := trimSubsegment(seg); b != nil {
     16384      16384    113:			outSegments = append(outSegments, b)
         .          .    114:		}
         .          .    115:	}
         .          .    116:	return outSegments
         .          .    117:}
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.packSegments.func1 in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/default_emitter.go
         0     280376 (flat, cum) 15.62% of Total
         .          .     96:				break
         .          .     97:			}
         .          .     98:			cb := ss.StreamCompletedSubsegments(s)
         .          .     99:			outSegments = append(outSegments, cb...)
         .          .    100:		}
         .     280376    101:		b, _ := json.Marshal(s)
         .          .    102:		return b
         .          .    103:	}
         .          .    104:
         .          .    105:	for _, s := range seg.rawSubsegments {
         .          .    106:		outSegments = packSegments(s, outSegments)
ROUTINE ======================== github.com/cep21/gobenchtraces.xrayRun in /Users/jlindamo/IdeaProjects/gobenchtraces/xray_test.go
         0    1640417 (flat, cum) 91.37% of Total
         .          .    207:	span.Finish()
         .          .    208:}
         .          .    209:
         .          .    210:func xrayRun(b *testing.B, run benchmarkTracesRun) {
         .          .    211:	ctx := context.Background()
         .     947141    212:	_, s := xray.BeginSegment(ctx, "start")
         .          .    213:	run.toCall()
         .     693276    214:	s.Close(nil)
         .          .    215:}
         .          .    216:
         .          .    217:func ddRun(b *testing.B, run benchmarkTracesRun) {
         .          .    218:	span := tracer.StartSpan("test")
         .          .    219:	run.toCall()
(pprof) 
```
