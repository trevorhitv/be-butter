package linearsearch

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	if LinearSearch(foo, 4) != true {
		t.Errorf("wanted true")
	}
	if LinearSearch(foo, 69) != true {
		t.Errorf("wanted true")
	}
	if LinearSearch(foo, 1336) != false {
		t.Errorf("wanted false")
	}
	if LinearSearch(foo, 69420) != true {
		t.Errorf("wanted true")
	}
	if LinearSearch(foo, 69421) != false {
		t.Errorf("wanted false")
	}
	if LinearSearch(foo, 1) != true {
		t.Errorf("wanted true")
	}
	if LinearSearch(foo, 0) != false {
		t.Errorf("wanted false")
	}

}
