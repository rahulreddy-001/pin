package pin

import (
	"fmt"
	"io"
	"time"
)

type Writer struct {
	FlushInterval time.Duration
	Buffer        *Buffer
	Encoder       Encoder
	Writer        io.Writer
	closeChan     chan struct{}
	exitChan      chan struct{}
}

func NewWriter(writer io.Writer, buffer *Buffer, encoder Encoder) *Writer {
	w := &Writer{
		FlushInterval: time.Microsecond,
		Buffer:        buffer,
		Writer:        writer,
		Encoder:       encoder,
		closeChan:     make(chan struct{}),
		exitChan:      make(chan struct{}),
	}
	go w.Flush()
	return w
}

func (b *Writer) flush() {
	for {
		logs := b.Buffer.GetLogs(100)
		if len(logs) == 0 {
			break
		}
		for _, log := range logs {
			fmt.Fprintln(b.Writer, b.Encoder.Encode(log))
		}
	}
}

func (b *Writer) Flush() {
	ticker := time.NewTicker(b.FlushInterval)
	defer func() {
		b.exitChan <- struct{}{}
		close(b.exitChan)
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
