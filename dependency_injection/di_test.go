package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("test greet", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Ram")

		got := buffer.String()
		want := "Hello, Ram\n"

		assertEquals(t, got, want)
	})
}

func assertEquals(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
