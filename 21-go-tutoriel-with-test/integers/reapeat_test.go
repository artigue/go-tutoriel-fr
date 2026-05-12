package integers

import "testing"

func Repeat(s string, count int) string {
	var repeated string= ""
	for i := 0; i < count; i++ {
		repeated += s
	}
	return repeated
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}	