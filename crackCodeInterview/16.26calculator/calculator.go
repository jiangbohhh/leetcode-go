package calculator

// 16.26 计算器
// 给定一个包含正整数、加(+)、减(-)、乘(*)、除(/)的算数表达式(括号除外)，计算其结果。
// 表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格 。 整数除法仅保留整数部分。
// 示例1:
// 输入: "3+2*2"
// 输出: 7
// 示例 2:
// 输入: " 3/2 "
// 输出: 1
// 示例 3:
// 输入: " 3+5 / 2 "
// 输出: 5
// 说明：
//
// 你可以假设所给定的表达式都是有效的。
// 请不要使用内置的库函数 eval。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/calculator-lcci
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func calculate(s string) int {
	var dataStack, operStack = make([]int, 0), make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		// 数字
		if s[i] >= '0' && s[i] <= '9' {
			val := int(s[i] - '0')
			for j := i + 1; j < len(s) && s[j] >= '0' && s[j] <= '9'; j++ {
				val = val*10 + int(s[j]-'0')
				i = j
			}
			// 出现 * / 优先计算
			if len(operStack) != 0 && len(dataStack) != 0 {
				if operStack[len(operStack)-1] == '*' {
					val = dataStack[len(dataStack)-1] * val
					dataStack = dataStack[:len(dataStack)-1]
					operStack = operStack[:len(operStack)-1]
				} else if operStack[len(operStack)-1] == '/' {
					val = dataStack[len(dataStack)-1] / val
					dataStack = dataStack[:len(dataStack)-1]
					operStack = operStack[:len(operStack)-1]
				}
			}
			dataStack = append(dataStack, val)
		}

		// 符合加入操作栈
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			operStack = append(operStack, s[i])
		}
	}
	if len(dataStack) > 0 {
		val := dataStack[0]
		dataStack = dataStack[1:]
		for len(dataStack) != 0 && len(operStack) != 0 {
			if operStack[0] == '+' {
				val += dataStack[0]
			} else if operStack[0] == '-' {
				val -= dataStack[0]
			}
			dataStack = dataStack[1:]
			operStack = operStack[1:]
		}
		return val
	}
	return 0
}
