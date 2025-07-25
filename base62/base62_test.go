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

//cpu: Intel(R) Core(TM) Ultra 7 165H
//BenchmarkBloomFilter_Check
//BenchmarkBloomFilter_Check-22    	54544496	        18.53 ns/op	      16 B/op	       1 allocs/op
func BenchmarkBloomFilter_Check(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		bs := New("")
		for pb.Next() {
			enc := bs.Encode(math.MaxUint64)
			bs.Decode(enc)
		}
	})
}
