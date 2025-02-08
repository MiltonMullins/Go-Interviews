package maxprofit

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	testPrices := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
		{},
		{1},
		{5, 3, 2, 2, 5, 7, 9, 4, 5, 3, 2, 2, 5, 7, 9, 4},
		{2, 4, 1, 11, 7},
	}

	expected := []int{5, 0, 0, 0, 7, 10}

	funcs := []func([]int) int{ maxProfitBF, maxProfitD1, maxProfit}

	for index, prices := range testPrices {
		for _, testFunc := range funcs {
			if res := testFunc(prices); res != expected[index] {
				t.Errorf("expected %v, got %v", expected[index], res)
			}
		}
	}
}
