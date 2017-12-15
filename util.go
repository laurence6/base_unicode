package main

import (
	"fmt"
	"math/big"
	"os"
)

func maxBit(length int) (i int) {
	for ; length > 0; length >>= 1 {
		i++
	}
	i -= 1
	return
}

func NCharsForNBytes(length int) (nc, nb int) {
	nc = 8
	nb = maxBit(length)
	return
}

func SplitBytes(bytes []byte, nc int) []uint64 {
	ret := make([]uint64, nc)

	nbi := len(bytes) * 8 / nc

	bitset := big.Int{}
	bitset.SetBytes(bytes)

	printBitset(bitset, len(bytes))

	for i := 0; i < nc; i++ {
		v := uint64(0)
		upper := (len(bytes)*8 - i*nbi)
		lower := Max(0, upper-nbi)
		for j := upper - 1; j >= lower; j-- {
			v <<= 1
			v += uint64(bitset.Bit(j))
		}
		ret[i] = v
	}

	return ret
}

func AssemleBytes(indices []uint64, nb int) []byte {
	nbi := nb * 8 / len(indices)

	bitset := big.Int{}

	for i := 0; i < len(indices); i++ {
		bitset.Lsh(&bitset, uint(nbi))
		bitset.Add(&bitset, big.NewInt(int64(indices[i])))
	}

	printBitset(bitset, nb)

	return bitset.Bytes()
}

func printBitset(bitset big.Int, nb int) {
	return
	for i := nb*8 - 1; i >= 0; i-- {
		fmt.Fprint(os.Stderr, bitset.Bit(i))
		if i%8 == 0 {
			fmt.Fprint(os.Stderr, " ")
		}
	}
	fmt.Fprintln(os.Stderr)
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
