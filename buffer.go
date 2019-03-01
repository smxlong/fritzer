package fritzer

type Buffer struct {
	buffer []Sample
	head int
	size int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{buffer: make([]Sample, 2 * size), head: size - 1, size: size}
}

func (b *Buffer) Write(value Sample) {
	b.buffer[b.head] = value
	b.buffer[b.head + b.size] = value
	b.head++
	if b.head >= b.size {
		b.head -= b.size
	}
}

func (b *Buffer) Get() []Sample {
	return b.buffer[b.head:b.head+b.size]
}

func (b *Buffer) Head() Sample {
	return b.buffer[b.head]
}
