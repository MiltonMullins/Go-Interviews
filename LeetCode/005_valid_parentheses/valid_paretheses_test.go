package validparentheses

import "testing"

func TestValidParetheses(t * testing.T){

	testData := []string{"()", "(((([{}]))))", "()()()()", "(((((((()", "((()(())))", "", ")(", "}{", "][", "(][)"}

	expectedData := []bool{true, true, true, false, true, true, false, false, false, false}

	for i, s := range testData {
		result := validParentheses(s)
		if result != expectedData[i] {
			t.Error("Test Failed!")
		}
	}
}