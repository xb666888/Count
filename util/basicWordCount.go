package util

// computeTemporaryArray 计算KMP算法中的临时数组
func computeTemporaryArray(pattern string) []int {
	temp := make([]int, len(pattern))
	j := 0
	for i := 1; i < len(pattern); {
		if pattern[i] == pattern[j] {
			temp[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = temp[j-1]
		} else {
			temp[i] = 0
			i++
		}
	}
	return temp
}

// GetWordCount 使用KMP算法统计模式在文本中出现的次数
func GetWordCount(text, pattern string) int {
	temp := computeTemporaryArray(pattern)
	count := 0
	i, j := 0, 0

	for i < len(text) {
		if text[i] == pattern[j] {
			i++
			j++
		}

		if j == len(pattern) { // 匹配成功
			count++
			j = temp[j-1] // 复位j以继续查找下一个匹配
		} else if i < len(text) && text[i] != pattern[j] {
			if j > 0 {
				j = temp[j-1]
			} else {
				i++
			}
		}
	}

	return count
}
