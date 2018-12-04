package main

import (
	"fmt"
	//"os"

	"github.com/jbittel/mnemo/mnemo"
)

func main() {
	arg := 25437225
	fmt.Println(mnemo.Encode(arg))
}
