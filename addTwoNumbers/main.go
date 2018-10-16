package main

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func toNum(l *ListNode) (result int, err error) {
	str := ""

	list := l
	for list != nil {
		str = strconv.Itoa(list.Val) + str
		list = list.Next
	}

	result, err = strconv.Atoi(str)

	return
}

func toListNode(num int) (result *ListNode, err error) {
	str := strconv.Itoa(num)
	result = new(ListNode)

	curr := result
	for i := len(str) - 1; i >= 0; i-- {
		n, err := strconv.Atoi(string(str[i]))

		if err != nil {
			return nil, err
		}

		curr.Val = n
		curr.Next = new(ListNode)
		curr = curr.Next
	}

	return result, nil
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1, _ := toNum(l1)
	n2, _ := toNum(l2)
	result, _ := toListNode(n1 + n2)

	return result
}
