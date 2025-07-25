package base62

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestUnit_Encode(t *testing.T) {
	b := New("")

	enc := b.Encode(math.MaxUint64)
	require.Equal(t, uint64(math.MaxUint64), b.Decode(enc))
}

// cpu: Intel(R) Core(TM) Ultra 7 165H
// BenchmarkBase62_Encode
// BenchmarkBase62_Encode-22    	178984873	         6.983 ns/op	       0 B/op	       0 allocs/op
func BenchmarkBase62_Encode(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		bs := New("")
		for pb.Next() {
			enc := bs.Encode(math.MaxUint64)
			bs.Decode(enc)
		}
	})
}
