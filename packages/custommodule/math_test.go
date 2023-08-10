package custommodule

import "testing"

func TestMathAdd(t *testing.T) {
	sum := Add(1, 2)
	expected := 3
	if (sum != expected) {
		t.Fatalf(`Error`)
	}
	
}