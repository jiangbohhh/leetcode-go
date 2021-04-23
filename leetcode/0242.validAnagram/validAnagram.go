package validAnagram

import (
	"sort"
)

// 0242. Valid Anagram
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 示例1:
// 输入: s = "anagram", t = "nagaram"
// 输出: true
// 示例 2:
// 输入: s = "rat", t = "car"
// 输出: false
// 说明:
// 你可以假设字符串只包含小写字母。
// 进阶:
// 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/valid-anagram
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ls, lt := make([]string, 0), make([]string, 0)
	for _, ch := range s {
		ls = append(ls, string(ch))
	}
	for _, ch := range t {
		lt = append(lt, string(ch))
	}
	sort.Strings(ls)
	sort.Strings(lt)
	for i := 0; i < len(ls); i++ {
		if ls[i] != lt[i] {
			return false
		}
	}
	return true
}
