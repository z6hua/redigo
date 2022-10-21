package server

type ReplBuffer struct {
	buf  []byte
	len  int
	free int
}

func NewReplBuffer(size int) *ReplBuffer {
	return &ReplBuffer{
		buf:  make([]byte, size),
		len:  0,
		free: size,
	}
}

func (b *ReplBuffer) Size() int {
	return len(b.buf)
}

func (b *ReplBuffer) Len() int {
	return b.len
}

func (b *ReplBuffer) Get(offset int) []byte {
	return (*b).buf[offset:(*b).len]
}

func (b *ReplBuffer) Push(s string) {
	sLen := len(s)
	if (*b).free < sLen {
		b.pop(sLen)
	}
	copy((*b).buf[(*b).len:(*b).len+sLen], s)
	(*b).len += sLen
	(*b).free -= sLen
}

func (b *ReplBuffer) pop(offset int) {
	if offset <= len((*b).buf) {
		tmp := (*b).buf[offset:]
		(*b).buf = make([]byte, len((*b).buf))
		(*b).buf = tmp
		(*b).len -= offset
		(*b).free += offset
	}
}
