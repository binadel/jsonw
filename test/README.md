# Testing

## Run All Tests
```bash
go test -v ./...
```

## Run Benchmarks
```bash
go test ./test -bench=Benchmark -run=^$ -benchmem
```

## Benchmark Results

```
BenchmarkEncodingJSON_Users-16               547           2182818 ns/op         1456607 B/op       5008 allocs/op
BenchmarkEncodingJSON_Posts-16               159           7470555 ns/op         5821733 B/op      20959 allocs/op
BenchmarkEasyJSON_Users-16                  2926            400893 ns/op          426085 B/op         26 allocs/op
BenchmarkEasyJSON_Posts-16                   922           1280384 ns/op         1412703 B/op         62 allocs/op
BenchmarkJsoniWriter_Users-16               2234            520458 ns/op          426485 B/op         26 allocs/op
BenchmarkJsoniWriter_Posts-16                722           1664082 ns/op         1421110 B/op         62 allocs/op
BenchmarkJsondiWriter_Users-16               948           1291177 ns/op         1679054 B/op      37035 allocs/op
BenchmarkJsondiWriter_Posts-16               280           4269173 ns/op         5908891 B/op     142547 allocs/op
BenchmarkJsondfWriter_Users-16               956           1263253 ns/op         1760587 B/op      37035 allocs/op
BenchmarkJsondfWriter_Posts-16               284           4176842 ns/op         6032322 B/op     142547 allocs/op
BenchmarkJsondsWriter_Users-16               708           1612749 ns/op         5024294 B/op       8043 allocs/op
BenchmarkJsondsWriter_Posts-16               201           6133315 ns/op        19495907 B/op      33015 allocs/op
```