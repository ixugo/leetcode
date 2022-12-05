/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */
package leetcode

// @lc code=start
type MyQueue struct {
	stack      []int
	begin, end int
}

func Constructor() MyQueue {
	return MyQueue{
		stack: make([]int, 0, 1),
	}
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
	this.end++
}

func (this *MyQueue) Pop() int {
	x := this.stack[this.begin]
	this.begin++
	return x
}

func (this *MyQueue) Peek() int {
	return this.stack[this.begin]
}

func (this *MyQueue) Empty() bool {
	return this.begin == this.end
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
