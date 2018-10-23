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
