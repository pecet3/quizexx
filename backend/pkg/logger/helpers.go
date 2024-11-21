package logger

import (
	"fmt"
	"time"
)

func getCurrentDate() string {
	now := time.Now()
	return now.Format("2006/01/02")
}
func getCurrentTime() string {
	now := time.Now()
	return now.Format("15:04:05")
}
func formatText(style, text string) string {
	return fmt.Sprintf("%s%s%s", style, text, reset)
}

func formatTextExt(style, style2, text string) string {
	return fmt.Sprintf("%s%s%s%s", style, style2, text, reset)
}
