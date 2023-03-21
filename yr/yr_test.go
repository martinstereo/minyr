package yr

import "testing"

func TestReplaceCelsiusInFile(t *testing.T) {

	//input og resultat ønsket
	type test struct {
		input string
		want  string
	}

	var tests = []test{
		{input: "Kjevik;SN39040;18.03.2023 00:20;4", want: "Kjevik;SN39040;18.03.2023 00:20;39.2"}, // Celsius to Fahr
		{input: "Kjevik;SN39040;17.03.2023 18:50;3", want: "Kjevik;SN39040;17.03.2023 18:50;37.4"},
		// .... flere tester
	}
	for _, tc := range tests {
		got := ReplaceCelsiusInFile(tc.input)

		if got != tc.want {
			t.Errorf("%s: want %s, got %s,", tc.input, tc.want, got)
		}
	}
}
