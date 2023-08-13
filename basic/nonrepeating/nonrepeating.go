package nonrepeating

func LengthOfNonRepeatingSubStr(s string) int {
	chs := []rune(s)
	lastIndexMap := make([]int, 0xffff)
	for i := range lastIndexMap {
		lastIndexMap[i] = -1
	}
	start := 0
	result := 0

	for i, c := range chs {
		if index := lastIndexMap[c]; index > -1 && index >= start {
			start = index + 1
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
