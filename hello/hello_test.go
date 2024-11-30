package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chad", "English")
		want := "Hello, Chad"

		assertCorrentMessage(t, got, want)
	})

	t.Run("say 'Hello, World'", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrentMessage(t, got, want)
	})

	t.Run("in Turkish", func(t *testing.T) {
		got := Hello("Chad", "Turkish")
		want := "Merhaba, Chad"
		assertCorrentMessage(t, got, want)
	})
}

func assertCorrentMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
