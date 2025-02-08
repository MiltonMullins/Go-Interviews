package productexceptself

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	testCases := [][]int{
		{1,2,3,4},
		{-1,1,0,-3,3},
	}

	expected := [][]int {
		{24,12,8,6},
		{0,0,9,0,0},
	}

	for index, nums := range testCases {
		if res := productExceptSelf(nums); !reflect.DeepEqual(res,expected[index]) {
			t.Errorf("\nexpected %v, got %v", expected[index], res)
		}
	}
}
