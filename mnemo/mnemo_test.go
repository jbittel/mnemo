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

type validTest struct {
	input  string
	output bool
}

var encodeTests = []encodeTest{
	{-999, "xachozo"},
	{-100, "xabiba"},
	{-99, "xazo"},
	{-1, "xabi"},
	{-0, "ba"},
	{0, "ba"},
	{1, "bi"},
	{34, "jo"},
	{99, "zo"},
	{100, "biba"},
	{101, "bibi"},
	{999, "chozo"},
	{392406, "kogochi"},
	{25437225, "haleshuha"},
}

var decodeTests = []decodeTest{
	{"xachozo", -999},
	{"xabaji", -31},
	{"xabi", -1},
	{"xaba", 0},
	{"ba", 0},
	{"bi", 1},
	{"jo", 34},
	{"bajo", 34},
	{"bababababababajo", 34},
	{"chozo", 999},
	{"kogochi", 392406},
	{"yoshida", 947110},
	{"tonukatsu", 79523582},
}

var validTests = []validTest{
	{"xachozo", true},
	{"chozo", true},
	{"hozo", true},
	{"ozo", false},
	{"ozo", false},
	{"zo", true},
	{"o", false},
}

func TestEncode(t *testing.T) {
	for _, test := range encodeTests {
		if s, _ := Encode(test.input); s != test.output {
			t.Errorf("encoding %d returned %s, expected %s", test.input, s, test.output)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, test := range decodeTests {
		if i, _ := Decode(test.input); i != test.output {
			t.Errorf("decoding %s returned %d, expected %d", test.input, i, test.output)
		}
	}

	_, err := Decode("invalid")
	if err == nil {
		t.Errorf("no error returned when decoding invalid syllables")
	}
}

func TestValid(t *testing.T) {
	for _, test := range validTests {
		if v := Valid(test.input); v != test.output {
			t.Errorf("validating %s returned %t, expected %t", test.input, v, test.output)
		}
	}
}

func benchmarkEncode(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encode(i)
	}
}

func BenchmarkEncode9(b *testing.B)         { benchmarkEncode(9, b) }
func BenchmarkEncode999(b *testing.B)       { benchmarkEncode(999, b) }
func BenchmarkEncode999999(b *testing.B)    { benchmarkEncode(999999, b) }
func BenchmarkEncode999999999(b *testing.B) { benchmarkEncode(999999999, b) }

func benchmarkDecode(s string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Decode(s)
	}
}

func BenchmarkDecode9(b *testing.B)         { benchmarkDecode("cho", b) }
func BenchmarkDecode999(b *testing.B)       { benchmarkDecode("chozo", b) }
func BenchmarkDecode999999(b *testing.B)    { benchmarkDecode("zozozo", b) }
func BenchmarkDecode999999999(b *testing.B) { benchmarkDecode("chozozozozo", b) }
