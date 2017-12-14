package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	var table []rune
	{
		table_file, err := os.Open("data")
		if err != nil {
			log.Fatalln(err)
		}
		file_reader := bufio.NewReader(table_file)

		table_bytes, err := file_reader.ReadBytes('\n')
		if err != nil {
			log.Fatalln(err)
		}
		table_bytes = table_bytes[:len(table_bytes)-1]

		for len(table_bytes) > 0 {
			r, size := utf8.DecodeRune(table_bytes)
			table = append(table, r)
			table_bytes = table_bytes[size:]
		}
	}

	encode(table, bufio.NewReader(os.Stdin), os.Stdout)
}

type BitReader struct {
	in     io.ByteReader
	n_bits uint // <= 64

	buf    uint16
	offset uint
	eof    bool
}

func (self *BitReader) fillBuffer() {
	for !self.eof && self.offset <= 8 {
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
func (self *BitReader) popNBits(n uint) uint64 {
	ret := uint64(self.buf) >> (16 - 1 - n)
	self.offset -= n
	return ret
}

func (self *BitReader) Read() (uint64, bool) {
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

func encode(table []rune, in io.ByteReader, out io.Writer) {
	n_bits := max_bits_represented(len(table))
	reader := BitReader{in: in, n_bits: n_bits}

	for {
		index, eof := reader.Read()
		ch := table[index]
		fmt.Fprintf(os.Stderr, "%d: %s\n", index, string(ch))
		bytes := make([]byte, utf8.RuneLen(ch))
		utf8.EncodeRune(bytes, ch)
		out.Write(bytes)
		if eof {
			break
		}
	}
}

func max_bits_represented(length int) (i uint) {
	for ; length > 0; length >>= 1 {
		i++
	}
	i -= 1
	return
}

func min(a, b uint) uint {
	if a < b {
		return a
	} else {
		return b
	}
}
