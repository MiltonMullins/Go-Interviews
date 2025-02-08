/*
Given an integer array nums, return an array answer such that answer[i] is equal to 
the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

*/

package productexceptself

//Brute forece Solution
//Time Complexity: O(n^2)
//Space Complexity: O(1)
func productExceptSelf(nums []int) []int {
    products, curProd := []int{},1

	for i := range nums {
		curProd = 1
		for j, num2 := range nums {
			if j != i {
				curProd *= num2
			}
		}
		products = append(products, curProd)
	}

	return products
}