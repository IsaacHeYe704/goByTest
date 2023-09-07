package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeat := Repeat("a", 5)
	expected := "aaaaa"
	if repeat != expected {
		t.Errorf("Exxpected %q but got %q", expected, repeat)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
