package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a1 := listNodeToArr(l1)
	a2 := listNodeToArr(l2)

	a1Len := len(a1)
	a2Len := len(a2)
	max := a1Len
	if a2Len > a1Len {
		max = a2Len
	}

	arr := make([]int, max)
	for i := 0; i < max; i++ {
		if i < a1Len {
			arr[i] += a1[i]
		}

		if i < a2Len {
			arr[i] += a2[i]
		}

		if arr[i] >= 10 {
			arr[i+1]++
			arr[i] -= 10
		}
	}

	return arrToListNode(arr)
}

func addTwoNumbersBeta(l1 *ListNode, l2 *ListNode) *ListNode {
	t1 := l1
	t2 := l2

	result := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp := result

	for t1 != nil && t2 != nil {

		if t1 != nil {
			tmp.Val += t1.Val
			t1 = t1.Next
		}

		if t2 != nil {
			tmp.Val += t2.Val
			t2 = t2.Next
		}

		nextVal := 0
		if tmp.Val >= 10 {
			nextVal++
			tmp.Val -= 10
		}

		tmp.Next = &ListNode{
			Val:  nextVal,
			Next: nil,
		}
		tmp = tmp.Next
	}

	return result
}
