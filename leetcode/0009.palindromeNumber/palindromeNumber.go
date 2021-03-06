package palindromeNumber

import "strconv"

// 9. 回文数
// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
//
//
// 示例 1：
// 输入：x = 121
// 输出：true
// 示例2：
// 输入：x = -121
// 输出：false
// 解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
// 示例 3：
// 输入：x = 10
// 输出：false
// 解释：从右向左读, 为 01 。因此它不是一个回文数。
// 示例 4：
// 输入：x = -101
// 输出：false
//
//
// 提示：
// -2^31<= x <= 2^31- 1
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/palindrome-number
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// isPalindrome
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	sl := make([]byte, 0)
	for _, ch := range []byte(s) {
		sl = append(sl, ch)
	}
	n := len(sl)
	for i := 0; i <= (n-1)>>1; i++ {
		sl[i], sl[n-1-i] = sl[n-1-i], sl[i]
	}
	var ss string
	for _, ch := range sl {
		ss += string(ch)
	}
	if ss == s {
		return true
	}
	return false
}

// isPalindromeNum
func isPalindromeNum(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	num := 0
	for x > num {
		num = num*10 + x%10
		x /= 10
	}
	return num == x || x == num/10
}
