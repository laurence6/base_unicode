package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

const PADDING_OFFSET = 0x20

func main() {
	table1 := []rune{}
	table2 := map[rune]uint64{}
	{
		table_file, err := os.Open("data")
		if err != nil {
			log.Fatalln(err)
		}
		file_reader := bufio.NewReader(table_file)

		r := rune(0)
		i := uint64(0)
		for {
			r, _, err = file_reader.ReadRune()
			if err != nil {
				log.Fatalln(err)
			}
			if r == '\n' {
				break
			}
			table1 = append(table1, r)
			table2[r] = i
			i++
		}
	}

	if true {
		Encode(table1, bufio.NewReader(os.Stdin), os.Stdout)
	} else {
		buf := &bytes.Buffer{}
		input_file, err := os.Open("input")
		if err != nil {
			log.Fatalln(err)
		}
		Encode(table1, input_file, buf)
		fmt.Fprintln(os.Stderr)
		Decode(table2, buf, os.Stdout)
	}
}

func Encode(table []rune, in io.Reader, out io.Writer) {
	nc, nb := NCharsForNBytes(len(table))
	nbi := nb * 8 / nc

	in_bytes := make([]byte, nb)
	buf := &bytes.Buffer{}
	for {
		for i := 0; i < nb; i++ {
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

		indices := SplitBytes(in_bytes, nc)

		_nc := nc
		if n < nb {
			_nc = (n * 8) / nbi
			if (n*8)%nbi > 0 {
				_nc += 1
			}
		}
		for i := 0; i < _nc; i++ {
			buf.WriteRune(table[indices[i]])
		}

		if nb-n > 0 {
			buf.WriteRune(rune(PADDING_OFFSET + nb - n))
		}

		out.Write(buf.Bytes())
	}
}

func Decode(table map[rune]uint64, in io.RuneReader, out io.Writer) {
	nc, nb := NCharsForNBytes(len(table))

	in_runes := make([]rune, nc)
	in_indices := make([]uint64, nc)
	for {
		for i := 0; i < nc; i++ {
			in_runes[i] = 0
			in_indices[i] = 0
		}

		var err error
		n := 0
		for n < nc && err == nil {
			r := rune(0)
			r, _, err = in.ReadRune()
			in_runes[n] = r
			n++
		}

		n_paddings := 0
		for i, r := range in_runes {
			if index, ok := table[r]; ok {
				in_indices[i] = index
			} else if n_paddings == 0 {
				n_paddings = int(r - PADDING_OFFSET)
			}
		}

		bytes := AssemleBytes(in_indices, nb)
		out.Write(bytes[:len(bytes)-n_paddings])

		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalln(err)
			}
		}
	}
}
