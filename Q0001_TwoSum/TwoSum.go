package Q0001_TwoSum

/*
https://leetcode.com/problems/two-sum/

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
 */

func twoSum(nums []int, target int) []int {
	counterMap := map[int]int{}
	for i, v := range nums {
		pairIndex, ok := counterMap[v]
		if ok {
			answer := []int{pairIndex, i}
			return answer
		}
		counterValue := target - v
		counterMap[counterValue] = i
	}
	return []int{-1, -1}
}