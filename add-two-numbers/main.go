package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 8,
		},
	}

	l2 := &ListNode{
		Val: 0,
		// Next: &ListNode{
		// 	Val: 8,
		// },
	}

	l3 := addTwoNumbers(l1, l2)
	for {
		if l3.Next == nil {
			fmt.Printf("%d\n", l3.Val)
			break
		}
		fmt.Printf("%d", l3.Val)
		l3 = l3.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0                // 进位标示
	head := &ListNode{Val: 0} // 头节点
	cur := head

	for l1 != nil || l2 != nil {
		// 读取链表
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		// 进位
		carry = sum / 10
		// 取余
		sum = sum % 10
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
	}

	if carry == 1 {
		cur.Next = &ListNode{Val: carry}
	}
	return head.Next
}
