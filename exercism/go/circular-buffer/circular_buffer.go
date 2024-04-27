package circular

import "errors"

// Buffer type.
type Buffer struct {
	Data        []byte
	Cap, Len    int // Capacity and current size
	Read, Write int // read and write index
}

func NewBuffer(size int) *Buffer {
	if size < 0 {
		return nil
	}
	return &Buffer{Data: make([]byte, size), Cap: size}
}

func (b *Buffer) ReadByte() (n byte, err error) {
	if b == nil || b.Len == 0 {
		err = errors.New("nothing to read")
		return
	}
	n = b.Data[b.Read]
	b.Read = (b.Read + 1) % len(b.Data)
	b.Len--
	return
}

func (b *Buffer) WriteByte(c byte) error {
	if b == nil || b.Len == b.Cap {
		return errors.New("buffer full")
	}
	b.Data[b.Write] = c
	b.Write = (b.Write + 1) % len(b.Data)
	b.Len++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b == nil || b.Cap == 0 {
		return
	}
	b.Data[b.Write] = c
	b.Write = (b.Write + 1) % len(b.Data)
	// move read index ahead if we are writing on a full buffer
	if b.Len == b.Cap {
		b.Read = (b.Read + 1) % len(b.Data)
	} else {
		b.Len++
	}
}

func (b *Buffer) Reset() {
	if b == nil {
		return
	}
	b.Len, b.Read, b.Write = 0, 0, 0
}
