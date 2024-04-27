package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	Bytes  int64
	Ops    int
	Reader io.Reader
	Lock   sync.Mutex
}

type writeCounter struct {
	Bytes  int64
	Ops    int
	Writer io.Writer
	Lock   sync.Mutex
}

type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{Writer: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{Reader: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{NewReadCounter(readwriter), NewWriteCounter(readwriter)}
}

func (c *readCounter) Read(p []byte) (int, error) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	n, err := c.Reader.Read(p)
	c.Bytes += int64(n)
	c.Ops++
	return n, err
}

func (c *readCounter) ReadCount() (int64, int) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	return c.Bytes, c.Ops
}

func (c *writeCounter) Write(p []byte) (int, error) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	n, err := c.Writer.Write(p)
	c.Bytes += int64(n)
	c.Ops++
	return n, err
}

func (c *writeCounter) WriteCount() (int64, int) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	return c.Bytes, c.Ops
}
