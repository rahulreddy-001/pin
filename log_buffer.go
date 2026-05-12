package pin

import "sync"

type Buffer struct {
	mu   sync.Mutex
	logs []Log
}

func NewBuffer() *Buffer {
	return &Buffer{
		mu:   sync.Mutex{},
		logs: []Log{},
	}
}

func (b *Buffer) Add(log Log) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.logs = append(b.logs, log)
}

func (b *Buffer) AddLogs(logs ...Log) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.logs = append(b.logs, logs...)
}

func (b *Buffer) GetLogs(count int) []Log {
	b.mu.Lock()
	defer b.mu.Unlock()
	count = min(count, len(b.logs))
	l := b.logs[:count]
	b.logs = b.logs[count:]
	return l
}
