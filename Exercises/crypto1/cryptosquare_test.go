package cryptosquare

import "testing"

var tests = []struct {
	plain string // plain text
	normalized string // cipher text
}{
	{
		"My name is Ade",
		"MynameisAde",
	},
	{
		"My name is Ade Baruwa a bo",
		"MynameisAdeBaruwaabo",
	},
	{
		"My name is Ade Baruwa, a bo",
		"MynameisAdeBaruwaabo",
	},
	{
		"My name is: Ade Baruwa; a, bo",
		"MynameisAdeBaruwaabo",
	},
}

func TestNoSpaces(t *testing.T) {
	for _, test := range tests {
		if normalized := remover(test.plain); normalized != test.normalized {
			t.Fatalf(`remover(%q);
got %q
want %q`, test.plain, normalized, test.normalized)

		}

	}
}
