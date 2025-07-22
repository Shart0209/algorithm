package bloom_filter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnit_Check(t *testing.T) {
	fb, err := New(50, EnableOptimal(0.01))
	require.NoError(t, err)

	fb.Add([]byte("test_1"))
	require.True(t, fb.Check([]byte("test_1")))

	fb.Add([]byte("test_12"))
	require.True(t, fb.Check([]byte("test_12")))

	fb.Add([]byte("test_30"))
	require.True(t, fb.Check([]byte("test_30")))

	require.False(t, fb.Check([]byte("test_510")))
	require.False(t, fb.Check([]byte("test_230")))
	require.False(t, fb.Check([]byte("test_131")))
}

func BenchmarkBloomFilter_Check(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		fb, err := New(10_000_000)
		require.NoError(b, err)
		fb.Add([]byte("test_5"))
		for pb.Next() {
			//fb, err := New(10_000_000)
			//require.NoError(b, err)
			//fb.Add([]byte("test_5"))
			fb.Check([]byte("test_505"))
		}
	})
}
