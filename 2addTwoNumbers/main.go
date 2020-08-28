package main

import "fmt"

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 根据int创建一个链表
// 链表头是num的个位
// 后面的节点依次为十位、百位。。。
func getListNode(num int) *ListNode {
	head := new(ListNode)
	x := 10
	if num%x != num {
		head.Val = num % x
		// 再创建一个node 赋值给head.Next
		head.Next = getListNode(num / x)
		x = x * 10
	} else {
		head.Val = num
		head.Next = nil
	}
	return head
}

func toInt(node *ListNode) int {
	x := 1
	num := 0
	pointer := node
	for {
		num = num + pointer.Val*x
		x = x * 10
		if pointer.Next == nil {
			return num
		}
		pointer = pointer.Next
	}

}

func (node *ListNode) show() {
	fmt.Print(node.Val)
	if node.Next != nil {
		fmt.Print(" -> ")
		node.Next.show()
	}
}

func addTwoNumbersToInt(l1 *ListNode, l2 *ListNode) int64 {
	if l1.Next != nil || l2.Next != nil {

	}
	return 0
}

// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

// 受到int长度的限制
func addTwoNumbersV1(l1 *ListNode, l2 *ListNode) *ListNode {
	return getListNode(toInt(l1) + toInt(l2))
}

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	// 定义伪head
	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	// 声明指针
	p, q, current := l1, l2, dummyHead
	for {
		if p != nil || q != nil {
			x := 0
			if p != nil {
				x = p.Val
			}
			y := 0
			if q != nil {
				y = q.Val
			}
			// 进位
			carry := (current.Val + x + y) / 10
			if carry > 0 {
				current.Val = (current.Val + x + y) % 10
			} else {
				current.Val = (current.Val + x + y)
			}

			if p == nil && q != nil {
				p = &ListNode{
					Val:  0,
					Next: nil,
				}
			} else if p != nil {
				p = p.Next
			} else if p == nil && carry > 0 {
				p = &ListNode{
					Val:  0,
					Next: nil,
				}
			}
			if q == nil && p != nil {
				q = &ListNode{
					Val:  0,
					Next: nil,
				}
			} else if q != nil {
				q = q.Next
			} else if q == nil && carry > 0 {
				q = &ListNode{
					Val:  0,
					Next: nil,
				}
			}

			// 添加链表条件
			// 1.进位
			// 2.下一位存在（p q已经指向下一个node）
			if carry > 0 || (p != nil || q != nil) {
				next := &ListNode{
					// 进位
					Val:  carry,
					Next: nil,
				}
				current.Next = next
				current = next
			}

		} else {
			break
		}
	}
	return dummyHead
}

func main() {
	// 0
	l1 := getListNode(0)
	// 7 -> 3
	l2 := getListNode(37)
	// l3 := addTwoNumbersV1(l1, l2)
	// l3.show()
	l3 := addTwoNumbersV2(l1, l2)
	l3.show()
}
