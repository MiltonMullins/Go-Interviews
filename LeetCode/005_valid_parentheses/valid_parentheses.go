/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "([{}])"
Output: true

Tip: push opening brace on stack, pop if matching close brace, at end if stack empty, return true;
*/

package validparentheses

//Stack Solution
//Time Complexity: O(n)
//Space Complexity: O(n)
func validParentheses(s string) bool {
	stack := []rune{}
	hash := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, char := range s {
		if match, found := hash[char]; found {
			// Check if stack is non-empty and matches
			if len(stack) > 0 && stack[len(stack)-1] == match {
				stack = stack[:len(stack)-1] //Pop ->  len-1 == last position and ":" operator exclude the last one
			} else {
				return false //invalid
			}
		} else {
			stack = append(stack, char) //Push
		}
	}

	return len(stack) == 0
}
