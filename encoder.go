package main

import (
	"bytes"
	"io"
	"log"
	"math/big"
)

type Encoder struct {
	table []rune

	nchars  int // number of characters
	nbits_c int // number of bits represented by a character
	nbytes  int // number of bytes represented by nchars characters
}

func NewEncoder(table []rune) Encoder {
	nchars, nbits_c := NCharsForNBits(len(table))
	nbytes := nchars * nbits_c / 8

	return Encoder{
		table:   table,
		nchars:  nchars,
		nbits_c: nbits_c,
		nbytes:  nbytes,
	}
}

func (self *Encoder) splitBytes(bytes []byte) []uint64 {
	ret := make([]uint64, self.nchars)

	bitset := big.Int{}
	bitset.SetBytes(bytes)

	for i := 0; i < self.nchars; i++ {
		v := uint64(0)
		upper := (len(bytes)*8 - i*self.nbits_c)
		lower := upper - self.nbits_c
		for j := upper - 1; j >= lower; j-- {
			v <<= 1
			v += uint64(bitset.Bit(j))
		}
		ret[i] = v
	}

	return ret
}

func (self *Encoder) Encode(in io.Reader, out io.Writer) {
	in_bytes := make([]byte, self.nbytes)
	buf := &bytes.Buffer{}
	for {
		for i := 0; i < self.nbytes; i++ {
			in_bytes[i] = 0
		}
		buf.Reset()

		n, err := io.ReadFull(in, in_bytes)

		if err != nil && err != io.ErrUnexpectedEOF {
			if err == io.EOF {
				break
			} else {
				log.Fatalln(err)
			}
		}

		indices := self.splitBytes(in_bytes)

		_nc := self.nchars
		if n < self.nbytes {
			_nc = (n * 8) / self.nbits_c
			if (n*8)%self.nbits_c > 0 {
				_nc += 1
			}
		}
		for i := 0; i < _nc; i++ {
			buf.WriteRune(self.table[indices[i]])
		}

		if self.nbytes-n > 0 {
			buf.WriteRune(rune(PADDING_OFFSET + self.nbytes - n))
		}

		out.Write(buf.Bytes())
	}
}
