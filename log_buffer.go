package pin

import "sync"

type buffer struct {
	mu   sync.Mutex
	logs []Log
}

func newBuffer() *buffer {
	return &buffer{
		mu:   sync.Mutex{},
		logs: []Log{},
	}
}

func (b *buffer) add(log Log) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.logs = append(b.logs, log)
}

func (b *buffer) addLogs(logs ...Log) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.logs = append(b.logs, logs...)
}

func (b *buffer) getLogs(count int) []Log {
	b.mu.Lock()
	defer b.mu.Unlock()
	count = min(count, len(b.logs))
	l := b.logs[:count]
	b.logs = b.logs[count:]
	return l
}
