package lengthOfLongestSubstring

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。

//示例 1:
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
// 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//提示：

//0 <= s.length <= 5 * 104
//s 由英文字母、数字、符号和空格组成

// lengthOfLongestSubstring 计算最长子串长度
func lengthOfLongestSubstring(s string) int {
	// 创建一个哈希表用于记录字符是否出现过以及其出现的位置
	charMap := make(map[byte]int)
	// 初始化滑动窗口的左右边界和最长子串长度
	left, right, maxLength := 0, 0, 0

	// 遍历字符串
	for right < len(s) {
		// 获取当前字符
		currChar := s[right]
		// 如果当前字符在哈希表中出现过，并且其位置在滑动窗口的左边界之后
		if index, ok := charMap[currChar]; ok && index >= left {
			// 将左边界移动到重复字符的位置的下一个位置
			left = index + 1
		}
		// 更新哈希表中当前字符的位置
		charMap[currChar] = right
		// 更新最长子串长度
		maxLength = Max(maxLength, right-left+1)
		// 右边界向右移动
		right++
	}

	return maxLength
}

// Max 辅助函数，返回两个整数中的最大值
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
