package logger

import (
	"log"
	"runtime"
)

func Error(err error) {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	log.Printf(`[ERROR] <%s> %s`, fName, err)
}
