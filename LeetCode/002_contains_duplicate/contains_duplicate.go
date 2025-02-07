/*
217. Contains Duplicate
https://leetcode.com/problems/contains-duplicate/

Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array,
and it should return false if every element is distinct.
*/
package containsduplicate

//Brute Force Solution
//Time Complexity: O(n^2)
//Space Complexity: O(1)
func containsDuplicate(nums []int) bool{
	//loop over array twice and compare
	for i, v := range nums {
		for j := i+1; j < len(nums); j++ {
			if v == nums[j] {
				return true
			}
		}
	}
	return false
}