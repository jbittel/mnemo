package main

import (
	"bytes"
	"fmt"
)

var fmne_syls = [...]string{
	"ba", "bi", "bu", "be", "bo",
	"cha", "chi", "chu", "che", "cho",
	"da", "di", "du", "de", "do",
	"fa", "fi", "fu", "fe", "fo",
	"ga", "gi", "gu", "ge", "go",
	"ha", "hi", "hu", "he", "ho",
	"ja", "ji", "ju", "je", "jo",
	"ka", "ki", "ku", "ke", "ko",
	"la", "li", "lu", "le", "lo",
	"ma", "mi", "mu", "me", "mo",
	"na", "ni", "nu", "ne", "no",
	"pa", "pi", "pu", "pe", "po",
	"ra", "ri", "ru", "re", "ro",
	"sa", "si", "su", "se", "so",
	"sha", "shi", "shu", "she", "sho",
	"ta", "ti", "tu", "te", "to",
	"tsa", "tsi", "tsu", "tse", "tso",
	"wa", "wi", "wu", "we", "wo",
	"ya", "yi", "yu", "ye", "yo",
	"za", "zi", "zu", "ze", "zo",
}

const fmne_neg = "xa"

func fmne_tos(i int, b *bytes.Buffer) {
	mod := i % len(fmne_syls)
	rst := i / len(fmne_syls)
	if rst > 0 {
		fmne_tos(rst, b)
	}
	b.WriteString(fmne_syls[mod])
}

func fmne_to_s(i int) string {
	var b bytes.Buffer

	if i < 0 {
		b.WriteString(fmne_neg)
		i = i * -1
	}

	fmne_tos(i, &b)

	return b.String()
}

func main() {
	fmt.Println(fmne_to_s(98765))
	fmt.Println(fmne_to_s(4567))
	fmt.Println(fmne_to_s(-4567))
}