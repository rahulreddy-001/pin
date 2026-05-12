package pin

import (
	"encoding/json"
	"fmt"
	"time"
)

type Encoder interface {
	Encode(Log) string
}

type Log struct {
	Name      string    `json:"name"`
	Level     string    `json:"level"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Fields    Fields    `json:"fields"`
	Source    string    `json:"source,omitempty"`
}

func (l *Log) Encode() string {
	return DefaultEncoder{}.Encode(*l)
}

type DefaultEncoder struct{}

func (DefaultEncoder) Encode(l Log) string {
	fieldsMap := map[string]any{
		"name":      l.Name,
		"level":     l.Level,
		"timestamp": l.Timestamp,
		"message":   l.Message,
	}

	if l.Source != "" {
		fieldsMap["source"] = l.Source
	}

	for _, field := range l.Fields {
		fieldsMap[field.Key] = field.Value
	}

	data, _ := json.Marshal(fieldsMap)
	return string(data)

}


type DefaultCLIEncoder struct{}

func (_ DefaultCLIEncoder) Encode(log Log) string {
	return fmt.Sprintf("[%s]  [%s] %s %s  %s  %s", log.Level, log.Name, log.Source, log.Timestamp.Local().Format(time.RFC1123), log.Message, log.Fields.Encode())
}
