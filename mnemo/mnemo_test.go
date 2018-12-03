package mnemo

import "testing"

type tosTest struct {
	input  int
	output string
}

type toiTest struct {
	input  string
	output int
}

var tosTests = []tosTest{
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

var toiTests = []toiTest{
	{"yoshida", 947110},
	{"bajo", 34},
	{"tonukatsu", 79523582},
	{"xabaji", -31},
}

func TestFmne_to_s(t *testing.T) {
	for _, test := range tosTests {
		if s := Fmne_to_s(test.input); s != test.output {
			t.Errorf("Input %d returned %s, expected %s", test.input, s, test.output)
		}
	}
}

func TestFmne_to_i(t *testing.T) {
	for _, test := range toiTests {
		if i := Fmne_to_i(test.input); i != test.output {
			t.Errorf("Input %s returned %d, expected %d", test.input, i, test.output)
		}
	}
}
