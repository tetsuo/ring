# ring

A Ring represents a data structure (known as [circular buffer](https://en.wikipedia.org/wiki/Circular_buffer)) that uses a single fixed-size buffer as if it were connected end-to-end.

In this representation, the capacity of a buffer must be a power of 2.

## Example

```go
r := ring.NewRing[int](2)

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
