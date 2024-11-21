package logger

import (
	"fmt"
	"runtime"
)

func Error(args ...interface{}) {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`%s %s [%s] <%s> %s`,
		formatText(blink, date),
		formatText(underline, time),
		formatTextExt(bold, red, "ERROR"),
		formatText(cyan, fName),
		msg)
	fmt.Println(content)

}

func Info(args ...interface{}) {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()
	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`%s %s [%s] <%s> %s`,
		formatText(blink, date),
		formatText(underline, time),
		formatTextExt(bold, green, "INFO"),
		formatText(cyan, fName),
		msg)
	fmt.Println(content)
}

func Warning(args ...interface{}) {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)

	content := fmt.Sprintf(`%s %s [%s] <%s> %s`,
		formatText(blink, date),
		formatText(underline, time),
		formatTextExt(bold, yellow, "WARN"),
		formatText(cyan, fName),
		msg)
	fmt.Println(content)
}
