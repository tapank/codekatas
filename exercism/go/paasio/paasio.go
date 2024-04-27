package paasio

import (
	"io"
	"sync"
)

type Counter struct {
	RBytes int64
	ROps   int
	Reader io.Reader
	RLock  sync.Mutex
	WBytes int64
	WOps   int
	Writer io.Writer
	WLock  sync.Mutex
}

type readCounter Counter
type writeCounter Counter

func (c *Counter) Read(p []byte) (n int, err error) {
	c.RLock.Lock()
	defer c.RLock.Unlock()

	n, err = c.Reader.Read(p)
	if err == nil {
		c.RBytes += int64(n)
		c.ROps++
	}
	return
}

func (c *Counter) Write(p []byte) (n int, err error) {
	c.WLock.Lock()
	defer c.WLock.Unlock()

	n, err = c.Writer.Write(p)
	if err == nil {
		c.WBytes += int64(n)
		c.WOps++
	}
	return
}

func (c *Counter) ReadCount() (int64, int) {
	c.RLock.Lock()
	defer c.RLock.Unlock()

	return c.RBytes, c.ROps
}

func (c *Counter) WriteCount() (int64, int) {
	c.WLock.Lock()
	defer c.WLock.Unlock()

	return c.WBytes, c.WOps
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{Writer: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{Reader: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &Counter{Reader: readwriter, Writer: readwriter}
}

func (rc *readCounter) Read(p []byte) (n int, err error) {
	rc.RLock.Lock()
	defer rc.RLock.Unlock()

	n, err = rc.Reader.Read(p)
	if err == nil {
		rc.RBytes += int64(n)
		rc.ROps++
	}
	return
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.RLock.Lock()
	defer rc.RLock.Unlock()

	return rc.RBytes, rc.ROps
}

func (wc *writeCounter) Write(p []byte) (n int, err error) {
	wc.WLock.Lock()
	defer wc.WLock.Unlock()

	n, err = wc.Writer.Write(p)
	if err == nil {
		wc.WBytes += int64(n)
		wc.WOps++
	}
	return
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.WLock.Lock()
	defer wc.WLock.Unlock()

	return wc.WBytes, wc.WOps
}
