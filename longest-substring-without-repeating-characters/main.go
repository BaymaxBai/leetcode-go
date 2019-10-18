package main

import "fmt"

func main() {
	fmt.Printf("%d\n", lengthOfLongestSubstring("a"))
}

// 动态规划方法解题 利用空间换时间
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	posOfChar := make(map[rune]int) // 动态规划 map
	startPos := 0                   // 起始下标

	for i, c := range s {
		// 当字符再次出现时 触发计算子串长度 更新起始下标值(不能小于当前下标值)
		if pos, ok := posOfChar[c]; ok == true && pos >= startPos {
			// 当前下标 - 起始下标 = 当前子串的长度
			if i-startPos > maxLength {
				maxLength = i - startPos
			}
			// 起始下标前移一位
			startPos = pos + 1
		}
		// 记录当前字符的下标位置
		posOfChar[c] = i
	}

	// 处理特殊情况子串等于串的时候 example: abcde
	if len(s)-startPos > maxLength {
		maxLength = len(s) - startPos
	}

	return maxLength
}
