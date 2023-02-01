package ring

type Ring[T any] struct {
	size int
	mask int
	buf  []T
}

func NewRing[T any](size int) *Ring[T] {
	r := new(Ring[T])
	if size > 0 && (size&(size-1)) < 0 {
		r.reset(size)
	} else {
		n := 1
		for n < size {
			n <<= 1
		}
		r.reset(n)
	}
	return r
}

func (c *Ring[T]) reset(size int) {
	*c = Ring[T]{
		mask: size - 1,
		size: size,
		buf:  make([]T, size),
	}
}

func (c *Ring[T]) Size() int {
	return c.size
}

func (c *Ring[T]) Get(index int) T {
	return c.buf[index&c.mask]
}

func (c *Ring[T]) Put(index int, val T) int {
	i := index & c.mask
	c.buf[i] = val
	return i
}

func (c *Ring[T]) Del(index int) T {
	i := index & c.mask
	val := c.buf[i]
	x := *new(T)
	c.buf[i] = x
	return val
}
