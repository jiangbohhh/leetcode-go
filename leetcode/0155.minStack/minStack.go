package minStack

// 0155. min Stack
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
// push(x) —— 将元素 x 推入栈中。
// pop()—— 删除栈顶的元素。
// top()—— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。
//
// 示例:
// 输入：
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
// 输出：
// [null,null,null,null,-3,null,0,-2]
//
// 解释：
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.getMin();   --> 返回 -2.
//
// 提示：
// pop、top 和 getMin 操作总是在 非空栈 上调用。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/min-stack
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type MinStack struct {
	Stack    []int
	MinStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Stack:    make([]int, 0),
		MinStack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	if len(this.MinStack) > 0 && this.MinStack[len(this.MinStack)-1] < val {
		this.MinStack = append(this.MinStack, this.MinStack[len(this.MinStack)-1])
	} else {
		this.MinStack = append(this.MinStack, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.Stack) > 0 && len(this.MinStack) > 0 {
		this.Stack = this.Stack[:len(this.Stack)-1]
		this.MinStack = this.MinStack[:len(this.MinStack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.Stack) > 0 {
		return this.Stack[len(this.Stack)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.MinStack) > 0 {
		return this.MinStack[len(this.MinStack)-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
