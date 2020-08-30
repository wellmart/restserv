# Rest Server

[![Build Status](https://travis-ci.org/wellmart/restserv.svg?branch=master)](https://travis-ci.org/wellmart/restserv)
[![Go Report Card](https://goreportcard.com/badge/github.com/wellmart/restserv)](https://goreportcard.com/report/github.com/wellmart/restserv)
[![Coverage Status](https://coveralls.io/repos/github/wellmart/restserv/badge.svg?branch=master)](https://coveralls.io/github/wellmart/restserv?branch=master)
![Version](https://img.shields.io/badge/version-0.1.0-blue)
[![Software License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![GoDoc](https://godoc.org/github.com/wellmart/restserv?status.svg)](https://godoc.org/github.com/wellmart/restserv)

A simplest way to write REST API services with http routing. It's ideal for writing simple, performant backend REST API services.

## Requirements

Go 1.1 and beyond.

## Installation

Use the go package manager to install Restserv.

```bash
go get github.com/wellmart/restserv
```

## Usage

```go
package main

import "github.com/wellmart/restserv"

func main() {
    app := restserv.New()

    app.Get("/api/test/", func(c *restserv.Context) {
        type test struct {
            Hello string
        }

        c.WriteJSON(&test{
            Hello: "Hello World"})
    })

    app.ListenAndServe("127.0.0.1:8080")
}
```

## Staying up to date

To update Restserv to the latest version, use `go get -u github.com/wellmart/restserv`.

## License

[MIT](https://choosealicense.com/licenses/mit/)
