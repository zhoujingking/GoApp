package custommodule_test

import "testing"
import "goapp/packages/custommodule"

func TestMathAdd(t *testing.T) {
	sum := custommodule.Add(1, 2)
	expected := 3
	if (sum != expected) {
		t.Fatalf(`Error`)
	}
	
}