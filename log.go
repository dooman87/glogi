package glogi

//Logger interface that provides only two log levels:
// * Print - is for all non-error messages
// * Error - for all errors
type Logger interface {
	Printf(format string, v ...interface{})
	Print(v ...interface{})
	Errorf(format string, v ...interface{})
	Error(v ...interface{})
}