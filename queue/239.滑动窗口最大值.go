/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */
package leetcode

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	q := NewDeque(k)
	res := make([]int, 0, 1)
	for i, v := range nums {
		// 滑动窗口右移时，删除左边第一个
		if i >= k && q.GetFront() <= i-k {
			q.DeleteFront()
		}
		// 保持左边的元素最大，删除所有比 v 小的元素
		for !q.IsEmpty() && nums[q.GetRear()] <= v {
			q.DeleteLast()
		}
		q.InsertLast(i)
		// 窗口中填满值
		if i >= k-1 {
			res = append(res, nums[q.GetFront()])
		}
	}
	return res
}

// @lc code=end

type Deque struct {
	d       []int
	i, j, k int
}

func NewDeque(k int) Deque {
	if k < 1 {
		k = 1
	}
	k++
	return Deque{
		d: make([]int, k),
		k: k,
	}
}

func (this *Deque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.i = (this.i - 1 + this.k) % this.k
	this.d[this.i] = value
	return true
}

func (this *Deque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.d[this.j] = value
	this.j = (this.j + 1) % this.k
	return true
}

func (this *Deque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.i = (this.i + 1) % this.k
	return true
}

func (this *Deque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.j = (this.j - 1 + this.k) % this.k
	return true
}

func (this *Deque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.d[this.i]
}

func (this *Deque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.d[(this.j-1+this.k)%this.k]
}

func (this *Deque) IsEmpty() bool {
	return this.i == this.j
}

func (this *Deque) IsFull() bool {
	return (this.j+1)%this.k == this.i
}
