package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	t1 := l1
	t2 := l2

	result := &ListNode{
		Val:  0,
		Next: nil,
	}

	for curr := result; t1 != nil || t2 != nil; curr = curr.Next {

		// Calculate value and go to next for input linked list
		if t1 != nil {
			curr.Val += t1.Val
			t1 = t1.Next
		}

		if t2 != nil {
			curr.Val += t2.Val
			t2 = t2.Next
		}

		// Handle next
		nextVal := 0
		if curr.Val >= 10 {
			nextVal += 1
			curr.Val -= 10
			curr.Next = &ListNode{
				Val:  nextVal,
				Next: nil,
			}
			continue
		}

		if t1 == nil && t2 == nil {
			curr.Next = nil
			continue
		}

		curr.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
	}

	return result
}

// Helpers
func arrToListNode(nums []int) *ListNode {
	result := new(ListNode)

	temp := result
	for i, n := range nums {
		temp.Val = n

		if i < len(nums)-1 {
			temp.Next = new(ListNode)
			temp = temp.Next
		}
	}

	return result
}

func listNodeToArr(l *ListNode) []int {
	result := []int{}

	temp := l
	for temp != nil {
		result = append(result, temp.Val)
		temp = temp.Next
	}

	return result
}
