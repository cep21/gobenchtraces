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

bench_out:
	echo "\`\`\`" > benchmark.md
	go test -v -benchmem -run=^$$ -bench=. ./... >> benchmark.md
	echo "\`\`\`" >> benchmark.md

bench_mem_xray:
	go test -benchmem -run=^$$ -bench="BenchmarkTraces/x-ray-1000" -memprofile=xray-mem.profile
	echo "\`\`\`" > benchmark_xray_mem.md
	echo "list xray" | go tool pprof -sample_index=alloc_objects gobenchtraces.test xray-mem.profile >> benchmark_xray_mem.md
	echo >> benchmark_xray_mem.md
	echo "\`\`\`" >> benchmark_xray_mem.md

profile_cpu:
	go test -run=^$$ -bench=$(PROFILE_RUN) -benchtime=15s -cpuprofile=cpu.profile ./benchmarking
	go tool pprof benchmarking.test cpu.profile
	rm -f cpu.profile benchmarking.test

bench_blocking_xray:
	go test -benchmem -run=^$$ -bench="BenchmarkTraces/x-ray-1000" -blockprofile=block_xray.profile
	echo "\`\`\`" > benchmark_xray_block.md
	echo "list xray" | go tool pprof gobenchtraces.test block_xray.profile >> benchmark_xray_block.md
	echo >> benchmark_xray_block.md
	echo "\`\`\`" >> benchmark_xray_block.md

# Lint the code
lint:
	golangci-lint run

# ci installs dep by direct version.  Users install with 'go get'
setup_ci:
	GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint
