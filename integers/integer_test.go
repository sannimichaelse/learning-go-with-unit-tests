package integers

import "testing"

func TestInteger(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("it should add two numbers up", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		assertCorrectMessage(t, got, want)
	})

	t.Run("it should subtract two numbers", func(t *testing.T) {
		got := Sub(10, 2)
		want := 8

		assertCorrectMessage(t, got, want)
	})

	t.Run("it should Multiply two numbers", func(t *testing.T) {
		got := Mult(10, 20)
		want := 200

		assertCorrectMessage(t, got, want)
	})

	t.Run("it should divide two numbers", func(t *testing.T) {
		got := Div(15, 3)
		want := 5

		assertCorrectMessage(t, got, want)
	})
}
