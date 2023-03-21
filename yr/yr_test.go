package yr_test

import "testing"

func TestYr(t *testing.T) {

	//input og resultat ønsket
	type test struct {
		input int
		want  int
	}

	var tests = []test{
		{input  , want },
		{1, 1},
		{2, 1},
		{3, 2},
		// .... flere tester
	}
	for _, tc := range tests {
		got := .....(tc.input)

		if got != tc.want {
			t.Errorf("%d: want %d, got %d,", tc.input, tc.want, got)
		}
	}
}
