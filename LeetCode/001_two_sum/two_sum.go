/*
1. Two Sum

Source: https://leetcode.com/problems/two-sum/

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

Note:
use hash map to instantly check for difference value,
map will add index of last occurrence of a num, don’t use same element twice;
*/

package twosum

// Time complexity: O(n)
// Space complexity: O(n)
func twoSum(nums []int, target int) []int {
	//create a map
	m := make(map[int]int)

	//loop the array nums from index 1 if not present in map we add this number
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		//check if the value of difference exist in the map
		if value, exist := m[diff]; exist {
			return []int{value, i}
		}
		//add to map if not present
		m[nums[i]] = i
	}

	return []int{}
}

// Brute Force Solution
// Time complexity: O(n^2)
// Space complexity: O(1)
func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
