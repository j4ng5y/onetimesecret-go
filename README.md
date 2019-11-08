[![Go Report Card](https://goreportcard.com/badge/github.com/j4ng5y/onetimesecret-go)](https://goreportcard.com/report/github.com/j4ng5y/onetimesecret-go) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/j4ng5y/onetimesecret-go) [![GoDoc](https://godoc.org/github.com/j4ng5y/onetimesecret-go?status.svg)](https://godoc.org/github.com/j4ng5y/onetimesecret-go)

# OneTimeSecret Go Client Library

This is a fairly straight-forward Go implementation of the https://onetimesecret.com service for Go applications

## Installation

`go get -u github.com/j4ng5y/onetimesecret-go`

```go
package main

import "github.com/j4ng5y/onetimesecret-go"

func main() {
    c := onetimesecret.New(&onetimesecret.Credentials{Username:"j4ng5y",APIToken:"abcdefg12345"})
    ...
}
```

## Examples

Please see the [examples](/examples) folder.

## Note

If you want to use a custom http client, instead of using the `onetimesecret.New()` function to generate the client, use the `onetimesecret.NewWithOptions()` function.