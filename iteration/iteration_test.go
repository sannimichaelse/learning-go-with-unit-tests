package iteration

import "testing"

func TestIteration(t *testing.T) {
	repeated := Iterate("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("a")
	}
}
