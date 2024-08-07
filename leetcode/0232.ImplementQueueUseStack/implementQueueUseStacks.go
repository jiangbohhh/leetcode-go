package implementQueueUseStack

// 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
//实现 MyQueue 类：
//void push(int x) 将元素 x 推到队列的末尾
//int pop() 从队列的开头移除并返回元素
//int peek() 返回队列开头的元素
//boolean empty() 如果队列为空，返回 true ；否则，返回 false
//说明：
//你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
//你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
//示例 1：
//输入：
//["MyQueue", "push", "push", "peek", "pop", "empty"]
//[[], [1], [2], [], [], []]
//输出：
//[null, null, null, 1, 1, false]
//解释：
//MyQueue myQueue = new MyQueue();
//myQueue.push(1); // queue is: [1]
//myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
//myQueue.peek(); // return 1
//myQueue.pop(); // return 1, queue is [2]
//myQueue.empty(); // return false
//提示：
//1 <= x <= 9
//最多调用 100 次 push、pop、peek 和 empty
//假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）

// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/implement-queue-using-stacks
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type MyQueue struct {
	Stack1 Stack
	Stack2 Stack
}

func Constructor() MyQueue {
	return MyQueue{
		Stack1: Stack{},
		Stack2: Stack{},
	}
}

func (this *MyQueue) Push(x int) {
	for !this.Stack1.IsEmpty() {
		this.Stack2.Push(this.Stack1.Pop())
	}
	this.Stack1.Push(x)
	for !this.Stack2.IsEmpty() {
		this.Stack1.Push(this.Stack2.Pop())
	}
}

func (this *MyQueue) Pop() int {
	return this.Stack1.Pop()
}

func (this *MyQueue) Peek() int {
	return this.Stack1.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.Stack1.IsEmpty()
}

type Stack struct {
	stackList []int
}

func (s *Stack) Push(val int) {
	s.stackList = append(s.stackList, val)
	return
}

func (s *Stack) Pop() int {
	val := s.stackList[len(s.stackList)-1]
	s.stackList = s.stackList[:len(s.stackList)-1]
	return val
}

func (s *Stack) IsEmpty() bool {
	return len(s.stackList) == 0
}

func (s *Stack) Size() int {
	return len(s.stackList)
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.stackList[len(s.stackList)-1]
}
