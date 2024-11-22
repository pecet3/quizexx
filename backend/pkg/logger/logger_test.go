package logger_test

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/pecet3/quizex/pkg/logger"
)

func TestLogger_Error(t *testing.T) {
	var logBuffer bytes.Buffer
	log.SetOutput(&logBuffer)

	testError := errors.New("test error")

	logger.Error(testError)

	logOutput := logBuffer.String()
	fmt.Println(logOutput)

}

func TestLogger_Info(t *testing.T) {
	var logBuffer bytes.Buffer
	log.SetOutput(&logBuffer)

	logger.Info("test")

	logOutput := logBuffer.String()
	fmt.Println(logOutput)

}

func TestLogger_Warning(t *testing.T) {
	var logBuffer bytes.Buffer
	log.SetOutput(&logBuffer)

	logger.Warn("test", "warning")

	logOutput := logBuffer.String()
	fmt.Println(logOutput)

}
