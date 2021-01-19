package main

import "fmt"

func main()  {
	l1 := newList([]int{8, 9, 9}) // 5	// 9 8
	l2 := newList([]int{2})       // 5	// 1
	cur := addTwoNumbers(l1, l2)
	for cur != nil {
		fmt.Print(cur.String())
		cur = cur.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func newList(nums []int) *ListNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0], Next: nil}
	cur := head
	for i := 1; i < n; i++ {
		newNode := &ListNode{Val: nums[i], Next: nil}
		cur.Next = newNode
		cur = newNode
	}
	return head
}

func (cur *ListNode) String() string {
	counts := 0
	var nums []int
	for cur != nil {
		nums = append(nums, cur.Val)
		counts++
		cur = cur.Next
	}
	return fmt.Sprintf("%d nodes: %v", counts, nums)
}

// 遍历两个链表，取出各自的数字再相加（难以解决整数溢出的问题，1560 / 1563 个通过测试用例，取值求解不可行）
// 遍历链表，相互相加记进位。注意特9殊情况的处理
func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail = &ListNode{Val: 0}
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}