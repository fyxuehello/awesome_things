package tools

// 统计中文字符串的长度 e.g:你好啊，长度为3
func CountCharacters(str string) int {
	l := []rune(str)
	return len(l)
}
