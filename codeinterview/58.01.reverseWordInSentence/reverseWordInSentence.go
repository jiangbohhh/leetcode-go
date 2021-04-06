package reverseWordInSentence

// 剑指 Offer 58.I. 翻转单词顺序
//
// 输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。
// 例如输入字符串"I am a student. "，则输出"student. a am I"。
//
// 示例 1：
// 输入: "the sky is blue"
// 输出:"blue is sky the"
//
// 示例 2：
// 输入: " hello world! "
// 输出:"world! hello"
// 解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//
// 示例 3：
// 输入: "a good  example"
// 输出:"example good a"
// 解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
//
// 说明：
// 无空格字符构成一个单词。
// 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
// 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// reverseWords
func reverseWords(s string) string {
	// 去除多于的空格
	s = trimSpace(s)
	if s == "" {
		return ""
	}

	// 翻转整个句子
	sList := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		sList = append(sList, s[i])
	}
	sList = reverse(sList)

	// 翻转每个单词
	var start, end int
	for start < len(sList) {
		if sList[start] == ' ' {
			start++
			end++
		} else if end == len(sList) || sList[end] == ' ' {
			reverse(sList[start:end])
			start = end
		} else {
			end++
		}
	}
	return string(sList)
}

// reverse
func reverse(s []byte) []byte {
	for i := 0; i <= (len(s)-1)>>1; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return s
}

// trimSpace 去除空格
func trimSpace(s string) string {
	var data string
	var start, end int
	for start < len(s) {
		if s[start] == ' ' {
			start++
			end++
		} else if end == len(s) || s[end] == ' ' {
			// 保证单词之间保留空格
			if data != "" {
				data += " "
			}
			data += s[start:end]
			start = end
		} else {
			end++
		}
	}
	return data
}
