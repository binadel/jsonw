# jsonw — High Performance Manual JSON Writing for Go

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/binadel/jsonw)](https://goreportcard.com/report/github.com/binadel/jsonw)
![CI](https://github.com/binadel/jsonw/actions/workflows/test.yml/badge.svg)

`jsonw` is a set of high-performance JSON writing utilities for Go — focused on **zero reflection**, **low allocations**, and **maximum encoding speed**.

It provides:
- an **imperative** writer (`jsoni`) that uses [`easyjson`](https://github.com/mailru/easyjson)'s jwriter under the hood and
- three **declarative** variants (which expose the **same public API**) implemented in different styles:
  - `jsonds` — declarative implementation using **structs**
  - `jsondi` — declarative implementation using **interfaces**
  - `jsondf` — declarative implementation using **functions**

---

## Installing

```bash
go get github.com/binadel/jsonw
```

Import what you need:

```go
import (
    "github.com/binadel/jsonw/jsoni"   // imperative writer
    "github.com/binadel/jsonw/jsondi"  // declarative (interfaces) — same API as jsondf/jsonds
    "github.com/binadel/jsonw/jsondf"  // declarative (functions)  — same API
    "github.com/binadel/jsonw/jsonds"  // declarative (structs)    — same API
)
```

---

## API Overview & Examples

All declarative packages (`jsondi`, `jsondf`, `jsonds`) present the same public API shape
(so switching implementation is straightforward).
jsoni is an imperative builder around `easyjson/jwriter`.

### Imperative

`jsoni` is manual: you call methods to write fields and control exact output. Because it writes directly into a jwriter, it achieves very low allocations.

```go
w := jsoni.NewObjectWriter(nil)
w.Open()
w.StringField("name", "John")
w.IntegerField("age", 30)
w.Close()

out, err := w.BuildBytes()
```

### Declarative

These packages (`jsonds`, `jsondi`, `jsondf`) expose the same declarative API.

```go
obj := json.New(
    json.String("name", "John"),
    json.Integer("age", 30),
)


out, err := obj.Build()
```

See `examples` directory for comprehensive usage examples.

---

## Comparison 

| Implementation         | Speed                   | Memory             | Best Feature         |
|------------------------|-------------------------|--------------------|----------------------|
| `encoding/json`        | 🐢 Slowest              | ❌ High allocs      | Zero setup           |
| `easyjson`             | ⚡ Fastest               | 🧠 Very low allocs | Code generation      |
| `jsoni`                | ⚡ Fast                  | 🧠 Very low allocs | Fully manual control |
| `jsondi/jsondf/jsonds` | 🚀 Faster than std json | ❌ Highest allocs   | Declarative API      |

See `test` directory for benchmark results.

---

## License

MIT License - see [LICENSE](LICENSE) file for details.
