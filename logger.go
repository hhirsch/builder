package main

import (
	"fmt"
	"log"
	charmLog "github.com/charmbracelet/log"
	"os"
	"time"
)

type Logger struct {
	file *os.File
}

func NewLogger() *Logger {
	file, err := os.OpenFile("builder.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	
	return &Logger{
		file: file,
	}
}

func (this *Logger) Print(string string) {
	this.print(string)
}

func (this *Logger) writeGeneric(level string, string string) {
	time := time.Now()
	this.writeLog(time.Format("2006/01/02 15:04:05") + " " + level + " " + string + "\n")
}

func (this *Logger) Info(string string) {
	this.writeGeneric("INFO", string)
	charmLog.Info(string)
}

func (this *Logger) Warn(string string) {
	this.writeGeneric("WARN", string)	
	charmLog.Warn(string)
}

func (this *Logger) Fatal(string string) {
	this.writeGeneric("FATAL", string)
	log.Fatal(string)
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
