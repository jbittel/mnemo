// Package mnemo converts integers into mnemonic
// "words" and back again.
package mnemo

import (
	"bytes"
	"errors"
	"strings"
)

var syllables = [...]string{
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

const negSyllable = "xa"

func encodeString(i int, buf *bytes.Buffer) {
	mod := i % len(syllables)
	rst := i / len(syllables)

	if rst > 0 {
		encodeString(rst, buf)
	}

	buf.WriteString(syllables[mod])
}

// Encode converts an integer into a mnemonic "word"
func Encode(i int) (string, error) {
	var buf bytes.Buffer

	if i < 0 {
		buf.WriteString(negSyllable)
		i = i * -1
	}

	encodeString(i, &buf)

	return buf.String(), nil
}

// Decode converts a "word" into it's integer equivalent
func Decode(s string) (int, error) {
	result := 0
	sign := 1

	if strings.HasPrefix(s, negSyllable) {
		sign = -1
		s = strings.TrimPrefix(s, negSyllable)
	}

DecodeLoop:
	for len(s) > 0 {
		for i, syllable := range syllables {
			if strings.HasPrefix(s, syllable) {
				s = strings.TrimPrefix(s, syllable)
				result = len(syllables)*result + i
				continue DecodeLoop
			}
		}

		return 0, errors.New("invalid syllable encountered")
	}

	return sign * result, nil
}
