package main

import (
	"io"
	"log"
	"math/big"
	"unicode"
)

type Decoder struct {
	table map[rune]uint64

	nchars  int // number of characters
	nbits_c int // number of bits represented by a character
	nbytes  int // number of bytes represented by nchars characters
}

func NewDecoder(table []rune) Decoder {
	nchars, nbits_c := NCharsForNBits(len(table))
	nbytes := nchars * nbits_c / 8

	table_r := map[rune]uint64{}
	for i, r := range table {
		table_r[r] = uint64(i)
	}

	return Decoder{
		table:   table_r,
		nchars:  nchars,
		nbits_c: nbits_c,
		nbytes:  nbytes,
	}
}

func (self *Decoder) assembleBytes(indices []uint64) []byte {
	bitset := big.Int{}

	for i := 0; i < len(indices); i++ {
		bitset.Lsh(&bitset, uint(self.nbits_c))
		bitset.Add(&bitset, big.NewInt(int64(indices[i])))
	}

	bytes := bitset.Bytes()
	if n_leading_zeros := self.nbytes - len(bytes); n_leading_zeros > 0 {
		leading_zeros := make([]byte, n_leading_zeros)
		bytes = append(leading_zeros, bytes...)
	}

	return bytes
}

func (self *Decoder) Decode(in io.RuneReader, out io.Writer) {
	in_runes := make([]rune, self.nchars)
	in_indices := make([]uint64, self.nchars)
	var bytes []byte
	n_paddings := 0
	for {
		for i := 0; i < self.nchars; i++ {
			in_runes[i] = 0
			in_indices[i] = 0
		}

		var err error
		n := 0
		for n < self.nchars && err == nil {
			r := rune(0)
			r, _, err = in.ReadRune()
			if r == PADDING_CHAR {
				for {
					r, _, err = in.ReadRune()
					if err != nil && err != io.EOF {
						log.Fatalln(err)
					}
					if '0' <= r && r <= '9' {
						n_paddings *= 10
						n_paddings += int(r - '0')
					} else {
						break
					}
				}
				break
			} else if unicode.IsPrint(r) && !unicode.IsSpace(r) {
				in_runes[n] = r
				n++
			}
		}
		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}

		for i, r := range in_runes {
			if index, ok := self.table[r]; ok {
				in_indices[i] = index
			}
		}

		if bytes == nil {
			if n == 0 {
				break
			}
		} else {
			if n_paddings > 0 {
				if n == 0 {
					out.Write(bytes[:len(bytes)-n_paddings])
					break
				} else {
					out.Write(bytes)
				}
			} else {
				if n == self.nchars {
					out.Write(bytes)
				} else if n == 0 {
					out.Write(bytes)
					break
				} else {
					break
				}
			}
		}

		bytes = self.assembleBytes(in_indices)
	}
}
