package pin

import (
	"fmt"
	"io"
	"time"
)

type writer struct {
	FlushInterval time.Duration
	Buffer        *buffer
	Encoder       Encoder
	Writer        io.Writer
	closeChan     chan struct{}
}

func newWriter(output io.Writer, buffer *buffer, encoder Encoder, flushInterval time.Duration) *writer {
	w := &writer{
		FlushInterval: flushInterval,
		Buffer:        buffer,
		Writer:        output,
		Encoder:       encoder,
		closeChan:     make(chan struct{}),
	}
	go w.flushLoop()
	return w
}

func (b *writer) flush() {
	for {
		logs := b.Buffer.getLogs(100)
		if len(logs) == 0 {
			break
		}
		for _, log := range logs {
			fmt.Fprintln(b.Writer, b.Encoder.Encode(log))
		}
	}
}

func (b *writer) flushLoop() {
	ticker := time.NewTicker(b.FlushInterval)
	defer func() {
		b.closeChan <- struct{}{}
		ticker.Stop()
	}()

	for {
		select {
		case <-ticker.C:
			b.flush()
		case <-b.closeChan:
			b.flush()
			return
		}
	}
}
