package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("check given char is repeated", func(t *testing.T) {
		got := Repeat("a", 3)
		want := "aaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("d", 100)
	}
}

func BenchmarkRepeat2(b *testing.B) {
	for b.Loop() {
		Repeat2("d", 100)
	}
}

func ExampleRepeat() {
	got := Repeat("a", 3)
	fmt.Println(got)
	//Output: aaa
}
