package maxsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	testCases := [][]int{
		{5, 4, -1, 7, 8},
		{1},
		{-2,1,-3,4,-1,2,1,-5,4},
	}

	expected := []int{23, 1, 6}

	for index, nums := range testCases {
		if res := maxSubArray(nums); res != expected[index] {
			t.Errorf("\nexpected %v, got %v", expected[index], res)
		}
	}
}
