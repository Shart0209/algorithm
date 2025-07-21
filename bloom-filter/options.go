package bloom_filter

type Option func(bf *BloomFilter)

func EnableOptimal(fpRate float64) Option {
	return func(bf *BloomFilter) {
		bf.fpRate = fpRate
		bf.enableOptimal = true
	}
}
