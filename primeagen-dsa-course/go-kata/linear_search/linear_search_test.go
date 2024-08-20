package linearsearch

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	got := LinearSearch(foo, 4)
	fmt.Printf("result: %t", got)
	if got != true {
		t.Errorf("Abs(-1) = %t; want 1", got)
	}
}
