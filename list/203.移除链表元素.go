/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package list

// 注意可能出现连续要删除的情况
// 所以只有当  list.Val != val 时，才能开始下一个结点
func removeElements(head *ListNode, val int) *ListNode {
	l := &ListNode{Next: head}

	t := l
	for t.Next != nil {
		if t.Next.Val == val {
			t.Next = t.Next.Next
		} else {
			t = t.Next
		}
	}
	return l.Next
}

// @lc code=end
