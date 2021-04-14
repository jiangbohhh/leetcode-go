package ticTacWin

// 面试题 16.04. 井字游戏
// 设计一个算法，判断玩家是否赢了井字游戏。输入是一个 N x N 的数组棋盘，由字符" "，"X"和"O"组成，其中字符" "代表一个空位。
//
// 以下是井字游戏的规则：
// 玩家轮流将字符放入空位（" "）中。
// 第一个玩家总是放字符"O"，且第二个玩家总是放字符"X"。
// "X"和"O"只允许放置在空位中，不允许对已放有字符的位置进行填充。
// 当有N个相同（且非空）的字符填充任何行、列或对角线时，游戏结束，对应该字符的玩家获胜。
// 当所有位置非空时，也算为游戏结束。
// 如果游戏结束，玩家不允许再放置字符。
// 如果游戏存在获胜者，就返回该游戏的获胜者使用的字符（"X"或"O"）；如果游戏以平局结束，则返回 "Draw"；如果仍会有行动（游戏未结束），则返回 "Pending"。
//
// 示例 1：
// 输入： board = ["O X"," XO","X O"]
// 输出： "X"
// 示例 2：
// 输入： board = ["OOX","XXO","OXO"]
// 输出： "Draw"
// 解释： 没有玩家获胜且不存在空位
// 示例 3：
// 输入： board = ["OOX","XXO","OX "]
// 输出： "Pending"
// 解释： 没有玩家获胜且仍存在空位
//
// 提示：
// 1 <= board.length == board[i].length <= 100
// 输入一定遵循井字棋规则
//

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/tic-tac-toe-lcci
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// tictactoe
func tictactoe(board []string) string {
	n := len(board)
	list := make([][]byte, n)
	for idx, str := range board {
		list[idx] = make([]byte, 0, n)
		for i := 0; i < len(str); i++ {
			list[idx] = append(list[idx], str[i])
		}
	}
	isWon := false
	// 行上判断是否相等
	for i := 0; i < n; i++ {
		if list[i][0] == ' ' {
			continue
		}
		isWon = true
		for j := 0; j < n; j++ {
			if list[i][j] != list[i][0] {
				isWon = false
				break
			}
		}
		if isWon {
			return string(list[i][0])
		}
	}

	// 列上判断是否相等
	for i := 0; i < n; i++ {
		if list[0][i] == ' ' {
			continue
		}
		isWon = true
		for j := 0; j < n; j++ {
			if list[j][i] != list[0][i] {
				isWon = false
				break
			}
		}
		if isWon {
			return string(list[0][i])
		}
	}

	// 对角判断是否相等(左上到右下）
	if list[0][0] != ' ' {
		i, j := 1, 1
		isWon = true
		for i < n && j < n {
			if list[i][j] != list[0][0] {
				isWon = false
				break
			}
			i++
			j++
		}
		if isWon {
			return string(list[0][0])
		}
	}

	// 对角线(左下到右上)
	if list[n-1][0] != ' ' {
		i, j := n-2, 1
		isWon = true
		for i >= 0 && j < n {
			if list[i][j] != list[n-1][0] {
				isWon = false
				break
			}
			i--
			j++
		}
		if isWon {
			return string(list[n-1][0])
		}
	}
	// 如果没人赢，有空则代表Pending
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if list[i][j] == ' ' {
				return "Pending"
			}
		}
	}
	return "Draw"
}
