package funcfile

import (
	"testing"
)

func TestRunMyTime(t *testing.T) {
	RunMyTime()

}

func BenchmarkRunMyTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunMyTime()
	}
}

func ExampleRunMyTime() {
	RunSayHello()
	// output:
	// [hello world]
}
