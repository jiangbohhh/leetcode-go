package oneAway

import (
	"math"
)

// 面试题 01.05. 一次编辑
// 字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
//
// 示例1:
// 输入:
// first = "pale"
// second = "ple"
// 输出: True
//
// 示例2:
// 输入:
// first = "pales"
// second = "pal"
// 输出: False
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/one-away-lcci
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func oneEditAway(first string, second string) bool {
	// 长度大于 1 则 false
	if math.Abs(float64(len(first)-len(second))) > 1 {
		return false
	}
	// 长度相等执行替换指令，长度相差 1 执行插入或者删除指令（删除可以认为是反向插入）

	short := first
	long := second
	if len(first) > len(second) {
		short, long = second, first
	} else {
		short, long = first, second
	}
	// i 代表短串移动的指针，j 代表长串移动的指针， flag 标识第几次不相同
	i, j := 0, 0
	flag := false
	for i < len(short) && j < len(long) {
		if short[i] != long[j] {
			if flag {
				return false
			}
			flag = true
			// 如果是替换 就让两个指针都加 1
			if len(short) == len(long) {
				i++
			}
		} else {
			i++
		}
		// 如果不相等 只加长的串的指针
		j++
	}
	return true
}
