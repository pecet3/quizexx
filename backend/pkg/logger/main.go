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
	SMTPHost      string
	SMTPPort      int
	Username      string
	Password      string
	FromAddress   string
	ToAddresses   []string
	SubjectRaport string
	SubjectAlert  string
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
					err := m.SendLogs(ctx, l)
					if c.IsDebugMode {
						if err != nil {
							debug("sending logs err: ", err)
							return
						}
						debug("sending logs successful")
					}
				}(method)
			}
			wg.Wait()
			l.cleanCache()
			if c.IsDebugMode {
				debug("cleaned cache")
			}
		}
	}()

	return l
}
func (l *Logger) Alert(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	currentTime := getCurrentTime()

	msg := fmt.Sprint(args...)
	contentRaw := fmt.Sprintf(`[%s] %s %s (%s:%s) %s`,
		" ALERT",
		date,
		currentTime,
		fName,
		strconv.Itoa(line),
		msg,
	)
	l.addCache(time.Now(), contentRaw)
	alert(msg)

	wg := sync.WaitGroup{}
	for _, method := range l.senders {
		wg.Add(1)
		ctx, cancel := context.WithCancel(context.Background())
		go func(m Sender) {
			defer cancel()
			defer wg.Done()
			err := m.SendAlert(ctx, l, msg)
			if l.c.IsDebugMode {
				if err != nil {
					debug("sending alert err: ", err)
					return
				}
				debug("sending alert successful")
			}
		}(method)
	}
	wg.Wait()

}

func (l *Logger) Error(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	currentTime := getCurrentTime()

	msg := fmt.Sprint(args...)
	contentRaw := fmt.Sprintf(`[%s] %s %s (%s:%s) %s`,
		" ERROR",
		date,
		currentTime,
		fName,
		strconv.Itoa(line),
		msg,
	)
	l.addCache(time.Now(), contentRaw)
	Error(msg)
}

func (l *Logger) Info(args ...interface{}) {
	date := getCurrentDate()
	currentTime := getCurrentTime()

	msg := fmt.Sprint(args...)
	contentRaw := fmt.Sprintf(`[%s] %s %s  %s`,
		" INFO ",
		date,
		currentTime,

		msg,
	)
	l.addCache(time.Now(), contentRaw)
	Info(msg)
}

func (l *Logger) Warn(args ...interface{}) {
	date := getCurrentDate()
	currentTime := getCurrentTime()

	msg := fmt.Sprint(args...)
	contentRaw := fmt.Sprintf(`[%s] %s %s  %s`,
		" WARN ",
		date,
		currentTime,

		msg,
	)
	l.addCache(time.Now(), contentRaw)
	Warn(msg)
}

func (l *Logger) Debug(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	currentTime := getCurrentTime()

	msg := fmt.Sprint(args...)
	contentRaw := fmt.Sprintf(`[%s] %s %s (%s:%s) %s`,
		" DBUG ",
		date,
		currentTime,
		fName,
		strconv.Itoa(line),
		msg,
	)
	l.addCache(time.Now(), contentRaw)
	Debug(msg)
}

func (l *Logger) InfoC(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	currentTime := getCurrentTime()

	msg := fmt.Sprint(args...)
	contentRaw := fmt.Sprintf(`[%s] %s %s (%s:%s) %s`,
		" INFO ",
		date,
		currentTime,
		fName,
		strconv.Itoa(line),
		msg,
	)
	l.addCache(time.Now(), contentRaw)
	InfoC(msg)
}

func (l *Logger) WarnC(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	currentTime := getCurrentTime()

	msg := fmt.Sprint(args...)
	contentRaw := fmt.Sprintf(`[%s] %s %s (%s:%s) %s`,
		" WARN ",
		date,
		currentTime,
		fName,
		strconv.Itoa(line),
		msg,
	)
	l.addCache(time.Now(), contentRaw)
	WarnC(msg)
}
