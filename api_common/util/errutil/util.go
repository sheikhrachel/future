package errutil

import (
	"log"
	"runtime"
)

// HandleError is a util to provide readable insight into errors
func HandleError(err error) (b bool) {
	if err != nil {
		/*
			0 = this func.
			use 1 to log where the error actually happened
		*/
		_, filename, line, _ := runtime.Caller(1)
		log.Printf("[error] %s:%d %v", filename, line, err)
		b = true
	}
	return
}
