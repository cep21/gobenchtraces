```
(pprof) Total: 3.69s
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/internal/logger.Debugf in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/internal/logger/logger.go
         0       10ms (flat, cum)  0.27% of Total
         .          .     19:
         .          .     20:// The Logger instance used by xray to log. Set via xray.SetLogger().
         .          .     21:var Logger xraylog.Logger = xraylog.NewDefaultLogger(os.Stdout, xraylog.LogLevelInfo)
         .          .     22:
         .          .     23:func Debugf(format string, args ...interface{}) {
         .       10ms     24:	Logger.Log(xraylog.LogLevelDebug, printfArgs{format, args})
         .          .     25:}
         .          .     26:
         .          .     27:func Debug(args ...interface{}) {
         .          .     28:	Logger.Log(xraylog.LogLevelDebug, printArgs(args))
         .          .     29:}
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*DefaultEmitter).Emit in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/default_emitter.go
         0      520ms (flat, cum) 14.09% of Total
         .          .     59:func (de *DefaultEmitter) Emit(seg *Segment) {
         .          .     60:	if seg == nil || !seg.ParentSegment.Sampled {
         .          .     61:		return
         .          .     62:	}
         .          .     63:
         .      120ms     64:	for _, p := range packSegments(seg, nil) {
         .          .     65:		// defer expensive marshal until log message is actually logged
         .          .     66:		logger.DebugDeferred(func() string {
         .          .     67:			var b bytes.Buffer
         .          .     68:			json.Indent(&b, p, "", " ")
         .          .     69:			return b.String()
         .          .     70:		})
         .       70ms     71:		de.Lock()
         .          .     72:
         .          .     73:		if de.conn == nil {
         .          .     74:			if err := de.refresh(de.addr); err != nil {
         .          .     75:				de.Unlock()
         .          .     76:				return
         .          .     77:			}
         .          .     78:		}
         .          .     79:
         .      320ms     80:		_, err := de.conn.Write(append(Header, p...))
         .          .     81:		if err != nil {
         .          .     82:			logger.Error(err)
         .          .     83:		}
         .       10ms     84:		de.Unlock()
         .          .     85:	}
         .          .     86:}
         .          .     87:
         .          .     88:func packSegments(seg *Segment, outSegments [][]byte) [][]byte {
         .          .     89:	trimSubsegment := func(s *Segment) []byte {
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).Close in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0      540ms (flat, cum) 14.63% of Total
         .          .    236:	seg.Lock()
         .          .    237:	defer seg.Unlock()
         .          .    238:	if seg.parent != nil {
         .          .    239:		logger.Debugf("Closing subsegment named %s", seg.Name)
         .          .    240:	} else {
         .       10ms    241:		logger.Debugf("Closing segment named %s", seg.Name)
         .          .    242:	}
         .       10ms    243:	seg.EndTime = float64(time.Now().UnixNano()) / float64(time.Second)
         .          .    244:	seg.InProgress = false
         .          .    245:
         .          .    246:	if err != nil {
         .          .    247:		seg.addError(err)
         .          .    248:	}
         .          .    249:
         .      520ms    250:	seg.flush()
         .          .    251:}
         .          .    252:
         .          .    253:// CloseAndStream closes a subsegment and sends it.
         .          .    254:func (subseg *Segment) CloseAndStream(err error) {
         .          .    255:	subseg.Lock()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).GetConfiguration in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment_model.go
         0       10ms (flat, cum)  0.27% of Total
         .          .    200:}
         .          .    201:
         .          .    202:// GetConfiguration returns a value of Config.
         .          .    203:func (s *Segment) GetConfiguration() *Config {
         .          .    204:	if s.Configuration == nil {
         .       10ms    205:		s.Configuration = &Config{}
         .          .    206:	}
         .          .    207:	return s.Configuration
         .          .    208:}
         .          .    209:
         .          .    210:// AddRuleName adds rule name, if present from sampling decision to xray context.
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).addSDKAndServiceInformation in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0       20ms (flat, cum)  0.54% of Total
         .          .    366:		seg.Origin = metadata.Origin
         .          .    367:	}
         .          .    368:}
         .          .    369:
         .          .    370:func (seg *Segment) addSDKAndServiceInformation() {
         .       20ms    371:	seg.GetAWS()["xray"] = SDK{Version: SDKVersion, Type: SDKType}
         .          .    372:
         .          .    373:	seg.GetService().Compiler = runtime.Compiler
         .          .    374:	seg.GetService().CompilerVersion = runtime.Version()
         .          .    375:}
         .          .    376:
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).assignConfiguration in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0       10ms (flat, cum)  0.27% of Total
         .          .    108:
         .          .    109:// assignConfiguration assigns value to seg.Configuration
         .          .    110:func (seg *Segment) assignConfiguration(cfg *Config) {
         .          .    111:	seg.Lock()
         .          .    112:	if cfg == nil {
         .       10ms    113:		seg.GetConfiguration().ContextMissingStrategy = globalCfg.contextMissingStrategy
         .          .    114:		seg.GetConfiguration().ExceptionFormattingStrategy = globalCfg.exceptionFormattingStrategy
         .          .    115:		seg.GetConfiguration().SamplingStrategy = globalCfg.samplingStrategy
         .          .    116:		seg.GetConfiguration().StreamingStrategy = globalCfg.streamingStrategy
         .          .    117:		seg.GetConfiguration().Emitter = globalCfg.emitter
         .          .    118:		seg.GetConfiguration().ServiceVersion = globalCfg.serviceVersion
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).emit in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0      520ms (flat, cum) 14.09% of Total
         .          .    292:	}
         .          .    293:	return false
         .          .    294:}
         .          .    295:
         .          .    296:func (seg *Segment) emit() {
         .      520ms    297:	seg.ParentSegment.GetConfiguration().Emitter.Emit(seg)
         .          .    298:}
         .          .    299:
         .          .    300:func (seg *Segment) handleContextDone() {
         .          .    301:	seg.Lock()
         .          .    302:	defer seg.Unlock()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).flush in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0      520ms (flat, cum) 14.09% of Total
         .          .    309:
         .          .    310:func (seg *Segment) flush() {
         .          .    311:	if (seg.openSegments == 0 && seg.EndTime > 0) || seg.ContextDone {
         .          .    312:		if seg.parent == nil {
         .          .    313:			seg.Emitted = true
         .      520ms    314:			seg.emit()
         .          .    315:		} else if seg.parent != nil && seg.parent.Facade {
         .          .    316:			seg.Emitted = true
         .          .    317:			seg.beforeEmitSubsegment(seg.parent)
         .          .    318:			logger.Debugf("emit lambda subsegment named: %v", seg.Name)
         .          .    319:			seg.emit()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.BeginSegment in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0      260ms (flat, cum)  7.05% of Total
         .          .     51:	return context.WithValue(ctx, ContextKey, seg), seg
         .          .     52:}
         .          .     53:
         .          .     54:// BeginSegment creates a Segment for a given name and context.
         .          .     55:func BeginSegment(ctx context.Context, name string) (context.Context, *Segment) {
         .       50ms     56:	seg := basicSegment(name, nil)
         .          .     57:
         .          .     58:	cfg := GetRecorder(ctx)
         .       10ms     59:	seg.assignConfiguration(cfg)
         .          .     60:
         .          .     61:	seg.Lock()
         .          .     62:	defer seg.Unlock()
         .          .     63:
         .          .     64:	seg.addPlugin(plugins.InstancePluginMetadata)
         .       20ms     65:	seg.addSDKAndServiceInformation()
         .          .     66:	if seg.ParentSegment.GetConfiguration().ServiceVersion != "" {
         .          .     67:		seg.GetService().Version = seg.ParentSegment.GetConfiguration().ServiceVersion
         .          .     68:	}
         .          .     69:
         .      170ms     70:	go func() {
         .          .     71:		select {
         .          .     72:		case <-ctx.Done():
         .          .     73:			seg.handleContextDone()
         .          .     74:		}
         .          .     75:	}()
         .          .     76:
         .       10ms     77:	return context.WithValue(ctx, ContextKey, seg), seg
         .          .     78:}
         .          .     79:
         .          .     80:func basicSegment(name string, h *header.Header) *Segment {
         .          .     81:	if len(name) > 200 {
         .          .     82:		name = name[:200]
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.BeginSegment.func1 in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0       10ms (flat, cum)  0.27% of Total
         .          .     67:		seg.GetService().Version = seg.ParentSegment.GetConfiguration().ServiceVersion
         .          .     68:	}
         .          .     69:
         .          .     70:	go func() {
         .          .     71:		select {
         .       10ms     72:		case <-ctx.Done():
         .          .     73:			seg.handleContextDone()
         .          .     74:		}
         .          .     75:	}()
         .          .     76:
         .          .     77:	return context.WithValue(ctx, ContextKey, seg), seg
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.NewTraceID in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0       10ms (flat, cum)  0.27% of Total
         .          .     26:	var r [12]byte
         .          .     27:	_, err := rand.Read(r[:])
         .          .     28:	if err != nil {
         .          .     29:		panic(err)
         .          .     30:	}
         .       10ms     31:	return fmt.Sprintf("1-%08x-%02x", time.Now().Unix(), r)
         .          .     32:}
         .          .     33:
         .          .     34:// NewSegmentID generates a string format of segment ID.
         .          .     35:func NewSegmentID() string {
         .          .     36:	var r [8]byte
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.basicSegment in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0       50ms (flat, cum)  1.36% of Total
         .          .     79:
         .          .     80:func basicSegment(name string, h *header.Header) *Segment {
         .          .     81:	if len(name) > 200 {
         .          .     82:		name = name[:200]
         .          .     83:	}
         .       20ms     84:	seg := &Segment{parent: nil}
         .          .     85:	logger.Debugf("Beginning segment named %s", name)
         .          .     86:	seg.ParentSegment = seg
         .          .     87:
         .          .     88:	seg.Lock()
         .          .     89:	defer seg.Unlock()
         .          .     90:
         .          .     91:	seg.Name = name
         .       20ms     92:	seg.StartTime = float64(time.Now().UnixNano()) / float64(time.Second)
         .          .     93:	seg.InProgress = true
         .          .     94:
         .          .     95:	if h == nil {
         .       10ms     96:		seg.TraceID = NewTraceID()
         .          .     97:		seg.ID = NewSegmentID()
         .          .     98:		seg.Sampled = true
         .          .     99:	} else {
         .          .    100:		seg.Facade = true
         .          .    101:		seg.ID = h.ParentID
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.packSegments in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/default_emitter.go
         0      120ms (flat, cum)  3.25% of Total
         .          .    107:		if b := trimSubsegment(s); b != nil {
         .          .    108:			seg.Subsegments = append(seg.Subsegments, b)
         .          .    109:		}
         .          .    110:	}
         .          .    111:	if seg.parent == nil {
         .      120ms    112:		if b := trimSubsegment(seg); b != nil {
         .          .    113:			outSegments = append(outSegments, b)
         .          .    114:		}
         .          .    115:	}
         .          .    116:	return outSegments
         .          .    117:}
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.packSegments.func1 in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/default_emitter.go
         0      120ms (flat, cum)  3.25% of Total
         .          .     96:				break
         .          .     97:			}
         .          .     98:			cb := ss.StreamCompletedSubsegments(s)
         .          .     99:			outSegments = append(outSegments, cb...)
         .          .    100:		}
         .      120ms    101:		b, _ := json.Marshal(s)
         .          .    102:		return b
         .          .    103:	}
         .          .    104:
         .          .    105:	for _, s := range seg.rawSubsegments {
         .          .    106:		outSegments = packSegments(s, outSegments)
ROUTINE ======================== github.com/cep21/gobenchtraces.xrayRun in /Users/jlindamo/IdeaProjects/gobenchtraces/xray_test.go
         0      800ms (flat, cum) 21.68% of Total
         .          .    207:	span.Finish()
         .          .    208:}
         .          .    209:
         .          .    210:func xrayRun(b *testing.B, run benchmarkTracesRun) {
         .          .    211:	ctx := context.Background()
         .      260ms    212:	_, s := xray.BeginSegment(ctx, "start")
         .          .    213:	run.toCall()
         .      540ms    214:	s.Close(nil)
         .          .    215:}
         .          .    216:
         .          .    217:func ddRun(b *testing.B, run benchmarkTracesRun) {
         .          .    218:	span := tracer.StartSpan("test")
         .          .    219:	run.toCall()
(pprof) 
```
