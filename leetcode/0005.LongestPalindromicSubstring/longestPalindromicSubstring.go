package _005_LongestPalindromicSubstring

// 给你一个字符串 s，找到 s 中最长的回文子串。
//
// 示例 1：
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 示例 2：
// 输入：s = "cbbd"
// 输出："bb"
//
// 提示：
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-palindromic-substring
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	starti, i := 0, 1
	list := make([]string, 0)
	for {
		if s[starti] == s[i] {
			list = append(list, s[starti:i+1])
			starti++
			i = starti + 1
			continue
		}
		i++
		if i == len(s) {
			starti++
			i = starti + 1
		}
		if starti == len(s) || i >= len(s) {
			break
		}

	}
	if len(list) == 0 {
		return ""
	}
	long := list[0]
	for _, str := range list {
		if len(str) > len(long) {
			long = str
		}
	}
	return long
}
