package logger

import "time"

func (l *Logger) addCache(t time.Time, content string) {
	l.cMu.Lock()
	defer l.cMu.Unlock()
	l.cache[t] = content
}

func (l *Logger) cleanCache() {
	l.cMu.Lock()
	defer l.cMu.Unlock()
	for key, _ := range l.cache {
		delete(l.cache, key)
	}
}
