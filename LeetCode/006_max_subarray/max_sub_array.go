/*
Given an integer array nums, find the
subarray with the largest sum, and return its sum.

Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Example 2:
Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.

Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.

Tip: pattern: prev subarray cant be negative, dynamic programming: compute max sum for each prefix
*/

package maxsubarray

//Time Complexity: O(n)
//Space Complexity: O(1)
func maxSubArray(nums []int) int {
	//declare variables
	maxSum, sum := nums[0], 0

	for _, num := range nums {
		sum += num
		if sum > maxSum {
			maxSum = sum
		}
		//If sum becomes negative, reset it to zero because a negative sum cannot contribute to the maximum.
		if sum < 0 {
			sum = 0
		}
	}

	return maxSum
}