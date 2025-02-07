package containsduplicate

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	numsTrue := []int{1, 2, 4, 56, 7, 8, 2}
	numsFalse := []int{1, 2, 4, 56, 7, 8}

	funcs := []func([]int) bool{containsDuplicateBF, containsDuplicateSorting, containsDuplicateMap}

	for _, testFunc := range funcs {
		if res := testFunc(numsTrue); !reflect.DeepEqual(res, true) {
			t.Error("Fail Test!")
		}

		if res := testFunc(numsFalse); !reflect.DeepEqual(res, false) {
			fmt.Println(res)
			t.Error("Fail Test!")
		}
	}

}

func TestContainsDuplicateV2(t *testing.T) {
	testsCases := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 3},
	}

	expected := []bool{false, true}

	funcs := []func([]int) bool{containsDuplicateBF, containsDuplicateSorting, containsDuplicateMap}

	for index, nums := range testsCases {
		for _, testFunc := range funcs {
			if res := testFunc(nums); res != expected[index] {
				t.Errorf("expected %t, got %t", expected[index], res)
			}
		}
	}
}
