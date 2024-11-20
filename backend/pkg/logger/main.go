package logger

import "log"

func Error(where string, err error) {
	log.Printf(`[ERROR] <%s>  %d\n`, where, err)
}
