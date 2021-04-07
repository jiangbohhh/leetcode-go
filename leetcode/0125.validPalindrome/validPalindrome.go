package validPalindrome

import (
	"strings"
	"unicode"
)

// 125. 验证回文串
// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
// 说明：本题中，我们将空字符串定义为有效的回文串。
//
// 示例 1:
// 输入: "A man, a plan, a canal: Panama"
// 输出: true
//
// 示例 2:
// 输入: "race a car"
// 输出: false
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/valid-palindrome
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// isPalindrome
func isPalindrome(s string) bool {
	// 剔除非数字和字母的字符
	s = trimPunctAndSpace(s)
	if s == "" {
		return true
	}
	// 将字母转换成小写
	s = strings.ToLower(s)

	// 将字符进行翻转，如果翻转后字符串和原字符串相等则代表是回文串
	var list = make([]byte, 0)
	for i := 0; i < len(s); i++ {
		list = append(list, s[i])
	}
	reverse(list)
	if string(list) == s {
		return true
	}
	return false
}

// reverse
func reverse(s []byte) {
	for i := 0; i <= (len(s)-1)>>1; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

// trimPunctAndSpace 剔除标点字符和空格
func trimPunctAndSpace(s string) string {
	var ss = make([]rune, 0)
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			if unicode.IsUpper(r) {
				r = unicode.ToLower(r)
			}
			ss = append(ss, r)
		}
	}
	return string(ss)
}
