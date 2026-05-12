package pin

import (
	"fmt"
	"io"
	"path/filepath"
	"runtime"
	"time"
)

type logger struct {
	Name   string
	Level  LogLevel
	Buffer *Buffer
	Writer *Writer
}

func NewLogger(name string, level LogLevel, writer io.Writer, encoder Encoder) *logger {
	buffer := NewBuffer()
	return &logger{
		Name:   name,
		Level:  level,
		Buffer: buffer,
		Writer: NewWriter(writer, buffer, encoder),
	}
}

func (l *logger) Clone(name string) *logger {
	return &logger{
		Name:   name,
		Level:  l.Level,
		Buffer: l.Buffer,
		Writer: l.Writer,
	}
}

func (l *logger) Flush() {
	l.Writer.closeChan <- struct{}{}
	close(l.Writer.closeChan)
	<-l.Writer.exitChan
}

func (l *logger) log(level LogLevel, msg string, fields ...Field) {
	if !level.IsGreaterThan(l.Level) {
		return
	}

	l.Buffer.Add(Log{
		Name:      l.Name,
		Level:     level.GetLogLevel(),
		Timestamp: time.Now(),
		Message:   msg,
		Fields:    fields,
		Source:    getCaller(3),
	})
}

func (l *logger) Debug(msg string, fields ...Field) {
	l.log(DEBUG, msg, fields...)
}

func (l *logger) Info(msg string, fields ...Field) {
	l.log(INFO, msg, fields...)
}

func (l *logger) Warn(msg string, fields ...Field) {
	l.log(WARN, msg, fields...)
}

func (l *logger) Error(msg string, fields ...Field) {
	l.log(ERROR, msg, fields...)
}

func (l *logger) Fatal(msg string, fields ...Field) {
	l.log(FATAL, msg, fields...)
}

func getCaller(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}

	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}
