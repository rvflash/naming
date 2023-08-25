# Naming conventions

[![GoDoc](https://godoc.org/github.com/rvflash/naming?status.svg)](https://godoc.org/github.com/rvflash/naming)
[![Build Status](https://github.com/rvflash/naming/workflows/build/badge.svg)](https://github.com/rvflash/naming/actions?workflow=build)
[![Code Coverage](https://codecov.io/gh/rvflash/naming/branch/master/graph/badge.svg)](https://codecov.io/gh/rvflash/naming)
[![Go Report Card](https://goreportcard.com/badge/github.com/rvflash/naming?)](https://goreportcard.com/report/github.com/rvflash/naming)


The package `naming` provides methods to parse and build compound names for variables types, 
functions, classes or other structures in source code.

Naming conventions bring consistency, better understanding, readability and automation.

List of supported multiple-word identifier formats:

* camelCase
* CONSTANT_CASE
* flatcase
* kebab-case
* PascalCase
* snake_case
* Train-Case
* UPPERFLATCASE


### Installation

To install it, you need to install Go and set your Go workspace first.
Then, download and install it:

```bash
$ go get -u github.com/rvflash/naming
```    
Import it in your code:

```go
import "github.com/rvflash/naming"
```

### Prerequisite

`naming` uses the Go modules that required Go 1.11 or later.