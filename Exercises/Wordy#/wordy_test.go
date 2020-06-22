package wordy

import (
	"strings"
	"testing"
	"fmt"
)

func TestAnswer(t *testing.T) {
	for _, test := range tests {
		switch ok, answer := calculate(test.question); {
		case !ok:
			if test.ok {
				t.Fatalf("FAIL: %s\nAnswer(%q)\nreturned ok = false, expecting true.", test.description, test.question)
			}
		case !test.ok:
			t.Errorf("FAIL: %s\nAnswer(%q)\nreturned %d, %t, expecting ok = false.", test.description, test.question, answer, ok)
		case answer != test.answer:
			t.Errorf("FAIL: %s\nAnswer(%q)\nreturned %d, expected %d.", test.description, test.question, answer, test.answer)
		}
		t.Logf("PASS: %s", test.description)
	}
}

func TestRemoveQuestionMark(t *testing.T) {
	test := "hello?"
	fmt.Println(strings.TrimRight(test, "?"))
}
