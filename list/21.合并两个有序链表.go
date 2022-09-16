/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	newList := &ListNode{}
	result := newList
	l1, l2 := list1, list2
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			newList.Next = l1
			l1 = l1.Next

		} else {
			newList.Next = l2
			l2 = l2.Next
		}
		newList = newList.Next
	}

	if l1 != nil {
		newList.Next = l1
	}
	if l2 != nil {
		newList.Next = l2
	}
	return result.Next
}

// @lc code=end
