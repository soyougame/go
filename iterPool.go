package jsoniter

import (
	"sync"
	"unsafe"

	"go.uber.org/atomic"
)

type bytesPool struct {
	*sync.Pool
	counter *atomic.Int64
}

func newBytesPool(length int, count bool) *bytesPool {
	result := &bytesPool{
		Pool: &sync.Pool{
			New: func() any {
				return make([]byte, length)
			},
		},
		counter: &atomic.Int64{},
	}

	if count {
		result.counter = &atomic.Int64{}
		result.Pool.New = func() any {
			result.counter.Inc()

			return make([]byte, length)
		}
	}

	return result
}

func (b *bytesPool) Get() []byte {
	return b.Pool.Get().([]byte)
}

func (b *bytesPool) PutBytes(data []byte) {
	if data == nil || cap(data) == 0 {
		return
	}

	b.Pool.Put(data)
}

func (b *bytesPool) PutString(data string) {
	b.PutBytes(*(*[]byte)(unsafe.Pointer(&data)))
}
