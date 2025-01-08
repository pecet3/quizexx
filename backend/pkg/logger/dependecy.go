package logger

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Email struct {
	SMTPHost     string
	SMTPPort     int
	Username     string
	Password     string
	FromAddress  string
	ToAddresses  []string
	Subject      string
	BodyTemplate string
}
type Config struct {
	IsDebugMode bool
	Email       *Email
	Duration    time.Duration
}

type Logger struct {
	cache map[time.Time]string
	cMu   sync.Mutex

	senders map[string]Sender

	// config
	c *Config
}

func New(c *Config) *Logger {
	l := &Logger{
		cache:   make(map[time.Time]string),
		c:       c,
		senders: make(map[string]Sender),
	}
	if c.Email != nil {
		l.senders["email"] = c.Email
	}
	go func() {
		for {
			time.Sleep(c.Duration)
			wg := sync.WaitGroup{}
			for _, method := range l.senders {
				wg.Add(1)
				ctx, cancel := context.WithCancel(context.Background())
				go func(m Sender) {
					defer cancel()
					defer wg.Done()
					err := m.Send(ctx, l)
					if c.IsDebugMode {
						if err != nil {
							debug("sending email err: ", err)
							return
						}
						debug("sent email successful", err)
					}
				}(method)
			}
			wg.Wait()
			l.clearCache()
			if c.IsDebugMode {
				debug("clear cache")
			}
		}
	}()

	return l
}

func (l *Logger) Error(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	currentTime := getCurrentTime()

	msg := fmt.Sprint(args...)
	contentRaw := fmt.Sprintf(`[%s] %s %s (%s:%s)`,
		" ERROR",
		date,
		currentTime,
		fName,
		strconv.Itoa(line),
	)
	l.addCache(time.Now(), contentRaw)
	Error(msg)
}

func (l *Logger) Info(args ...interface{}) {
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

func (l *Logger) InfoC(args ...interface{}) {
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
func (l *Logger) Warn(args ...interface{}) {

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

func (l *Logger) WarnC(args ...interface{}) {
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
func (l *Logger) Debug(args ...interface{}) {
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
