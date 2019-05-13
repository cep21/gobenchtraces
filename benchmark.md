```
goos: darwin
goarch: amd64
pkg: github.com/cep21/gobenchtraces
BenchmarkTraces/x-ray-1-8         	  100000	     23053 ns/op	    2732 B/op	      36 allocs/op
BenchmarkTraces/x-ray-1000-8      	  100000	     24625 ns/op	    2730 B/op	      36 allocs/op
BenchmarkTraces/datadog-1-8       	  300000	      3676 ns/op	    1252 B/op	      16 allocs/op
BenchmarkTraces/datadog-1000-8    	 1000000	      2109 ns/op	    1244 B/op	      17 allocs/op
BenchmarkTraces/openjaeger-1-8    	 1000000	      1460 ns/op	     485 B/op	       5 allocs/op
BenchmarkTraces/openjaeger-1000-8 	 2000000	      1005 ns/op	     418 B/op	       4 allocs/op
BenchmarkTraces/newrelic-1-8      	  300000	      6766 ns/op	    1744 B/op	      11 allocs/op
BenchmarkTraces/newrelic-1000-8   	  200000	      8397 ns/op	    1747 B/op	      11 allocs/op
PASS
ok  	github.com/cep21/gobenchtraces	19.983s
```
