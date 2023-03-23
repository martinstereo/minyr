package yr

import (
	"reflect"
	"testing"
)

/*
	Tester som skal være implementert:

	* antall linjer i filen er 16756
	* gitt "Kjevik;SN39040;18.03.2022 01:50;6" ønsker å få (want) "Kjevik;SN39040;18.03.2022 01:50;42,8"
	* gitt "Kjevik;SN39040;07.03.2023 18:20;0" ønsker å få (want) "Kjevik;SN39040;07.03.2023 18:20;32"
	* gitt "Kjevik;SN39040;08.03.2023 02:20;-11" ønsker å få (want) "Kjevik;SN39040;08.03.2023 02:20;12,2"
	* gitt "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;" ønsker å få (want) "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av STUDENTENS_NAVN", hvor STUDENTENS_NAVN er navn på studenten som leverer besvarelsen
	* er gjennomsnittstemperatur 8.56?
	* FOR EKSTRA OPPGAVE: gjennomsnittstemperatur for august 2022 i grader Celsius er ...
*/

func TestCountLines(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	var tests = []test{
		{input: "../kjevik-temp-celsius-20220318-20230318.csv",
			want: 16756},
	}
	for _, tc := range tests {
		got := CountLines(tc.input)

		if got != tc.want {
			t.Errorf("%v: want %v, got %v,", tc.input, tc.want, got)
		}
	}
}
func TestReplaceCelsiusInFile(t *testing.T) {

	//input og resultat ønsket
	type test struct {
		input string
		want  string
	}

	var tests = []test{
		{input: "Kjevik;SN39040;18.03.2023 00:20;4", want: "Kjevik;SN39040;18.03.2023 00:20;39.2"}, // Celsius to Fahr
		{input: "Kjevik;SN39040;17.03.2023 18:50;3", want: "Kjevik;SN39040;17.03.2023 18:50;37.4"},
		{input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
		{input: "Kjevik;SN39040;08.03.2023 02:20;0", want: "Kjevik;SN39040;08.03.2023 02:20;32"},
	}
	for _, tc := range tests {
		got := ConvertCelsiusToFahr(tc.input)

		if !reflect.DeepEqual(tc.want, got) {

			t.Errorf("%s: want %s, got %s", tc.input, tc.want, got)
		}
	}
}

func TestReplaceEndLine(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	var tests = []test{
		{input: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;",
			want: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);endringen er gjort av Martin Steiro"},
	}
	for _, tc := range tests {
		got := EditEndLine(tc.input)

		if got != tc.want {
			t.Errorf("%v: want %v, got %v,", tc.input, tc.want, got)
		}
	}
}

func TestAverageTemp(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	var tests = []test{
		{input: "../kjevik-temp-celsius-20220318-20230318.csv",
			want: "8.56"},
	}
	for _, tc := range tests {
		got := AverageTemp(tc.input)

		if got != tc.want {
			t.Errorf("%v: want %v, got %v,", tc.input, tc.want, got)
		}
	}
}
