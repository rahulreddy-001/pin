package pin

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

type logger struct {
	name   string
	level  LogLevel
	buffer *buffer
	writer *writer
}

func NewLogger(config *Config) *logger {
	buffer := newBuffer()
	return &logger{
		name:   config.Name,
		level:  config.LogLevel,
		buffer: buffer,
		writer: newWriter(config.Writer, buffer, config.Encoder, config.FlushInterval),
	}
}

func (l *logger) Clone(name string) *logger {
	return &logger{
		name:   name,
		level:  l.level,
		buffer: l.buffer,
		writer: l.writer,
	}
}

func (l *logger) Flush() {
	l.writer.closeChan <- struct{}{}
	<-l.writer.closeChan
	close(l.writer.closeChan)
}

func (l *logger) log(level LogLevel, msg string, fields ...Field) {
	if !level.IsGreaterThan(l.level) {
		return
	}

	l.buffer.add(Log{
		Name:      l.name,
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
