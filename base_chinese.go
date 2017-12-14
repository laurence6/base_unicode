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

		var r rune
		for r != '\n' {
			r, _, err = file_reader.ReadRune()
			if err != nil {
				log.Fatalln(err)
			}
			table = append(table, r)
		}
	}

	Encode(table, bufio.NewReader(os.Stdin), os.Stdout)
}

func Encode(table []rune, in io.ByteReader, out io.Writer) {
	n_bits := max_bits_represented(len(table))
	reader := NewBitsReader(in, n_bits)

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
