/*
242. Valid Anagram
Given two strings s and t, return true if t is an
anagram of s, and false otherwise.

Example 1:

Input: s = "anagram", t = "nagaram"

Output: true

Example 2:

Input: s = "rat", t = "car"

Output: false

Tip: hashmap to count each char in str1, decrement for str2;
*/
package validanagram

// Map Solution
// Time Complexity: O(n)
// Space Complexity: O(1)
func validAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	//creat data structure map
	dictionary := make(map[rune]int)

	//loop the strings s and add to a map if present
	for _, char := range s {
		dictionary[char]++
	}

	//variable of control
	cantLetters := len(dictionary)
	//loop the strings t and dreceasse value if key present
	for _, char := range t {
		if _, ok := dictionary[char]; ok {
			dictionary[char]--
		} else {
			return false
		}

		if dictionary[char] < 0 {
			return false
		}

		if dictionary[char] == 0 {
			cantLetters--
		}
	}

	return cantLetters == 0
}
