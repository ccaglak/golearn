package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"
	if repeated != expected {
		t.Errorf("expected %s but got %s", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("A", 100)
	}
}
