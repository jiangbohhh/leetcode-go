package sort

import (
	"math/rand"
	"time"
)

// RandomSequence
func RandomSequence(num int) (sequence []int) {
	sequence = make([]int, num, num)
	rand.Seed(time.Now().UnixNano())
	var i int
	for i = 0; i < num; i++ {
		sequence[i] = rand.Intn(999) - rand.Intn(999)
	}
	return
}
