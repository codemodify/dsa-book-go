package main

//
// The Go way SOLUTION
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
// This solution is leveraging Go modern principles and implements
// Factory based on that.
//

type Logger interface {
	LogInfo(string)
	LogWarning(string)
	LogError(string)
}

type consoleLogger struct{}

func (l consoleLogger) LogInfo(string)    {}
func (l consoleLogger) LogWarning(string) {}
func (l consoleLogger) LogError(string)   {}

type fileLogger struct{}

func (l fileLogger) LogInfo(string)    {}
func (l fileLogger) LogWarning(string) {}
func (l fileLogger) LogError(string)   {}

type emptyLogger struct{}

func (l emptyLogger) LogInfo(string)    {}
func (l emptyLogger) LogWarning(string) {}
func (l emptyLogger) LogError(string)   {}

type loggerType int

const (
	loggerTypeConsole loggerType = iota
	loggerTypeFile
)

func createLogger(lt loggerType) Logger {
	switch lt {
	case loggerTypeConsole:
		return &consoleLogger{}
	case loggerTypeFile:
		return &fileLogger{}
	}

	return &emptyLogger{}
}

func main() {
	logger := createLogger(loggerTypeConsole)
	logger.LogInfo("test")
}
