package logger

type Sender interface {
	Send(l *Logger)
}

func (e Email) Send(l *Logger) {
	if l.c.IsDebugMode {
		debug("Sent email", l.c.Email)
	}
}
