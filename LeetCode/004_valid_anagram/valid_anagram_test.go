package validanagram

import	"testing"

func TestValidAnagram(t *testing.T) {
	s1 := "anagram"
	s2 := "nagaram"
	s3 := "sds"
	s4 := "sxs"

	if !validAnagram(s1,s2) {
		t.Error("Test Fail")
	}

	if validAnagram(s4,s3) {
		t.Error("Test Fail")
	}
}
