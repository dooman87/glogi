# glogi

> Golang Log Interface

This is a small package that provides an interface for logger. The idea is to have logger 
interface and reuse it in the libraries, also providing an override point so applications
that using libraries can provide their own implementation.

The interface provides only two levels of logging: debug (Print* functions) and error.

Partially inspired by Dave Chaney's [article about log levels](https://dave.cheney.net/2015/11/05/lets-talk-about-logging).

## Install

```
go get github.com/dooman87/glogi
```

## Usage

You can use this package as a dependency or just copy-paste interface of a logger:

```go
package log

//Logger interface that provides only two log levels:
// * Print - is for all non-error messages
// * Error - for all errors
type Logger interface {
	Printf(format string, v ...interface{})
	Print(v ...interface{})
	Errorf(format string, v ...interface{})
	Error(v ...interface{})
}
```

## Contribute

PRs accepted.

## License

MIT © Dmitry Pokidov