package containsduplicate

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	numsTrue := []int{1, 2, 4, 56, 7, 8, 2}
	numsFalse := []int{1, 2, 4, 56, 7, 8}

	funcs := []func([]int) bool{containsDuplicate}

	for _, testFunc := range funcs {
		if res := testFunc(numsTrue); !reflect.DeepEqual(res, true){
			t.Error("Fail Test!")
		}

		if res := testFunc(numsFalse); !reflect.DeepEqual(res, false){
			fmt.Println(res)
			t.Error("Fail Test!")
		}
	}
}
