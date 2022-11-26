package logger_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sr-codefreak/portfolio-visualiser/foundation/logger"
	"strings"
	"testing"
)

// TestNewLogger helps test the logger.
func TestNewLogger(t *testing.T) {

	log := logger.NewLogger()
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.WithField("test", "test").Infof("Testing ongoing")
	s := buf.String()

	//Resetting the logger writer
	logFile, err := os.Create("log.json")
	multi := io.MultiWriter(os.Stdout)
	if err == nil {
		multi = io.MultiWriter(logFile, os.Stdout)
	}
	log.SetOutput(multi)

	want := fmt.Sprintln(`time="2022-11-27T00:40:43+05:30" level=info msg="Testing ongoing" test=test`)
	if strings.Contains(s, want) {
		t.Fatalf("Creating New logger test failed, got %s, want %s", s, want)
	}
}
