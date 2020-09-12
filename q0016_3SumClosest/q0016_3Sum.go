package q0016_3SumClosest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	ans := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				//fmt.Printf("%v + %v + %v = %v\n", nums[i], nums[j], nums[k], sum)
				if sum == target {
					return sum
				}
				diff := target - sum
				if diff < 0 {
					diff = -diff
				}
				ansDiff := target - ans
				if ansDiff < 0 {
					ansDiff = -ansDiff
				}
				if diff < ansDiff {
					//fmt.Printf("diff(%v) < ansDiff(%v). %v\n", diff, ansDiff, sum)
					ans = sum
				}
			}
		}
	}
	return ans
}
