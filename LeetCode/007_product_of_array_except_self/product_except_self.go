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

// Brute forece Solution
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func productExceptSelf(nums []int) []int {
	products, curProd := []int{}, 1

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
// if the count of zeros are 0, which means all the elements in the nums array are non-zero 
// and output of each element of the nums array in the ans array will be equal to product / nums[i].
func productExceptSelf2(nums []int) []int {
	n := len(nums)
    ans := make([]int, n)
    product := 1
    zeros := 0

    for _, num := range nums {
        if num == 0 {
            zeros++
            continue
        }
        product *= num
    }

    if zeros == 1 {
        for i, num := range nums {
            if num == 0 {
                ans[i] = product
            }
        }
    } else if zeros == 0 {
        for i, num := range nums {
            ans[i] = product / num //without using the division operation. This isn't allowed
        }
    }

    return ans
}
