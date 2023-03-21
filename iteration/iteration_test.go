package iteration

import (
	"fmt"
	"testing"
)

const repeatCount = 5

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", repeatCount)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", repeatCount)
	}
}

func ExampleRepeat() {
	repeated := Repeat("p", 4)
	fmt.Println(repeated)
	// Output: pppp
}
