```
goos: darwin
goarch: amd64
pkg: github.com/cep21/gobenchtraces
BenchmarkTraces/x-ray-1-8         	   50000	     22314 ns/op	    2728 B/op	      36 allocs/op
BenchmarkTraces/x-ray-1000-8      	  100000	     24902 ns/op	    2741 B/op	      36 allocs/op
BenchmarkTraces/datadog-1-8       	  500000	      3590 ns/op	    1242 B/op	      16 allocs/op
BenchmarkTraces/datadog-1000-8    	 1000000	      2112 ns/op	    1248 B/op	      17 allocs/op
BenchmarkTraces/openjaeger-1-8    	 1000000	      1124 ns/op	     485 B/op	       5 allocs/op
BenchmarkTraces/openjaeger-1000-8 	 3000000	       624 ns/op	     416 B/op	       4 allocs/op
BenchmarkTraces/newrelic-1-8      	  500000	      5018 ns/op	    1744 B/op	      11 allocs/op
BenchmarkTraces/newrelic-1000-8   	  300000	      6203 ns/op	    1746 B/op	      11 allocs/op
PASS
ok  	github.com/cep21/gobenchtraces	18.186s
```
