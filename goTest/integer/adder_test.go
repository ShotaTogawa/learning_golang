package integer

import "testing"

func Add(x, y int) (sum int) {
	sum = x + y
	return
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
