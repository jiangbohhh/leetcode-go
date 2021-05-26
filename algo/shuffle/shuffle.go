package shuffle

import "math/rand"

// 洗牌算法

func shuffle(list []int, n int) {
	for i := 0; i < n; i++ {
		// 拿到剩下 n -i 中的一个位置，在加上当前位置 即为交换位置
		j := rand.Int()%(n-i) + i
		list[i], list[j] = list[j], list[i]
	}
}
