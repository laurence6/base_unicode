package main

const PADDING_OFFSET = 0x20

func leftmostBit(length int) (i int) {
	for ; length > 0; length >>= 1 {
		i++
	}
	i -= 1
	return
}

func gcd(a, b int) int {
	if b > 0 {
		for {
			if a %= b; a == 0 {
				break
			}
			if b %= a; b == 0 {
				break
			}
		}
	}
	return a + b
}

func NCharsForNBits(length int) (nchars, nbits_c int) {
	nbits_c = leftmostBit(length)
	nchars = 8 / gcd(nbits_c, 8)
	return
}
