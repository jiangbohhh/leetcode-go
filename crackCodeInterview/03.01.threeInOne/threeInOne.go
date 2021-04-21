package threeInOne

// 03.01 Three In One
// 三合一。描述如何只用一个数组来实现三个栈。
// 你应该实现push(stackNum, value)、pop(stackNum)、isEmpty(stackNum)、peek(stackNum)方法。stackNum表示栈下标，value表示压入的值。
// 构造函数会传入一个stackSize参数，代表每个栈的大小。
//
// 示例1:
// 输入：
// ["TripleInOne", "push", "push", "pop", "pop", "pop", "isEmpty"]
// [[1], [0, 1], [0, 2], [0], [0], [0], [0]]
// 输出：
// [null, null, null, 1, -1, -1, true]
// 说明：当栈为空时`pop, peek`返回-1，当栈满时`push`不压入元素。
// 示例2:
// 输入：
// ["TripleInOne", "push", "push", "push", "pop", "pop", "pop", "peek"]
// [[2], [0, 1], [0, 2], [0, 3], [0], [0], [0], [0]]
// 输出：
// [null, null, null, null, 2, 1, -1, -1]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/three-in-one-lcci
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TripleInOne struct {
	stack    []int
	size     int
	sizeList [3]int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		stack:    make([]int, stackSize*3),
		size:     stackSize,
		sizeList: [3]int{},
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	if this.sizeList[stackNum] < this.size {
		this.stack[this.sizeList[stackNum]+this.size*stackNum] = value
		this.sizeList[stackNum]++
	}
}

func (this *TripleInOne) Pop(stackNum int) int {
	if this.sizeList[stackNum] == 0 {
		return -1
	}
	val := this.stack[this.sizeList[stackNum]+this.size*stackNum-1]
	this.sizeList[stackNum]--
	return val
}

func (this *TripleInOne) Peek(stackNum int) int {
	if this.sizeList[stackNum] == 0 {
		return -1
	}
	return this.stack[this.sizeList[stackNum]+this.size*stackNum-1]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.sizeList[stackNum] == 0
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */
