# jsonw

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/binadel/jsonw)](https://goreportcard.com/report/github.com/binadel/jsonw)
![CI](https://github.com/binadel/jsonw/actions/workflows/jsonw.yml/badge.svg)

A lightweight, high-performance Go library for manually building JSON objects and arrays.

## Features

- 🚀 **High Performance**: Built on top of `easyjson` for optimal speed
- 🎯 **Type Safety**: Strongly typed methods for different JSON value types
- 🔧 **Manual Control**: Complete control over JSON structure building
- 🧪 **Well Tested**: Comprehensive test suite with 100% method coverage
- 📊 **Benchmarked**: ~4x faster than standard `json.Marshal`

## Installation

```bash
go get github.com/binadel/jsonw
```

## Quick Start

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    "github.com/binadel/jsonw/pkg"
)

func main() {
    // Create a simple object
    obj := jsonw.NewObjectWriter(nil)
    obj.Open()
    obj.StringField("name", "John Doe")
    obj.IntegerField("age", 30)
    obj.BooleanField("active", true)
    obj.Close()

    result, err := obj.BuildBytes()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(string(result))
    // Output: {"name":"John Doe","age":30,"active":true}
}
```

## Examples

See `examples/basic_usage.go` for comprehensive usage examples.

## License

MIT License - see [LICENSE](LICENSE) file for details.
