package logger

import (
	"fmt"
	"runtime"
	"strconv"
)

func alert(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`[%s] %s %s (%s:%s)`,
		formatTextExt(bold, blue, " ALERT"),
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
	)
	fmt.Println(content)
	if len(args) > 0 {
		fmt.Println("↳", formatText(bgBlue, formatTextExt(bold, brightYellow, msg)))
	}

}
func debug(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`[%s] %s %s (%s:%s)`,
		formatTextExt(bold, blue, " LOGGER DEBUG"),
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
	)
	fmt.Println(content)
	if len(args) > 0 {
		fmt.Println("↳", formatText(bgBlue, formatTextExt(bold, brightYellow, msg)))
	}

}
func Error(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`[%s] %s %s (%s:%s)`,
		formatTextExt(bold, red, " ERROR"),
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
	)
	fmt.Println(content)
	if len(args) > 0 {
		fmt.Println("↳", formatText(bgRed, formatTextExt(bold, brightYellow, msg)))
	}

}

func Info(args ...interface{}) {

	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`[%s] %s %s %s`,
		formatTextExt(bold, brightGreen, " INFO "),
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatText(bold, msg),
	)
	fmt.Println(content)

}

func InfoC(args ...interface{}) {

	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`[%s] %s %s (%s:%s)`,
		formatTextExt(bold, brightGreen, " INFO "),
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
	)
	fmt.Println(content)
	if len(args) > 0 {
		fmt.Println("↳", formatTextExt(bold, brightYellow, msg))
	}
}
func Warn(args ...interface{}) {

	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`[%s] %s %s %s`,
		formatTextExt(bold, orange, " WARN "),
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatText(bold, msg),
	)
	fmt.Println(content)
}

func WarnC(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`[%s] %s %s (%s:%s)`,
		formatTextExt(bold, orange, " WARN "),
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
	)
	fmt.Println(content)
	if len(args) > 0 {
		fmt.Println("↳", formatTextExt(bold, brightYellow, msg))
	}

}
func Debug(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`[%s] %s %s (%s:%s)`,
		formatTextExt(bold, magenta, " DBUG "),
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
	)
	fmt.Println(content)
	if len(args) > 0 {
		fmt.Println("↳", formatTextExt(bold, brightYellow, msg))
	}
}
