# Testing

## Run All Tests
```bash
go test -v ./...
```

## Run Benchmarks
```bash
# Run all benchmarks
go test ./test -bench=Benchmark -run=^$ -benchmem

# Run specific benchmark
go test ./test -bench=BenchmarkObjectWriter_SimpleObject -benchmem
```