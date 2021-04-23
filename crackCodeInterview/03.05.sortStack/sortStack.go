package sortStack

// 03.05. Sort Stack
// 栈排序。 编写程序，对栈进行排序使最小元素位于栈顶。最多只能使用一个其他的临时栈存放数据，但不得将元素复制到别的数据结构（如数组）中。
// 该栈支持如下操作：push、pop、peek 和 isEmpty。当栈为空时，peek返回 -1。
// 示例1:
// 输入：
// ["SortedStack", "push", "push", "peek", "pop", "peek"]
// [[], [1], [2], [], [], []]
// 输出：
// [null,null,null,1,null,2]
// 示例2:
// 输入：
// ["SortedStack", "pop", "pop", "push", "pop", "isEmpty"]
// [[], [], [], [1], [], []]
// 输出：
// [null,null,null,null,null,true]
// 说明:
// 栈中的元素数目在[0, 5000]范围内。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sort-of-stacks-lcci
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type SortedStack struct {
	dataStack   []int
	assistStack []int
}

func Constructor() SortedStack {
	return SortedStack{
		dataStack:   make([]int, 0),
		assistStack: make([]int, 0),
	}
}

func (this *SortedStack) Push(val int) {

	if len(this.dataStack) == 0 {
		this.dataStack = append(this.dataStack, val)
	} else {
		// 把 小于 val 的值挪到辅助栈中
		for len(this.dataStack) > 0 && this.dataStack[len(this.dataStack)-1] < val {
			this.assistStack = append(this.assistStack, this.dataStack[len(this.dataStack)-1])
			this.dataStack = this.dataStack[:len(this.dataStack)-1]
		}
		this.dataStack = append(this.dataStack, val)
		// 将 辅助栈中值依次弹出，压入数据栈
		for len(this.assistStack) != 0 {
			this.dataStack = append(this.dataStack, this.assistStack[len(this.assistStack)-1])
			this.assistStack = this.assistStack[:len(this.assistStack)-1]
		}
	}
}

func (this *SortedStack) Pop() {
	if len(this.dataStack) >= 1 {
		this.dataStack = this.dataStack[:len(this.dataStack)-1]
	}
}

func (this *SortedStack) Peek() int {
	if len(this.dataStack) > 0 {
		return this.dataStack[len(this.dataStack)-1]
	}
	return -1
}

func (this *SortedStack) IsEmpty() bool {
	if len(this.dataStack) == 0 {
		return true
	}
	return false
}

/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */
