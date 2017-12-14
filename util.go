package main

func maxBit(length int) (i uint) {
	for ; length > 0; length >>= 1 {
		i++
	}
	i -= 1
	return
}

func NCharsForNBytes(length int) (nc uint, nb uint) {
	nc = 8
	nb = maxBit(length)
	return
}

func Min(a, b uint) uint {
	if a < b {
		return a
	} else {
		return b
	}
}
