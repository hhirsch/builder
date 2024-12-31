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

func (this *Logger) Print(string string) {
	this.print(string)
}

func (this *Logger) writeFile(level string, message string) {
	time := time.Now()
	this.writeLog(time.Format("2006/01/02 15:04:05") + " " + level + " " + message + "\n")
}

func (this *Logger) getCallerName() string {
	callAdresses := make([]uintptr, 10)
	callers := runtime.Callers(3, callAdresses)

	if callers == 0 {
		return ""
	}

	function := runtime.FuncForPC(callAdresses[0])
	if function == nil {
		return ""
	}

	fullName := function.Name()
	parts := strings.Split(fullName, ".")

	if len(parts) == 0 {
		return ""
	}
	expression := regexp.MustCompile(`[()*]`)
	callerType := expression.ReplaceAllString(parts[len(parts)-2], "")

	return callerType + "->" + parts[len(parts)-1]
}

func (this *Logger) genericLog(module string, logLevel string, message interface{}, data ...interface{}) {
	charmLog.SetPrefix("")
	if module != "" {
		charmLog.SetPrefix(module)
		this.writeFile(logLevel, module+": "+message.(string))
	} else {
		this.writeFile(logLevel, message.(string))
	}
	if operation, exists := this.logOperations[logLevel]; exists {
		if len(data) >= 2 {
			operation(message)
		} else {
			operation(message)
		}
	}
	charmLog.SetPrefix("")
}

func (this *Logger) Debug(message interface{}, data ...interface{}) {
	this.genericLog(this.getCallerName(), "DEBUG", message, data)
}

func (this *Logger) Debugf(message string, data ...interface{}) {
	this.Debug(fmt.Sprintf(message, data...))
}

func (this *Logger) Info(message interface{}, data ...interface{}) {
	this.genericLog(this.getCallerName(), "INFO", message, data)
}

func (this *Logger) Infof(message string, data ...interface{}) {
	this.Info(fmt.Sprintf(message, data...))
}

func (this *Logger) Warn(message interface{}, data ...interface{}) {
	this.genericLog(this.getCallerName(), "WARN", message, data)
}

func (this *Logger) Error(message interface{}, data ...interface{}) {
	this.genericLog(this.getCallerName(), "ERROR", message, data)
}

func (this *Logger) Fatal(message interface{}, data ...interface{}) {
	this.genericLog(this.getCallerName(), "FATAL", message, data)
}

func (this *Logger) Fatalf(message string, data ...interface{}) {
	this.Fatal(fmt.Sprintf(message, data...))
}

func (this *Logger) writeLog(string string) {
	_, err := this.file.WriteString(string)
	if err != nil {
		panic(err)
	}
}

func (this *Logger) print(string string) {
	fmt.Println(string)
	this.writeLog(string)
}
