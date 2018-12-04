package main

import (
	"flag"
	"fmt"
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

func init() {
	flag.Var(&encode, "encode", "Encode an integer to a string")
	flag.Var(&decode, "decode", "Decode a string to an integer")
}

func main() {
	flag.Parse()

	if encode.set {
		fmt.Println(mnemo.Encode(encode.value))
	} else if decode.set {
		fmt.Println(mnemo.Decode(decode.value))
	}
}
