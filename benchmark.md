goos: darwin
goarch: amd64
pkg: github.com/cep21/gobenchtraces
BenchmarkTraces/x-ray-1-8         	  100000	     22084 ns/op	    2732 B/op	      36 allocs/op
BenchmarkTraces/x-ray-1000-8      	  100000	     23159 ns/op	    2729 B/op	      36 allocs/op
BenchmarkTraces/datadog-1-8       	  500000	      3387 ns/op	    1243 B/op	      16 allocs/op
BenchmarkTraces/datadog-1000-8    	 1000000	      2035 ns/op	    1241 B/op	      17 allocs/op
BenchmarkTraces/openjaeger-1-8    	 1000000	      1013 ns/op	     482 B/op	       5 allocs/op
BenchmarkTraces/openjaeger-1000-8 	 3000000	       594 ns/op	     417 B/op	       4 allocs/op
BenchmarkTraces/newrelic-1-8      	  500000	      4813 ns/op	    1744 B/op	      11 allocs/op
BenchmarkTraces/newrelic-1000-8   	  300000	      6174 ns/op	    1746 B/op	      11 allocs/op
PASS
ok  	github.com/cep21/gobenchtraces	19.171s
