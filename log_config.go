package pin

import (
	"io"
	"os"
	"time"
)

type Config struct {
	Name          string
	LogLevel      LogLevel
	Writer        io.Writer
	Encoder       Encoder
	FlushInterval time.Duration
}

func NewLoggerConfig() *Config {
	return &Config{
		Name:          "main",
		LogLevel:      LEVEL_PROD,
		Writer:        os.Stdout,
		Encoder:       DefaultCLIEncoder{},
		FlushInterval: time.Second,
	}
}

func (c *Config) WithName(name string) *Config {
	c.Name = name
	return c
}

func (c *Config) WithLogLevel(level LogLevel) *Config {
	c.LogLevel = level
	return c
}

func (c *Config) WithWriter(writer io.Writer) *Config {
	c.Writer = writer
	return c

}
func (c *Config) WithEncoder(encoder Encoder) *Config {
	c.Encoder = encoder
	return c

}
func (c *Config) WithFlushInterval(flushInterval time.Duration) *Config {
	c.FlushInterval = flushInterval
	return c
}
