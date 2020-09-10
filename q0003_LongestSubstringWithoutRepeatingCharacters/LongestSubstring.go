package q0003

/*
Given a string, find the length of the longest substring without repeating characters.

== Example 1:
Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

== Example 2:
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

== Example 3:
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	strLen := len(s)
	if strLen == 0 {
		return 0
	}
	curPoint := 1
	maxPoint := 1
	left := 0
	indexMap := map[byte]int{}
	indexMap[s[0]] = 0
	for i := 1; i < strLen; i++ {
		// fmt.Printf("cur char s[%d]:%c", i, s[i])
		prevIndex, ok := indexMap[s[i]]
		if ok {
			if prevIndex < left {
				// fmt.Printf("1.prevIndex:%d, left:%d, ignore", prevIndex, left)
				indexMap[s[i]] = i
				curPoint++
				// fmt.Printf("1.adjust curPoint:%d", curPoint)
			} else {
				// fmt.Printf("2.Repeat!! %c, prevIndex:%d, curIndex:%d (curPoint:%d, maxPoint:%d)\n", s[i], prevIndex, i, curPoint, maxPoint)
				left = prevIndex + 1
				// fmt.Printf("2.adjust left:%d", left)
				if curPoint > maxPoint {
					maxPoint = curPoint
				}
				curPoint = i - left + 1
				// fmt.Printf("2.adjust curPoint:%d", curPoint)
				indexMap[s[i]] = i
			}
		} else {
			indexMap[s[i]] = i
			curPoint++
			// fmt.Printf("3.adjust curPoint:%d", curPoint)
		}
	}
	if curPoint > maxPoint {
		maxPoint = curPoint
	}
	return maxPoint
}
