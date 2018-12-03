package main

import (
	"fmt"

	"github.com/jbittel/mnemo/mnemo"
)

func main() {
	fmt.Println(mnemo.Fmne_to_s(98765))
	fmt.Println(mnemo.Fmne_to_s(4567))
	fmt.Println(mnemo.Fmne_to_s(-4567))

	fmt.Println(mnemo.Fmne_to_i("tonukatsu"))
}
