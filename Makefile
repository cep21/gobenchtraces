 build:
	go build ./...

# Run unit tests
test:
	env "GORACE=halt_on_error=1" go test -v -race ./...

# Format the code
fix:
	find . -iname '*.go' -not -path '*/vendor/*' -print0 | xargs -0 gofmt -s -w
	find . -iname '*.go' -not -path '*/vendor/*' -print0 | xargs -0 goimports -w

bench:
	go test -v -benchmem -run=^$$ -bench=. ./...

bench_all:
	echo "\`\`\`" > benchmark.md
	go test -v -benchmem -run=^$$ -bench=. ./... >> benchmark.md
	echo >> benchmark.md
	echo "\`\`\`" >> benchmark.md

bench_xray:
	go test -benchmem -run=^$$ -bench="BenchmarkTraces/x-ray-1000" -memprofile=xray_mem.profile -cpuprofile=xray_cpu.profile -mutexprofile=xray_mutex.profile -blockprofile=xray_block.profile

	echo "\`\`\`go" > benchmark_xray_mem.md
	echo >> benchmark_xray_mem.md
	echo "top" | go tool pprof -sample_index=alloc_objects gobenchtraces.test xray_mem.profile >> benchmark_xray_mem.md
	echo >> benchmark_xray_mem.md
	echo "top -cum" | go tool pprof -sample_index=alloc_objects gobenchtraces.test xray_mem.profile >> benchmark_xray_mem.md
	echo >> benchmark_xray_mem.md
	echo "list xray" | go tool pprof -sample_index=alloc_objects gobenchtraces.test xray_mem.profile >> benchmark_xray_mem.md
	echo >> benchmark_xray_mem.md
	echo "\`\`\`" >> benchmark_xray_mem.md

	echo "\`\`\`go" > benchmark_xray_cpu.md

	echo >> benchmark_xray_cpu.md
	echo "top" | go tool pprof gobenchtraces.test xray_cpu.profile >> benchmark_xray_cpu.md
	echo >> benchmark_xray_cpu.md
	echo "top -cum" | go tool pprof gobenchtraces.test xray_cpu.profile >> benchmark_xray_cpu.md
	echo >> benchmark_xray_cpu.md
	echo "list xray" | go tool pprof gobenchtraces.test xray_cpu.profile >> benchmark_xray_cpu.md
	echo >> benchmark_xray_cpu.md
	echo "\`\`\`" >> benchmark_xray_cpu.md

	echo "\`\`\`go" > benchmark_xray_block.md
	echo >> benchmark_xray_block.md
	echo "top" | go tool pprof gobenchtraces.test xray_block.profile >> benchmark_xray_block.md
	echo >> benchmark_xray_block.md
	echo "top -cum" | go tool pprof gobenchtraces.test xray_block.profile >> benchmark_xray_block.md
	echo >> benchmark_xray_block.md
	echo "list xray" | go tool pprof gobenchtraces.test xray_block.profile >> benchmark_xray_block.md
	echo >> benchmark_xray_block.md
	echo "\`\`\`" >> benchmark_xray_block.md

	echo "\`\`\`go" > benchmark_xray_mutex.md
	echo >> benchmark_xray_mutex.md
	echo "top" | go tool pprof gobenchtraces.test xray_mutex.profile >> benchmark_xray_mutex.md
	echo >> benchmark_xray_mutex.md
	echo "top -cum" | go tool pprof gobenchtraces.test xray_mutex.profile >> benchmark_xray_mutex.md
	echo >> benchmark_xray_mutex.md
	echo "list xray" | go tool pprof gobenchtraces.test xray_mutex.profile >> benchmark_xray_mutex.md
	echo >> benchmark_xray_mutex.md
	echo "\`\`\`" >> benchmark_xray_mutex.md

clean:
	rm -f *.profile gobenchtraces.test

generate: bench_all bench_xray clean

# Lint the code
lint:
	golangci-lint run

# ci installs dep by direct version.  Users install with 'go get'
setup_ci:
	GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint
