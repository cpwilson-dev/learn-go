package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat default number of times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertMessage(t, got, want)
	})
	t.Run("repeat 3 times", func(t *testing.T) {
		got := Repeat("a", 3)
		want := "aaa"
		assertMessage(t, got, want)
	})

}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q but got %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	res := Repeat("a", 3)
	fmt.Println(res)
	//Output: aaa
}
