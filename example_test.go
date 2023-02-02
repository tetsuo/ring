package ring_test

import (
	"fmt"

	"github.com/onur1/ring"
)

func ExampleRing() {
	r := ring.NewRing[int](2)

	r.Put(0, 42)
	r.Put(1, 555)

	fmt.Println(r.Get(0))
	fmt.Println(r.Get(1))
	fmt.Println(r.Get(2))
	fmt.Println(r.Get(3))
	fmt.Println(r.Get(4))

	// Output:
	// 42
	// 555
	// 42
	// 555
	// 42
}
