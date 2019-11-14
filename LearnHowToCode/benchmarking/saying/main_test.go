package saying

import (
	"fmt"
	"testing"
)

func TestGreeting(t *testing.T) {
	s := Greet("tom")
	if s != "Welcome my dear tom" {
		t.Error("got", s, "expected", "Welcome my dear tom")
	}
}

//  godoc -http=:8080
func ExampleGreet() {
	fmt.Println(Greet("tom"))
	// Output:
	// Welcome my dear tom
}

// go test -bench .
// コードの実行速度を図る
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("tom")
	}
}

// Coverage
//  go test -cover
// どれくらいテストされているかを図る。
// 100%になることはまずない。
