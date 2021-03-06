package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var DECODE = false
var COLS = 40
var TABLE_PATH = ""

func init() {
	flag.BoolVar(&DECODE, "d", DECODE, "decode data")
	flag.BoolVar(&DECODE, "decode", DECODE, "decode data")

	flag.IntVar(&COLS, "w", COLS, "wrap lines after n characters (0 to disable wrap)")
	flag.IntVar(&COLS, "wrap", COLS, "wrap lines after n characters (0 to disable wrap)")

	flag.StringVar(&TABLE_PATH, "t", TABLE_PATH, "path of table (if empty, use embedded table)")
	flag.StringVar(&TABLE_PATH, "table", TABLE_PATH, "path of table (if empty, use embedded table)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]... [FILE]\nIf FILE is empty or '-', read from standard input.\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	table := []rune{}
	if TABLE_PATH != "" {
		table_file, err := os.Open(TABLE_PATH)
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
	} else {
		for _, r := range DEFAULT_TABLE {
			table = append(table, r)
		}
	}

	type inputReader interface {
		io.Reader
		io.RuneReader
	}
	var input inputReader = bufio.NewReader(os.Stdin)
	if input_path := flag.Arg(0); input_path != "" && input_path != "-" {
		input_file, err := os.Open(input_path)
		if err != nil {
			log.Fatalln(err)
		}
		input = bufio.NewReader(input_file)
	}

	if !DECODE {
		encoder := NewEncoder(table)
		encoder.Encode(input, os.Stdout)
	} else {
		decoder := NewDecoder(table)
		decoder.Decode(input, os.Stdout)
	}
}
