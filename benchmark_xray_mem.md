```
(pprof) Total: 1920135
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/internal/logger.Debugf in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/internal/logger/logger.go
    142000     142000 (flat, cum)  7.40% of Total
         .          .     19:
         .          .     20:// The Logger instance used by xray to log. Set via xray.SetLogger().
         .          .     21:var Logger xraylog.Logger = xraylog.NewDefaultLogger(os.Stdout, xraylog.LogLevelInfo)
         .          .     22:
         .          .     23:func Debugf(format string, args ...interface{}) {
    142000     142000     24:	Logger.Log(xraylog.LogLevelDebug, printfArgs{format, args})
         .          .     25:}
         .          .     26:
         .          .     27:func Debug(args ...interface{}) {
         .          .     28:	Logger.Log(xraylog.LogLevelDebug, printArgs(args))
         .          .     29:}
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*DefaultEmitter).Emit in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/default_emitter.go
    150750     763548 (flat, cum) 39.77% of Total
         .          .     59:func (de *DefaultEmitter) Emit(seg *Segment) {
         .          .     60:	if seg == nil || !seg.ParentSegment.Sampled {
         .          .     61:		return
         .          .     62:	}
         .          .     63:
     65538     671782     64:	for _, p := range packSegments(seg, nil) {
         .          .     65:		// defer expensive marshal until log message is actually logged
     32768      32768     66:		logger.DebugDeferred(func() string {
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
     52444      58998     80:		_, err := de.conn.Write(append(Header, p...))
         .          .     81:		if err != nil {
         .          .     82:			logger.Error(err)
         .          .     83:		}
         .          .     84:		de.Unlock()
         .          .     85:	}
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).Close in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
     65537     916470 (flat, cum) 47.73% of Total
         .          .    236:	seg.Lock()
         .          .    237:	defer seg.Unlock()
         .          .    238:	if seg.parent != nil {
         .          .    239:		logger.Debugf("Closing subsegment named %s", seg.Name)
         .          .    240:	} else {
     65537     152922    241:		logger.Debugf("Closing segment named %s", seg.Name)
         .          .    242:	}
         .          .    243:	seg.EndTime = float64(time.Now().UnixNano()) / float64(time.Second)
         .          .    244:	seg.InProgress = false
         .          .    245:
         .          .    246:	if err != nil {
         .          .    247:		seg.addError(err)
         .          .    248:	}
         .          .    249:
         .     763548    250:	seg.flush()
         .          .    251:}
         .          .    252:
         .          .    253:// CloseAndStream closes a subsegment and sends it.
         .          .    254:func (subseg *Segment) CloseAndStream(err error) {
         .          .    255:	subseg.Lock()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).GetAWS in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment_model.go
     43692      43692 (flat, cum)  2.28% of Total
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
     58262      58262 (flat, cum)  3.03% of Total
         .          .    200:}
         .          .    201:
         .          .    202:// GetConfiguration returns a value of Config.
         .          .    203:func (s *Segment) GetConfiguration() *Config {
         .          .    204:	if s.Configuration == nil {
     58262      58262    205:		s.Configuration = &Config{}
         .          .    206:	}
         .          .    207:	return s.Configuration
         .          .    208:}
         .          .    209:
         .          .    210:// AddRuleName adds rule name, if present from sampling decision to xray context.
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).GetService in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment_model.go
     54615      54615 (flat, cum)  2.84% of Total
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
    138371     236678 (flat, cum) 12.33% of Total
         .          .    366:		seg.Origin = metadata.Origin
         .          .    367:	}
         .          .    368:}
         .          .    369:
         .          .    370:func (seg *Segment) addSDKAndServiceInformation() {
    138371     182063    371:	seg.GetAWS()["xray"] = SDK{Version: SDKVersion, Type: SDKType}
         .          .    372:
         .      54615    373:	seg.GetService().Compiler = runtime.Compiler
         .          .    374:	seg.GetService().CompilerVersion = runtime.Version()
         .          .    375:}
         .          .    376:
         .          .    377:func (sub *Segment) beforeEmitSubsegment(seg *Segment) {
         .          .    378:	// Only called within a subsegment locked code block
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).assignConfiguration in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0      58262 (flat, cum)  3.03% of Total
         .          .    108:
         .          .    109:// assignConfiguration assigns value to seg.Configuration
         .          .    110:func (seg *Segment) assignConfiguration(cfg *Config) {
         .          .    111:	seg.Lock()
         .          .    112:	if cfg == nil {
         .      58262    113:		seg.GetConfiguration().ContextMissingStrategy = globalCfg.contextMissingStrategy
         .          .    114:		seg.GetConfiguration().ExceptionFormattingStrategy = globalCfg.exceptionFormattingStrategy
         .          .    115:		seg.GetConfiguration().SamplingStrategy = globalCfg.samplingStrategy
         .          .    116:		seg.GetConfiguration().StreamingStrategy = globalCfg.streamingStrategy
         .          .    117:		seg.GetConfiguration().Emitter = globalCfg.emitter
         .          .    118:		seg.GetConfiguration().ServiceVersion = globalCfg.serviceVersion
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).emit in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0     763548 (flat, cum) 39.77% of Total
         .          .    292:	}
         .          .    293:	return false
         .          .    294:}
         .          .    295:
         .          .    296:func (seg *Segment) emit() {
         .     763548    297:	seg.ParentSegment.GetConfiguration().Emitter.Emit(seg)
         .          .    298:}
         .          .    299:
         .          .    300:func (seg *Segment) handleContextDone() {
         .          .    301:	seg.Lock()
         .          .    302:	defer seg.Unlock()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).flush in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0     763548 (flat, cum) 39.77% of Total
         .          .    309:
         .          .    310:func (seg *Segment) flush() {
         .          .    311:	if (seg.openSegments == 0 && seg.EndTime > 0) || seg.ContextDone {
         .          .    312:		if seg.parent == nil {
         .          .    313:			seg.Emitted = true
         .     763548    314:			seg.emit()
         .          .    315:		} else if seg.parent != nil && seg.parent.Facade {
         .          .    316:			seg.Emitted = true
         .          .    317:			seg.beforeEmitSubsegment(seg.parent)
         .          .    318:			logger.Debugf("emit lambda subsegment named: %v", seg.Name)
         .          .    319:			seg.emit()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.BeginSegment in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0     928484 (flat, cum) 48.36% of Total
         .          .     51:	return context.WithValue(ctx, ContextKey, seg), seg
         .          .     52:}
         .          .     53:
         .          .     54:// BeginSegment creates a Segment for a given name and context.
         .          .     55:func BeginSegment(ctx context.Context, name string) (context.Context, *Segment) {
         .     589852     56:	seg := basicSegment(name, nil)
         .          .     57:
         .          .     58:	cfg := GetRecorder(ctx)
         .      58262     59:	seg.assignConfiguration(cfg)
         .          .     60:
         .          .     61:	seg.Lock()
         .          .     62:	defer seg.Unlock()
         .          .     63:
         .          .     64:	seg.addPlugin(plugins.InstancePluginMetadata)
         .     236678     65:	seg.addSDKAndServiceInformation()
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
         .      43692     77:	return context.WithValue(ctx, ContextKey, seg), seg
         .          .     78:}
         .          .     79:
         .          .     80:func basicSegment(name string, h *header.Header) *Segment {
         .          .     81:	if len(name) > 200 {
         .          .     82:		name = name[:200]
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.NewSegmentID in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
     32768     196610 (flat, cum) 10.24% of Total
         .          .     36:	var r [8]byte
         .          .     37:	_, err := rand.Read(r[:])
         .          .     38:	if err != nil {
         .          .     39:		panic(err)
         .          .     40:	}
     32768     196610     41:	return fmt.Sprintf("%02x", r)
         .          .     42:}
         .          .     43:
         .          .     44:// BeginFacadeSegment creates a Segment for a given name and context.
         .          .     45:func BeginFacadeSegment(ctx context.Context, name string, h *header.Header) (context.Context, *Segment) {
         .          .     46:	seg := basicSegment(name, h)
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.NewTraceID in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
     98305     180226 (flat, cum)  9.39% of Total
         .          .     21:	"github.com/aws/aws-xray-sdk-go/internal/plugins"
         .          .     22:)
         .          .     23:
         .          .     24:// NewTraceID generates a string format of random trace ID.
         .          .     25:func NewTraceID() string {
     98305      98305     26:	var r [12]byte
         .       5461     27:	_, err := rand.Read(r[:])
         .          .     28:	if err != nil {
         .          .     29:		panic(err)
         .          .     30:	}
         .      76460     31:	return fmt.Sprintf("1-%08x-%02x", time.Now().Unix(), r)
         .          .     32:}
         .          .     33:
         .          .     34:// NewSegmentID generates a string format of segment ID.
         .          .     35:func NewSegmentID() string {
         .          .     36:	var r [8]byte
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.basicSegment in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
    158401     589852 (flat, cum) 30.72% of Total
         .          .     79:
         .          .     80:func basicSegment(name string, h *header.Header) *Segment {
         .          .     81:	if len(name) > 200 {
         .          .     82:		name = name[:200]
         .          .     83:	}
     60096      60096     84:	seg := &Segment{parent: nil}
     98305     152920     85:	logger.Debugf("Beginning segment named %s", name)
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
         .     180226     96:		seg.TraceID = NewTraceID()
         .     196610     97:		seg.ID = NewSegmentID()
         .          .     98:		seg.Sampled = true
         .          .     99:	} else {
         .          .    100:		seg.Facade = true
         .          .    101:		seg.ID = h.ParentID
         .          .    102:		seg.TraceID = h.TraceID
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.packSegments in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/default_emitter.go
    131076     606244 (flat, cum) 31.57% of Total
         .          .     83:		}
         .          .     84:		de.Unlock()
         .          .     85:	}
         .          .     86:}
         .          .     87:
     32769      32769     88:func packSegments(seg *Segment, outSegments [][]byte) [][]byte {
     65538      65538     89:	trimSubsegment := func(s *Segment) []byte {
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
         .     475168    112:		if b := trimSubsegment(seg); b != nil {
     32769      32769    113:			outSegments = append(outSegments, b)
         .          .    114:		}
         .          .    115:	}
         .          .    116:	return outSegments
         .          .    117:}
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.packSegments.func1 in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/default_emitter.go
         0     475168 (flat, cum) 24.75% of Total
         .          .     96:				break
         .          .     97:			}
         .          .     98:			cb := ss.StreamCompletedSubsegments(s)
         .          .     99:			outSegments = append(outSegments, cb...)
         .          .    100:		}
         .     475168    101:		b, _ := json.Marshal(s)
         .          .    102:		return b
         .          .    103:	}
         .          .    104:
         .          .    105:	for _, s := range seg.rawSubsegments {
         .          .    106:		outSegments = packSegments(s, outSegments)
ROUTINE ======================== github.com/cep21/gobenchtraces.xrayRun in /Users/jlindamo/IdeaProjects/gobenchtraces/xray_test.go
         0    1844954 (flat, cum) 96.08% of Total
         .          .    207:	span.Finish()
         .          .    208:}
         .          .    209:
         .          .    210:func xrayRun(b *testing.B, run benchmarkTracesRun) {
         .          .    211:	ctx := context.Background()
         .     928484    212:	_, s := xray.BeginSegment(ctx, "start")
         .          .    213:	run.toCall()
         .     916470    214:	s.Close(nil)
         .          .    215:}
         .          .    216:
         .          .    217:func ddRun(b *testing.B, run benchmarkTracesRun) {
         .          .    218:	span := tracer.StartSpan("test")
         .          .    219:	run.toCall()
(pprof) 
```
