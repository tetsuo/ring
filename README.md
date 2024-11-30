# ring

A `Ring` is a fixed size [circular buffer](https://en.wikipedia.org/wiki/Circular_buffer).

The capacity of a buffer is always rounded up to the nearest power of 2 for optimal performance.

## Example

```go
r := ring.New[int](2)

r.Put(0, 42)
r.Put(1, 555)

fmt.Println(r.Get(0))
fmt.Println(r.Get(1))
fmt.Println(r.Get(2))
fmt.Println(r.Get(3))
fmt.Println(r.Get(4))
```

Output:

```
42
555
42
555
42
```
