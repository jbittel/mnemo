package main

import (
	"flag"
	"fmt"

	"github.com/jbittel/mnemo/mnemo"
)

func main() {
	var encode int
	var decode string

	flag.IntVar(&encode, "encode", 0, "Encode an integer to a string")
	flag.StringVar(&decode, "decode", "", "Decode a string to an integer")
	flag.Parse()

	if encode != 0 {
		fmt.Println(mnemo.Encode(encode))
	} else if decode != "" {
		fmt.Println(mnemo.Decode(decode))
	}
}
