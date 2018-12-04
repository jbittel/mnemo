package mnemo

import "testing"

type encodeTest struct {
	input  int
	output string
}

type decodeTest struct {
	input  string
	output int
}

var encodeTests = []encodeTest{
	{0, "ba"},
	{1, "bi"},
	{99, "zo"},
	{100, "biba"},
	{101, "bibi"},
	{392406, "kogochi"},
	{25437225, "haleshuha"},
	{-1, "xabi"},
	{-99, "xazo"},
	{-100, "xabiba"},
}

var decodeTests = []decodeTest{
	{"yoshida", 947110},
	{"bajo", 34},
	{"kogochi", 392406},
	{"tonukatsu", 79523582},
	{"xabaji", -31},
}

func TestEncode(t *testing.T) {
	for _, test := range encodeTests {
		if s := Encode(test.input); s != test.output {
			t.Errorf("Input %d returned %s, expected %s", test.input, s, test.output)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, test := range decodeTests {
		if i := Decode(test.input); i != test.output {
			t.Errorf("Input %s returned %d, expected %d", test.input, i, test.output)
		}
	}
}
