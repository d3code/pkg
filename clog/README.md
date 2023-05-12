# Console Logging

## Introduction

This is a simple logging library for go. It is designed to be simple to use and provide convenience functions for logging to the console.

It is opinionated in the way it logs to the console, and is not designed to be a full logging library. It is designed to be used in small projects where you want to log to the console with a little more flair than using `fmt.Println()` or the like.

```go
// TODO: be more configurable and performant
```


## Usage

To use this library, simply install it in your project:

```shell
go get github.com/d3code/pkg/clog
```

Then, you can use the `clog` package to log to the console:

```go
import "github.com/d3code/pkg/clog"

func main() {
    clog.Info("Hello World!")
}
```
