```go

(pprof) Showing nodes accounting for 1710.81ms, 100% of 1710.81ms total
Dropped 15 nodes (cum <= 8.55ms)
Showing top 10 nodes out of 14
      flat  flat%   sum%        cum   cum%
 1710.81ms   100%   100%  1710.81ms   100%  sync.(*Mutex).Unlock
         0     0%   100%    34.30ms  2.01%  crypto/rand.(*devReader).Read
         0     0%   100%    34.30ms  2.01%  crypto/rand.Read
         0     0%   100%  1676.49ms 97.99%  github.com/aws/aws-xray-sdk-go/xray.(*DefaultEmitter).Emit
         0     0%   100%  1676.49ms 97.99%  github.com/aws/aws-xray-sdk-go/xray.(*Segment).Close
         0     0%   100%  1676.49ms 97.99%  github.com/aws/aws-xray-sdk-go/xray.(*Segment).emit
         0     0%   100%  1676.49ms 97.99%  github.com/aws/aws-xray-sdk-go/xray.(*Segment).flush
         0     0%   100%    34.30ms  2.01%  github.com/aws/aws-xray-sdk-go/xray.BeginSegment
         0     0%   100%    29.75ms  1.74%  github.com/aws/aws-xray-sdk-go/xray.NewTraceID
         0     0%   100%    34.30ms  2.01%  github.com/aws/aws-xray-sdk-go/xray.basicSegment
(pprof) 
(pprof) Showing nodes accounting for 1710.81ms, 100% of 1710.81ms total
Dropped 15 nodes (cum <= 8.55ms)
Showing top 10 nodes out of 14
      flat  flat%   sum%        cum   cum%
 1710.81ms   100%   100%  1710.81ms   100%  sync.(*Mutex).Unlock
         0     0%   100%  1710.79ms   100%  github.com/cep21/gobenchtraces.BenchmarkTraces.func2.1
         0     0%   100%  1710.79ms   100%  github.com/cep21/gobenchtraces.xrayRun
         0     0%   100%  1676.49ms 97.99%  github.com/aws/aws-xray-sdk-go/xray.(*DefaultEmitter).Emit
         0     0%   100%  1676.49ms 97.99%  github.com/aws/aws-xray-sdk-go/xray.(*Segment).Close
         0     0%   100%  1676.49ms 97.99%  github.com/aws/aws-xray-sdk-go/xray.(*Segment).emit
         0     0%   100%  1676.49ms 97.99%  github.com/aws/aws-xray-sdk-go/xray.(*Segment).flush
         0     0%   100%    34.30ms  2.01%  crypto/rand.(*devReader).Read
         0     0%   100%    34.30ms  2.01%  crypto/rand.Read
         0     0%   100%    34.30ms  2.01%  github.com/aws/aws-xray-sdk-go/xray.BeginSegment
(pprof) 
(pprof) Total: 1.71s
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*DefaultEmitter).Emit in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/default_emitter.go
         0      1.68s (flat, cum) 97.99% of Total
         .          .     59:func (de *DefaultEmitter) Emit(seg *Segment) {
         .          .     60:	if seg == nil || !seg.ParentSegment.Sampled {
         .          .     61:		return
         .          .     62:	}
         .          .     63:
         .     4.14us     64:	for _, p := range packSegments(seg, nil) {
         .          .     65:		// defer expensive marshal until log message is actually logged
         .          .     66:		logger.DebugDeferred(func() string {
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
         .          .     80:		_, err := de.conn.Write(append(Header, p...))
         .          .     81:		if err != nil {
         .          .     82:			logger.Error(err)
         .          .     83:		}
         .      1.68s     84:		de.Unlock()
         .          .     85:	}
         .          .     86:}
         .          .     87:
         .          .     88:func packSegments(seg *Segment, outSegments [][]byte) [][]byte {
         .          .     89:	trimSubsegment := func(s *Segment) []byte {
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).Close in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0      1.68s (flat, cum) 97.99% of Total
         .          .    245:
         .          .    246:	if err != nil {
         .          .    247:		seg.addError(err)
         .          .    248:	}
         .          .    249:
         .      1.68s    250:	seg.flush()
         .          .    251:}
         .          .    252:
         .          .    253:// CloseAndStream closes a subsegment and sends it.
         .          .    254:func (subseg *Segment) CloseAndStream(err error) {
         .          .    255:	subseg.Lock()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).emit in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0      1.68s (flat, cum) 97.99% of Total
         .          .    292:	}
         .          .    293:	return false
         .          .    294:}
         .          .    295:
         .          .    296:func (seg *Segment) emit() {
         .      1.68s    297:	seg.ParentSegment.GetConfiguration().Emitter.Emit(seg)
         .          .    298:}
         .          .    299:
         .          .    300:func (seg *Segment) handleContextDone() {
         .          .    301:	seg.Lock()
         .          .    302:	defer seg.Unlock()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.(*Segment).flush in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0      1.68s (flat, cum) 97.99% of Total
         .          .    309:
         .          .    310:func (seg *Segment) flush() {
         .          .    311:	if (seg.openSegments == 0 && seg.EndTime > 0) || seg.ContextDone {
         .          .    312:		if seg.parent == nil {
         .          .    313:			seg.Emitted = true
         .      1.68s    314:			seg.emit()
         .          .    315:		} else if seg.parent != nil && seg.parent.Facade {
         .          .    316:			seg.Emitted = true
         .          .    317:			seg.beforeEmitSubsegment(seg.parent)
         .          .    318:			logger.Debugf("emit lambda subsegment named: %v", seg.Name)
         .          .    319:			seg.emit()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.BeginSegment in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0    34.30ms (flat, cum)  2.01% of Total
         .          .     51:	return context.WithValue(ctx, ContextKey, seg), seg
         .          .     52:}
         .          .     53:
         .          .     54:// BeginSegment creates a Segment for a given name and context.
         .          .     55:func BeginSegment(ctx context.Context, name string) (context.Context, *Segment) {
         .    34.30ms     56:	seg := basicSegment(name, nil)
         .          .     57:
         .          .     58:	cfg := GetRecorder(ctx)
         .          .     59:	seg.assignConfiguration(cfg)
         .          .     60:
         .          .     61:	seg.Lock()
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.NewSegmentID in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0     4.55ms (flat, cum)  0.27% of Total
         .          .     32:}
         .          .     33:
         .          .     34:// NewSegmentID generates a string format of segment ID.
         .          .     35:func NewSegmentID() string {
         .          .     36:	var r [8]byte
         .     4.55ms     37:	_, err := rand.Read(r[:])
         .          .     38:	if err != nil {
         .          .     39:		panic(err)
         .          .     40:	}
         .          .     41:	return fmt.Sprintf("%02x", r)
         .          .     42:}
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.NewTraceID in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0    29.75ms (flat, cum)  1.74% of Total
         .          .     22:)
         .          .     23:
         .          .     24:// NewTraceID generates a string format of random trace ID.
         .          .     25:func NewTraceID() string {
         .          .     26:	var r [12]byte
         .    29.75ms     27:	_, err := rand.Read(r[:])
         .          .     28:	if err != nil {
         .          .     29:		panic(err)
         .          .     30:	}
         .          .     31:	return fmt.Sprintf("1-%08x-%02x", time.Now().Unix(), r)
         .          .     32:}
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.basicSegment in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/segment.go
         0    34.30ms (flat, cum)  2.01% of Total
         .          .     91:	seg.Name = name
         .          .     92:	seg.StartTime = float64(time.Now().UnixNano()) / float64(time.Second)
         .          .     93:	seg.InProgress = true
         .          .     94:
         .          .     95:	if h == nil {
         .    29.75ms     96:		seg.TraceID = NewTraceID()
         .     4.55ms     97:		seg.ID = NewSegmentID()
         .          .     98:		seg.Sampled = true
         .          .     99:	} else {
         .          .    100:		seg.Facade = true
         .          .    101:		seg.ID = h.ParentID
         .          .    102:		seg.TraceID = h.TraceID
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.packSegments in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/default_emitter.go
         0     4.14us (flat, cum) 0.00024% of Total
         .          .    107:		if b := trimSubsegment(s); b != nil {
         .          .    108:			seg.Subsegments = append(seg.Subsegments, b)
         .          .    109:		}
         .          .    110:	}
         .          .    111:	if seg.parent == nil {
         .     4.14us    112:		if b := trimSubsegment(seg); b != nil {
         .          .    113:			outSegments = append(outSegments, b)
         .          .    114:		}
         .          .    115:	}
         .          .    116:	return outSegments
         .          .    117:}
ROUTINE ======================== github.com/aws/aws-xray-sdk-go/xray.packSegments.func1 in /Users/jlindamo/go/pkg/mod/github.com/aws/aws-xray-sdk-go@v1.0.0-rc.11/xray/default_emitter.go
         0     4.14us (flat, cum) 0.00024% of Total
         .          .     96:				break
         .          .     97:			}
         .          .     98:			cb := ss.StreamCompletedSubsegments(s)
         .          .     99:			outSegments = append(outSegments, cb...)
         .          .    100:		}
         .     4.14us    101:		b, _ := json.Marshal(s)
         .          .    102:		return b
         .          .    103:	}
         .          .    104:
         .          .    105:	for _, s := range seg.rawSubsegments {
         .          .    106:		outSegments = packSegments(s, outSegments)
ROUTINE ======================== github.com/cep21/gobenchtraces.xrayRun in /Users/jlindamo/IdeaProjects/gobenchtraces/xray_test.go
         0      1.71s (flat, cum)   100% of Total
         .          .    207:	span.Finish()
         .          .    208:}
         .          .    209:
         .          .    210:func xrayRun(b *testing.B, run benchmarkTracesRun) {
         .          .    211:	ctx := context.Background()
         .    34.30ms    212:	_, s := xray.BeginSegment(ctx, "start")
         .          .    213:	run.toCall()
         .      1.68s    214:	s.Close(nil)
         .          .    215:}
         .          .    216:
         .          .    217:func ddRun(b *testing.B, run benchmarkTracesRun) {
         .          .    218:	span := tracer.StartSpan("test")
         .          .    219:	run.toCall()
(pprof) 
```
