package main

const PADDING_OFFSET = 0x20

func leftmostBit(length int) (i int) {
	for ; length > 0; length >>= 1 {
		i++
	}
	i -= 1
	return
}

// FIXME: compute LCM
func NCharsForNBits(length int) (nchars, nbits_c int) {
	nchars = 8
	nbits_c = leftmostBit(length)
	return
}
