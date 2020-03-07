package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jbittel/mnemo/mnemo"
)

type intFlag struct {
	set   bool
	value int
}

func (flag *intFlag) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	flag.set = true
	flag.value = int(v)
	return err
}

func (flag *intFlag) String() string {
	return strconv.Itoa(int(flag.value))
}

type stringFlag struct {
	set   bool
	value string
}

func (flag *stringFlag) Set(s string) error {
	flag.set = true
	flag.value = s
	return nil
}

func (flag *stringFlag) String() string {
	return flag.value
}

var encode intFlag
var decode stringFlag
var validate stringFlag

func init() {
	flag.Var(&encode, "encode", "Encode an integer to a word")
	flag.Var(&decode, "decode", "Decode a word to an integer")
	flag.Var(&validate, "validate", "Validate a word")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n\n  %s [-encode | -decode | -validate]\n\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if encode.set {
		s, _ := mnemo.Encode(encode.value)
		fmt.Println(s)
	} else if decode.set {
		i, err := mnemo.Decode(decode.value)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(i)
		}
	} else if validate.set {
		v := mnemo.Validate(validate.value)
		fmt.Println(v)
	} else {
		flag.Usage()
	}
}
