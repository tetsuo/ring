package ring

type Ring struct {
	size int
	mask int
	buf  []interface{}
}

func NewRing(size int) *Ring {
	r := new(Ring)
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

func (c *Ring) reset(size int) {
	*c = Ring{
		mask: size - 1,
		size: size,
		buf:  make([]interface{}, size),
	}
}

func (c *Ring) Size() int {
	return c.size
}

func (c *Ring) Get(index int) interface{} {
	return c.buf[index&c.mask]
}

func (c *Ring) Put(index int, val interface{}) int {
	i := index & c.mask
	c.buf[i] = val
	return i
}

func (c *Ring) Del(index int) interface{} {
	i := index & c.mask
	val := c.buf[i]
	c.buf[i] = nil
	return val
}
