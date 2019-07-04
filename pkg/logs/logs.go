package logs

import (
	"io"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init() {
	/*	Warnings	*/
	traceFile, err := os.OpenFile("pkg/logs/trace.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", ":", err)
	}
	multiTrace := io.MultiWriter(traceFile, os.Stdout)

	Trace = log.New(multiTrace,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	/*	Info	*/
	infoFile, err := os.OpenFile("pkg/logs/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", ":", err)
	}
	multiInfo := io.MultiWriter(infoFile, os.Stdout)
	Info = log.New(multiInfo,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	/*	Warnings	*/
	warningsFile, err := os.OpenFile("pkg/logs/warnings.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", ":", err)
	}
	multiWarnings := io.MultiWriter(warningsFile, os.Stdout)
	Warning = log.New(multiWarnings,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	/*	Error	*/
	errorFile, err := os.OpenFile("pkg/logs/errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", ":", err)
	}
	multiError := io.MultiWriter(errorFile, os.Stderr)

	Error = log.New(multiError,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
