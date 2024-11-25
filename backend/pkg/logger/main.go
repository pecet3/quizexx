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
	content := fmt.Sprintf(`[%s] %s %s (%s:%s)`,
		formatTextExt(bold, red, "ERROR"),
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
	)
	fmt.Println(content)
	if len(args) > 0 {
		fmt.Println("↳", formatTextExt(bold, yellow, msg))
	}
}

func Info(args ...interface{}) {
	date := getCurrentDate()
	time := getCurrentTime()
	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`%s %s [%s] %s`,
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatTextExt(bold, green, "INFO"),
		formatText(italic, msg))
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
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatTextExt(bold, green, "INFO"),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
		formatText(italic, msg))
	fmt.Println(content)
}
func Warn(args ...interface{}) {
	date := getCurrentDate()
	time := getCurrentTime()
	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`%s %s [%s] %s`,
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatTextExt(bold, orange, "WARN"),
		formatText(bold, msg))
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
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatTextExt(bold, orange, "WARN"),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
		formatText(bold, msg))
	fmt.Println(content)

}
func Debug(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`[%s] %s %s (%s:%s)`,
		formatTextExt(bold, blue, "DEBUG"),
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
	)
	fmt.Println(content)
	if len(args) > 0 {
		fmt.Println("↳", formatTextExt(bold, yellow, msg))
	}
}
func DebugT(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`%s %s [%s] (%s:%s) %s`,
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatTextExt(bold, orange, "DEBUG"),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
		formatText(italic, msg))
	fmt.Println(content)
}
