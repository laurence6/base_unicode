package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	table := []rune{}
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
			table = append(table, r)
			i++
		}
	}

	encoder := NewEncoder(table)
	decoder := NewDecoder(table)

	if true {
		encoder.Encode(bufio.NewReader(os.Stdin), os.Stdout)
	} else {
		buf := &bytes.Buffer{}
		input_file, err := os.Open("input")
		if err != nil {
			log.Fatalln(err)
		}
		encoder.Encode(input_file, buf)
		fmt.Fprintln(os.Stderr)
		decoder.Decode(buf, os.Stdout)
	}
}
