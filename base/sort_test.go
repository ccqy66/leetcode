package base

import (
	"fmt"
	"testing"
)

func TestGetTopK(t *testing.T) {
	fmt.Println(getTopK([]int{2,3,4,1,6,5}, 0, 6, 1))
}
