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

// longestPalindrome1 使用中心拓展法，以每个字符为中心(奇数长度回文)或每两个相邻字符为中心(偶数长度回文),向两边扩展。
func longestPalindrome1(s string) string {
	if len(s) < 2 {
		return s
	}
	start, maxLen := 0, 1
	for i := 0; i < len(s); i++ {
		len1 := expendAroundCenter(s, i, i)
		len2 := expendAroundCenter(s, i, i+1)
		length := max(len1, len2)
		if length > maxLen {
			start = i - (length-1)/2
			maxLen = length
		}
	}
	return s[start : start+maxLen]
}

func expendAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// longestPalindrome2 使用动态规划方式
func longestPalindrome2(s string) string {
	n := len(s)
	if len(s) < 2 {
		return s
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	start, maxLen := 0, 1

	// 所有长度为1的字串都是回文串
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// 检查长度为2的回文串
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLen = 2
		}
	}

	// 检查长度大于3的子串
	for length := 3; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if length > maxLen {
					start = i
					maxLen = length
				}
			}
		}
	}

	return s[start : start+maxLen]
}
