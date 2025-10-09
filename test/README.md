# Testing

## Run Unit Tests
```bash
go test ./pkg -v
```

## Run Integration Tests
```bash
go test ./test -v
```

## Run All Tests
```bash
go test ./...
```

## Run Benchmarks
```bash
# Run all benchmarks
go test ./test -bench=Benchmark -run=^$ -benchmem

# Run specific benchmark
go test ./test -bench=BenchmarkObjectWriter_SimpleObject -benchmem
```