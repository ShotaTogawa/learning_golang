package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func Repeat(char string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += char
	}
	return repeated
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func Sum(nums ...int) int {
	var total int
	for _, v := range nums {
		total += v
	}
	return total
}

func TestSum(t *testing.T) {
	want := 15
	expected := Sum(1,2,3,4,5)

	if want != expected {
		t.Errorf("expected %d, but got %d", expected, want)
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1,2,3,4,5)
	}
}