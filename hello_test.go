package main

import "testing"

func TestMain(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Ram", "English")
		want := "Hello, Ram!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello, World! when empty string is given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Support Spanish", func(t *testing.T) {
		got := Hello("Krishna", "Spanish")
		want := "Hola, Krishna!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Support French", func(t *testing.T) {
		got := Hello("Krishna", "French")
		want := "Bonjour, Krishna!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
