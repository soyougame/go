package jsoniter

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestBytesPool_PutString(t *testing.T) {
	var (
		capacity = 10
		pool     = newBytesPool(capacity, true)
		times    = 10
		address  []byte
		current  []byte
	)

	for i := 0; i < times; i++ {
		x := pool.Get()
		x = x[:0]

		current = x

		if address == nil {
			address = current
		} else {
			require.EqualValuesf(t, fmt.Sprintf(`%p`, address), fmt.Sprintf(`%p`, current), `第%d次`, i+1)
		}

		x = append(x, []byte(`a`)...)

		pool.PutString(*(*string)(unsafe.Pointer(&x)))

		require.EqualValues(t, 1, pool.counter.Load())
	}
}
