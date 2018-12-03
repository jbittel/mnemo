package mnemo

import "testing"

type tosTest struct {
	input  int
	output string
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

func TestFmne_to_s(t *testing.T) {
	for _, test := range tosTests {
		if s := Fmne_to_s(test.input); s != test.output {
			t.Errorf("Input %d returned %s, expected %s", test.input, s, test.output)
		}
	}
}
