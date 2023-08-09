package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	chs := []byte(s)
	lastIndexMap := make(map[byte]int)
	start := 0
	result := 0
	for i, c := range chs {
		if index, ok := lastIndexMap[c]; ok {
			if index >= start {
				start = index + 1
			}
		}
		lastIndexMap[c] = i
		// 计算最大长度
		curLength := i - start + 1
		if curLength > result {
			result = curLength
		}
	}
	return result
}

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcbb"),
		lengthOfNonRepeatingSubStr("bbbbbb"),
		lengthOfNonRepeatingSubStr("pwwkew"),
	)
}
