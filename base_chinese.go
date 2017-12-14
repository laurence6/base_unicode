package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
)

func main() {
	table1 := []rune{}
	table2 := map[rune]uint64{}
	{
		table_file, err := os.Open("data")
		if err != nil {
			log.Fatalln(err)
		}
		file_reader := bufio.NewReader(table_file)

		var r rune
		var i uint64 = 0
		for r != '\n' {
			r, _, err = file_reader.ReadRune()
			if err != nil {
				log.Fatalln(err)
			}
			table1 = append(table1, r)
			table2[r] = i
			i++
		}
	}

	//Encode(table1, bufio.NewReader(os.Stdin), os.Stdout)

	buf := &bytes.Buffer{}
	input_file, err := os.Open("input")
	if err != nil {
		log.Fatalln(err)
	}
	Encode(table1, bufio.NewReader(input_file), buf)
	Decode(table2, buf, bufio.NewWriter(os.Stdout))
}

func Encode(table []rune, in io.Reader, out io.Writer) {
	nc, nb := NCharsForNBytes(len(table))

	bytes := make([]byte, nb)
	for {
		n, err := in.Read(bytes)
		if uint(n) == nb {
			for i := uint(0); i < nc; i++ {

			}
		} else {

		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalln(err)
			}
		}
	}
}

func Decode(table map[rune]uint64, in io.RuneReader, out io.Writer) {
	nc, nb := NCharsForNBytes(len(table))

	for {
		r, _, err := in.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalln(err)
			}
		}
		b, ok := table[r]
		if !ok {
			log.Fatalf("unexpected rune: %s\n", string(r))
		}
	}
}
