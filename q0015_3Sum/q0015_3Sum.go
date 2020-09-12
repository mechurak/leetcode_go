package q0015_3Sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	answerMap := map[[3]int]bool{}

	for i := 0; i < len(nums) -2 ; i++ {
		v := nums[i]
		target := -v
		counterMap := map[int]int{}
		tempNums := nums[i+1:]
		//fmt.Printf("target: %v, counterMap: %v, tempNums: %v\n", target, counterMap, tempNums)
		for _, w := range tempNums {
			pairValue, ok := counterMap[w]
			if ok {
				answer := [3]int{v, pairValue, w}
				//fmt.Printf("answer: %v\n", answer)
				_, ok := answerMap[answer]
				if !ok {
					answerMap[answer] = true
					//fmt.Printf("Add it to map: %v\n", answer)
				} else {
					//fmt.Printf("Already exist. answer: %v\n", answer)
				}
			}
			counterValue := target - w
			counterMap[counterValue] = w
		}
	}

	if len(answerMap) == 0 {
		return [][]int{}
	}

	var answerArray [][]int
	for key := range answerMap {
		answerArray = append(answerArray, []int{key[0], key[1], key[2]})
	}
	return answerArray
}
