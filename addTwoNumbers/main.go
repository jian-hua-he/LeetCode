package main

import (
	"strconv"
)

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

func numToListNode(num int) (result *ListNode, err error) {
	str := strconv.Itoa(num)
	result = new(ListNode)

	curr := result
	for i := len(str) - 1; i >= 0; i-- {
		n, err := strconv.Atoi(string(str[i]))

		if err != nil {
			return nil, err
		}

		curr.Val = n
		if i != 0 {
			curr.Next = new(ListNode)
			curr = curr.Next
		}
	}

	return result, nil
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
	n1, _ := toNum(l1)
	n2, _ := toNum(l2)
	result, _ := numToListNode(n1 + n2)

	return result
}
