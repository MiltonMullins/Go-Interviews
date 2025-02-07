/*
217. Contains Duplicate
https://leetcode.com/problems/contains-duplicate/

Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array,
and it should return false if every element is distinct.
*/
package containsduplicate

import (
	"slices"
)

//Brute Force Solution
//Time Complexity: O(n^2)
//Space Complexity: O(1)
func containsDuplicateBF(nums []int) bool{
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

//Sorting Solution
//Time Complexity: O(n log n)
//Space Complexity: O(1)
func containsDuplicateSorting(nums []int) bool {
	//soting the array you only need to loop one time and compare adjacetns positions
	slices.Sort(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

//Map Solution
//Time Complexity: O(n)
//Space Complexity: O(1)
func containsDuplicateMap(nums []int) bool {
	//create map
	mNums := make(map[int]int)

	//loop nums and add to map if not exist
	for _, v := range nums {
		if _, ok := mNums[v]; ok {
			return true
		}
		mNums[v] = v
	}

	return false
}