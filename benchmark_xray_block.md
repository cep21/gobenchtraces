```go

(pprof) Showing nodes accounting for 27.15mins, 99.23% of 27.36mins total
Dropped 122 nodes (cum <= 0.14mins)
Showing top 10 nodes out of 14
      flat  flat%   sum%        cum   cum%
 27.15mins 99.23% 99.23%  27.15mins 99.23%  sync.(*Mutex).Lock
         0     0% 99.23%   0.16mins  0.57%  crypto/rand.(*devReader).Read
         0     0% 99.23%   0.16mins  0.57%  crypto/rand.Read
         0     0% 99.23%  26.99mins 98.66%  github.com/aws/aws-xray-sdk-go/xray.(*DefaultEmitter).Emit
         0     0% 99.23%  26.99mins 98.66%  github.com/aws/aws-xray-sdk-go/xray.(*Segment).Close
         0     0% 99.23%  26.99mins 98.66%  github.com/aws/aws-xray-sdk-go/xray.(*Segment).emit
         0     0% 99.23%  26.99mins 98.66%  github.com/aws/aws-xray-sdk-go/xray.(*Segment).flush
         0     0% 99.23%   0.16mins  0.57%  github.com/aws/aws-xray-sdk-go/xray.BeginSegment
         0     0% 99.23%   0.15mins  0.54%  github.com/aws/aws-xray-sdk-go/xray.NewTraceID
         0     0% 99.23%   0.16mins  0.57%  github.com/aws/aws-xray-sdk-go/xray.basicSegment
(pprof) 
(pprof) Showing nodes accounting for 27.15mins, 99.23% of 27.36mins total
Dropped 122 nodes (cum <= 0.14mins)
Showing top 10 nodes out of 14
      flat  flat%   sum%        cum   cum%
         0     0%     0%  27.15mins 99.23%  github.com/cep21/gobenchtraces.BenchmarkTraces.func2.1
         0     0%     0%  27.15mins 99.23%  github.com/cep21/gobenchtraces.xrayRun
 27.15mins 99.23% 99.23%  27.15mins 99.23%  sync.(*Mutex).Lock
         0     0% 99.23%  26.99mins 98.66%  github.com/aws/aws-xray-sdk-go/xray.(*DefaultEmitter).Emit
         0     0% 99.23%  26.99mins 98.66%  github.com/aws/aws-xray-sdk-go/xray.(*Segment).Close
         0     0% 99.23%  26.99mins 98.66%  github.com/aws/aws-xray-sdk-go/xray.(*Segment).emit
         0     0% 99.23%  26.99mins 98.66%  github.com/aws/aws-xray-sdk-go/xray.(*Segment).flush
         0     0% 99.23%   0.16mins  0.57%  crypto/rand.(*devReader).Read
         0     0% 99.23%   0.16mins  0.57%  crypto/rand.Read
         0     0% 99.23%   0.16mins  0.57%  github.com/aws/aws-xray-sdk-go/xray.BeginSegment
(pprof) 
(pprof) Total: 27.36mins
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*DefaultEmitter).Emit in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/default_emitter.go
         0  26.99mins (flat, cum) 98.66% of Total
         .          .     59:func (de *DefaultEmitter) Emit(seg *Segment) {
         .          .     60:	if seg == nil || !seg.ParentSegment.Sampled {
         .          .     61:		return
         .          .     62:	}
         .          .     63:
         .    16.41ms     64:	for _, p := range packSegments(seg, nil) {
         .          .     65:		// defer expensive marshal until log message is actually logged
         .          .     66:		logger.DebugDeferred(func() string {
         .          .     67:			var b bytes.Buffer
         .          .     68:			json.Indent(&b, p, "", " ")
         .          .     69:			return b.String()
         .          .     70:		})
         .  26.99mins     71:		de.Lock()
         .          .     72:
         .          .     73:		if de.conn == nil {
         .          .     74:			if err := de.refresh(de.addr); err != nil {
         .          .     75:				de.Unlock()
         .          .     76:				return
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).Close in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0  26.99mins (flat, cum) 98.66% of Total
         .          .    245:
         .          .    246:	if err != nil {
         .          .    247:		seg.addError(err)
         .          .    248:	}
         .          .    249:
         .  26.99mins    250:	seg.flush()
         .          .    251:}
         .          .    252:
         .          .    253:// CloseAndStream closes a subsegment and sends it.
         .          .    254:func (subseg *Segment) CloseAndStream(err error) {
         .          .    255:	subseg.Lock()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).emit in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0  26.99mins (flat, cum) 98.66% of Total
         .          .    292:	}
         .          .    293:	return false
         .          .    294:}
         .          .    295:
         .          .    296:func (seg *Segment) emit() {
         .  26.99mins    297:	seg.ParentSegment.GetConfiguration().Emitter.Emit(seg)
         .          .    298:}
         .          .    299:
         .          .    300:func (seg *Segment) handleContextDone() {
         .          .    301:	seg.Lock()
         .          .    302:	defer seg.Unlock()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).flush in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0  26.99mins (flat, cum) 98.66% of Total
         .          .    309:
         .          .    310:func (seg *Segment) flush() {
         .          .    311:	if (seg.openSegments == 0 && seg.EndTime > 0) || seg.ContextDone {
         .          .    312:		if seg.parent == nil {
         .          .    313:			seg.Emitted = true
         .  26.99mins    314:			seg.emit()
         .          .    315:		} else if seg.parent != nil && seg.parent.Facade {
         .          .    316:			seg.Emitted = true
         .          .    317:			seg.beforeEmitSubsegment(seg.parent)
         .          .    318:			logger.Debugf("emit lambda subsegment named: %v", seg.Name)
         .          .    319:			seg.emit()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.BeginSegment in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0      9.39s (flat, cum)  0.57% of Total
         .          .     51:	return context.WithValue(ctx, ContextKey, seg), seg
         .          .     52:}
         .          .     53:
         .          .     54:// BeginSegment creates a Segment for a given name and context.
         .          .     55:func BeginSegment(ctx context.Context, name string) (context.Context, *Segment) {
         .      9.39s     56:	seg := basicSegment(name, nil)
         .          .     57:
         .          .     58:	cfg := GetRecorder(ctx)
         .          .     59:	seg.assignConfiguration(cfg)
         .          .     60:
         .          .     61:	seg.Lock()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.NewSegmentID in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0   475.55ms (flat, cum) 0.029% of Total
         .          .     32:}
         .          .     33:
         .          .     34:// NewSegmentID generates a string format of segment ID.
         .          .     35:func NewSegmentID() string {
         .          .     36:	var r [8]byte
         .   475.55ms     37:	_, err := rand.Read(r[:])
         .          .     38:	if err != nil {
         .          .     39:		panic(err)
         .          .     40:	}
         .          .     41:	return fmt.Sprintf("%02x", r)
         .          .     42:}
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.NewTraceID in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0      8.92s (flat, cum)  0.54% of Total
         .          .     22:)
         .          .     23:
         .          .     24:// NewTraceID generates a string format of random trace ID.
         .          .     25:func NewTraceID() string {
         .          .     26:	var r [12]byte
         .      8.92s     27:	_, err := rand.Read(r[:])
         .          .     28:	if err != nil {
         .          .     29:		panic(err)
         .          .     30:	}
         .          .     31:	return fmt.Sprintf("1-%08x-%02x", time.Now().Unix(), r)
         .          .     32:}
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.basicSegment in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0      9.39s (flat, cum)  0.57% of Total
         .          .     91:	seg.Name = name
         .          .     92:	seg.StartTime = float64(time.Now().UnixNano()) / float64(time.Second)
         .          .     93:	seg.InProgress = true
         .          .     94:
         .          .     95:	if h == nil {
         .      8.92s     96:		seg.TraceID = NewTraceID()
         .   475.55ms     97:		seg.ID = NewSegmentID()
         .          .     98:		seg.Sampled = true
         .          .     99:	} else {
         .          .    100:		seg.Facade = true
         .          .    101:		seg.ID = h.ParentID
         .          .    102:		seg.TraceID = h.TraceID
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.packSegments in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/default_emitter.go
         0    16.41ms (flat, cum) 0.001% of Total
         .          .    107:		if b := trimSubsegment(s); b != nil {
         .          .    108:			seg.Subsegments = append(seg.Subsegments, b)
         .          .    109:		}
         .          .    110:	}
         .          .    111:	if seg.parent == nil {
         .    16.41ms    112:		if b := trimSubsegment(seg); b != nil {
         .          .    113:			outSegments = append(outSegments, b)
         .          .    114:		}
         .          .    115:	}
         .          .    116:	return outSegments
         .          .    117:}
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.packSegments.func1 in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/default_emitter.go
         0    16.41ms (flat, cum) 0.001% of Total
         .          .     96:				break
         .          .     97:			}
         .          .     98:			cb := ss.StreamCompletedSubsegments(s)
         .          .     99:			outSegments = append(outSegments, cb...)
         .          .    100:		}
         .    16.41ms    101:		b, _ := json.Marshal(s)
         .          .    102:		return b
         .          .    103:	}
         .          .    104:
         .          .    105:	for _, s := range seg.rawSubsegments {
         .          .    106:		outSegments = packSegments(s, outSegments)
ROUTINE ======================== github.com/cep21/gobenchtraces.xrayRun in /Users/jlindamo/IdeaProjects/gobenchtraces/xray_test.go
         0  27.15mins (flat, cum) 99.23% of Total
         .          .    207:	span.Finish()
         .          .    208:}
         .          .    209:
         .          .    210:func xrayRun(b *testing.B, run benchmarkTracesRun) {
         .          .    211:	ctx := context.Background()
         .      9.39s    212:	_, s := xray.BeginSegment(ctx, "start")
         .          .    213:	run.toCall()
         .  26.99mins    214:	s.Close(nil)
         .          .    215:}
         .          .    216:
         .          .    217:func ddRun(b *testing.B, run benchmarkTracesRun) {
         .          .    218:	span := tracer.StartSpan("test")
         .          .    219:	run.toCall()
(pprof) 
```
