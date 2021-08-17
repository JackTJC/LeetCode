package bit

import (
	"testing"
)

func TestSgNum(t *testing.T) {
	t.Logf("result = %v", singleNumber([]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 9, 9, 9, 999}))
	t.Logf("result = %v", singleNumber([]int{111, 111, 111, 12, 12, 12, 3}))
	t.Logf("result = %v", singleNumber([]int{0, 1, 1, 1, 2, 2, 2}))
}
