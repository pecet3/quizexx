package logger

import (
	"fmt"
	"runtime"
	"strconv"
)

func Error(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`%s %s [%s] (%s:%s) %s`,
		formatText(italic, date),
		formatText(underline, time),
		formatTextExt(bold, red, "ERROR"),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
		msg)
	fmt.Println(content)

}

func Info(args ...interface{}) {
	date := getCurrentDate()
	time := getCurrentTime()
	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`%s %s [%s] %s`,
		formatText(italic, date),
		formatText(underline, time),
		formatTextExt(bold, green, "INFO"),
		msg)
	fmt.Println(content)
}

func InfoC(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`%s %s [%s] (%s:%s) %s`,
		formatText(italic, date),
		formatText(underline, time),
		formatTextExt(bold, green, "INFO"),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
		msg)
	fmt.Println(content)
}
func Warn(args ...interface{}) {
	date := getCurrentDate()
	time := getCurrentTime()
	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`%s %s [%s] %s`,
		formatText(italic, date),
		formatText(underline, time),
		formatTextExt(bold, orange, "WARN"),
		msg)
	fmt.Println(content)
}

func WarnC(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`%s %s [%s] (%s:%s) %s`,
		formatText(italic, date),
		formatText(underline, time),
		formatTextExt(bold, yellow, "WARN"),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
		msg)
	fmt.Println(content)

}

func Debug(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`%s %s [%s] (%s:%s) %s`,
		formatText(italic, date),
		formatText(underline, time),
		formatTextExt(bold, orange, "DEBUG"),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
		msg)
	fmt.Println(content)
}
