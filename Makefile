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

bench_mem:
	go test -benchmem -run=^$$ -bench="BenchmarkTraces/openjaeger-1000" -memprofile=mem.out
	go tool pprof -sample_index=alloc_objects -list xray -call_tree -cum gobenchtraces.test mem.out > xray_mem.txt

profile_cpu:
	go test -run=^$$ -bench=$(PROFILE_RUN) -benchtime=15s -cpuprofile=cpu.out ./benchmarking
	go tool pprof benchmarking.test cpu.out
	rm -f cpu.out benchmarking.test

profile_blocking:
	go test -benchmem -run=^$$ -bench="BenchmarkTraces/x-ray-1000" -blockprofile=block_xray.out
	go tool pprof -list xray -call_tree -cum -nodecount=30 gobenchtraces.test block_xray.out
	rm -f block.out benchmarking.test


# Lint the code
lint:
	golangci-lint run

# ci installs dep by direct version.  Users install with 'go get'
setup_ci:
	GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint
