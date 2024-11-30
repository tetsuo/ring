package ring

// A Ring is a fixed-size circular buffer.
type Ring[T any] struct {
	size int // Capacity of the buffer
	mask int // Bitmask for efficient modulo operation
	buf  []T // Underlying buffer
}

// New creates a new Ring with the specified size.
// The size is rounded up to the nearest power of 2 if it isn't already.
func New[T any](size int) *Ring[T] {
	if size <= 0 {
		panic("size must be greater than 0")
	}
	size = nextPowerOfTwo(size)
	return &Ring[T]{
		size: size,
		mask: size - 1,
		buf:  make([]T, size),
	}
}

// Size returns the capacity of the internal buffer.
func (r *Ring[T]) Size() int {
	return r.size
}

// Get returns the value at given index.
func (r *Ring[T]) Get(index int) T {
	return r.buf[index&r.mask]
}

// Put inserts a value at given index.
func (r *Ring[T]) Put(index int, val T) int {
	i := index & r.mask
	r.buf[i] = val
	return i
}

// Del deletes a value at given index and returns it.
func (r *Ring[T]) Del(index int) T {
	i := index & r.mask
	val := r.buf[i]
	var zero T
	r.buf[i] = zero
	return val
}

// Helper function to calculate the next power of 2 for a given integer.
// If the input is already a power of 2, it returns the same value.
func nextPowerOfTwo(n int) int {
	if n&(n-1) == 0 {
		return n
	}
	power := 1
	for power < n {
		power <<= 1
	}
	return power
}
