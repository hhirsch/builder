package helpers

import (
	"fmt"
	charmLog "github.com/charmbracelet/log"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"
)

type Logger struct {
	file          *os.File
	logOperations map[string]LogOperation
}

type LogOperation func(msg interface{}, keyvals ...interface{})

func NewLogger(fileName string) *Logger {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file for writing:", err)
	}
	logOperations := map[string]LogOperation{
		"DEBUG": charmLog.Debug,
		"INFO":  charmLog.Info,
		"WARN":  charmLog.Warn,
		"ERROR": charmLog.Error,
		"FATAL": charmLog.Fatal,
	}

	charmLog.SetLevel(charmLog.DebugLevel)
	return &Logger{
		file:          file,
		logOperations: logOperations,
	}
}

func (logger *Logger) Print(string string) {
	logger.print(string)
}

func (logger *Logger) writeFile(level string, message string) {
	time := time.Now()
	logger.writeLog(time.Format("2006/01/02 15:04:05") + " " + level + " " + message + "\n")
}

func (logger *Logger) splitNamespaceFromFunctionName(functionName string) (isolatedNamespaceName string, isolatedFunctionName string) {
	lastSlash := strings.LastIndex(functionName, "/")
	if lastSlash == -1 {
		return "", functionName
	}
	lastDot := strings.LastIndex(functionName[lastSlash:], ".")
	if lastDot == -1 {
		return functionName[:lastSlash], functionName[lastSlash+1:]
	}

	return functionName[:lastSlash], functionName[lastSlash+lastDot+1:]
}

func (logger *Logger) getCallerName() (result string) {
	callAdresses := make([]uintptr, 10)
	callers := runtime.Callers(4, callAdresses)
	result = ""
	if callers == 0 {
		return
	}

	function := runtime.FuncForPC(callAdresses[0])
	if function == nil {
		return
	}

	fullName := function.Name()
	parts := strings.Split(fullName, ".")

	if len(parts) == 0 {
		return
	}
	expression := regexp.MustCompile(`[()*]`)
	callerTypeWithNamespace := parts[len(parts)-2]
	callerType := expression.ReplaceAllString(callerTypeWithNamespace, "")
	if callerType == callerTypeWithNamespace {
		_, callerType = logger.splitNamespaceFromFunctionName(callerTypeWithNamespace)
	}
	return callerType + "->" + parts[len(parts)-1]
}

func (logger *Logger) genericLog(module string, logLevel string, message interface{}, data ...interface{}) {
	charmLog.SetPrefix("")
	if module != "" {
		charmLog.SetPrefix(module)
		logger.writeFile(logLevel, module+": "+message.(string))
	} else {
		logger.writeFile(logLevel, message.(string))
	}
	if operation, exists := logger.logOperations[logLevel]; exists {
		if len(data) >= 2 {
			operation(message)
		} else {
			operation(message)
		}
	}
	charmLog.SetPrefix("")
}

func (logger *Logger) Debug(message interface{}, data ...interface{}) {
	logger.genericLog(logger.getCallerName(), "DEBUG", message, data)
}

func (logger *Logger) Debugf(message string, data ...interface{}) {
	logger.Debug(fmt.Sprintf(message, data...))
}

func (logger *Logger) Info(message interface{}, data ...interface{}) {
	logger.genericLog(logger.getCallerName(), "INFO", message, data)
}

func (logger *Logger) Infof(message string, data ...interface{}) {
	logger.Info(fmt.Sprintf(message, data...))
}

func (logger *Logger) Warn(message interface{}, data ...interface{}) {
	logger.genericLog(logger.getCallerName(), "WARN", message, data)
}

func (logger *Logger) Error(message interface{}, data ...interface{}) {
	logger.genericLog(logger.getCallerName(), "ERROR", message, data)
}

func (logger *Logger) Errorf(message string, data ...interface{}) {
	logger.Error(fmt.Sprintf(message, data...))
}

func (logger *Logger) Fatal(message interface{}, data ...interface{}) {
	logger.genericLog(logger.getCallerName(), "FATAL", message, data)
}

func (logger *Logger) Fatalf(message string, data ...interface{}) {
	logger.Fatal(fmt.Sprintf(message, data...))
}

func (logger *Logger) writeLog(string string) {
	_, err := logger.file.WriteString(string)
	if err != nil {
		panic(err)
	}
}

func (logger *Logger) print(string string) {
	fmt.Println(string)
	logger.writeLog(string)
}
