package main

import "math/big"

const PADDING_OFFSET = 0x20

func leftmostBit(length int) (i int) {
	for ; length > 0; length >>= 1 {
		i++
	}
	i -= 1
	return
}

func NCharsForNBits(length int) (nchars, nbits_c int) {
	nbits_c = leftmostBit(length)
	nchars = 8

	gcd_b := &big.Int{}
	gcd_b.GCD(nil, nil, big.NewInt(int64(nchars)), big.NewInt(int64(nbits_c)))
	gcd := int(gcd_b.Int64())

	nchars /= gcd

	return
}
