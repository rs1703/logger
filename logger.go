package logger

import (
	"io"
	"log"
	"os"
	"runtime"
	"time"
)

var Err = log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
var Inf = log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime|log.Lshortfile)
var track = log.New(os.Stdout, "DEBUG:\t", log.Ldate|log.Ltime)

var file *os.File
var out io.Writer

func SetOutput(fileName string) {
	if file != nil {
		file.Close()
	}

	var err error
	file, err = os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err.Error())
	}

	out = io.MultiWriter(os.Stdout, file)
	Err.SetOutput(out)
	Inf.SetOutput(out)
	track.SetOutput(out)
}

func Track() func() {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)

	now := time.Now()
	track.Println(fn.Name())
	return func() {
		track.Println(fn.Name(), time.Since(now))
	}
}
