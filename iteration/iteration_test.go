package iteration

import (
	"fmt"
	"strings"
	"testing"
)

const repeatCount = 5

func TestRepeat(t *testing.T) {

	t.Run("repeat 'a' five times", func(t *testing.T) {
		t.Helper()
		repeated := Repeat("a", repeatCount)
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("compare 'Repeat()' and 'strings.Repeat()'", func(t *testing.T) {
		t.Helper()
		customRepeat := Repeat("a", repeatCount)
		stringsRepeat := strings.Repeat("a", repeatCount)
		if customRepeat != stringsRepeat {
			t.Errorf("Repeat: %q not equal strings.Repeat: %q", customRepeat, stringsRepeat)
		}
	})
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
