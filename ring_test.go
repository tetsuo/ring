package ring_test

import (
	"testing"

	"github.com/onur1/ring"
	"github.com/stretchr/testify/assert"
)

func TestRing(t *testing.T) {
	testCases := []struct {
		desc string
		f    func(t *testing.T)
	}{
		{
			desc: "put and get",
			f: func(t *testing.T) {
				list := ring.NewRing(2)

				list.Put(0, 42)
				list.Put(1, 555)

				assert.EqualValues(t, 42, list.Get(0))
				assert.EqualValues(t, 555, list.Get(1))
			},
		},
		{
			desc: "overflow",
			f: func(t *testing.T) {
				list := ring.NewRing(2)

				list.Put(0, 42)
				list.Put(1, 555)
				list.Put(2, 30)

				assert.EqualValues(t, 30, list.Get(0))
				assert.EqualValues(t, 555, list.Get(1))
				assert.EqualValues(t, 30, list.Get(2))
			},
		},
		{
			desc: "del",
			f: func(t *testing.T) {
				list := ring.NewRing(2)

				list.Put(0, 42)
				assert.EqualValues(t, 42, list.Get(0))

				list.Del(0)
				assert.EqualValues(t, nil, list.Get(0))
			},
		},
		{
			desc: "multiple of two",
			f: func(t *testing.T) {
				list := ring.NewRing(3)

				assert.EqualValues(t, 4, list.Size())
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, tC.f)
	}
}
