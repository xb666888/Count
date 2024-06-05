package util

import (
	"regexp"
	"strings"
)

// Tokenize 分词函数，处理英文大小写和标点符号
func Tokenize(input string) []string {
	// 正则表达式匹配单词和标点符号
	re := regexp.MustCompile(`[\p{Ll}\p{Lu}\p{Lt}\p{Lo}\p{Nd}\p{Pd}]+|[^\p{L}\p{N}\s]+`)
	words := re.FindAllString(input, -1)

	// 将所有单词转换为小写
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	return words
}
