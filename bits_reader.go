package main

import (
	"io"
	"log"
)

type BitsReader struct {
	in     io.ByteReader
	n_bits uint // <= 64

	buf    uint16
	offset uint
	eof    bool
}

func NewBitsReader(in io.ByteReader, n_bits uint) BitsReader {
	return BitsReader{
		in:     in,
		n_bits: n_bits,
	}
}

func (self *BitsReader) fillBuffer() {
	for !self.eof && self.offset <= (16-8) {
		b, err := self.in.ReadByte()
		if err != nil {
			if err == io.EOF {
				self.eof = true
			} else {
				log.Fatalln(err)
			}
		}

		self.offset += 8
		self.buf &= 0xff << (16 - self.offset)
		self.buf += uint16(b) << (8 - self.offset)
	}
}

// n <= self.offset
func (self *BitsReader) popNBits(n uint) uint64 {
	ret := uint64(self.buf) >> (16 - 1 - n)
	self.offset -= n
	return ret
}

func (self *BitsReader) Read() (uint64, bool) {
	var ret uint64 = 0
	n_bits := self.n_bits

	for n_bits > 0 {
		n := min(n_bits, self.offset)
		ret <<= n
		ret += self.popNBits(n)
		n_bits -= n

		if !self.eof {
			self.fillBuffer()
		} else {
			break
		}
	}

	if n_bits > 0 && (self.offset > 0 || !self.eof) {
		log.Fatalln("Error!")
	}

	if n_bits > 0 {
		ret <<= n_bits
		return ret, true
	} else {
		return ret, false
	}
}
