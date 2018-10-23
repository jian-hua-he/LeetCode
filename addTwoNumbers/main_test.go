package main

import "testing"

func TestArrToListNode(t *testing.T) {
	nums1 := []int{1, 2, 3}

	l1 := arrToListNode(nums1)
	t1 := l1

	if t1.Val != 1 {
		t.Errorf("Data was incorrect, got %v, want %v", t1.Val, 1)
	}

	t1 = t1.Next
	if t1.Val != 2 {
		t.Errorf("Data was incorrect, got %v, want %v", t1.Val, 2)
	}

	t1 = t1.Next
	if t1.Val != 3 {
		t.Errorf("Data was incorrect, got %v, want %v", t1.Val, 3)
	}

	if t1.Next != nil {
		t.Errorf("Data was incorrect, got %v, want %v", t1.Next, nil)
	}
}

func TestListNodeToArr(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	a1 := listNodeToArr(l1)
	if len(a1) != 3 {
		t.Errorf("Data was incorrect, got %v, want %v", len(a1), 3)
	}
	if a1[0] != 2 {
		t.Errorf("Data was incorrect, got %v, want %v", a1[0], 2)
	}
	if a1[1] != 4 {
		t.Errorf("Data was incorrect, got %v, want %v", a1[1], 4)
	}
	if a1[2] != 3 {
		t.Errorf("Data was incorrect, got %v, want %v", a1[0], 3)
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	a2 := listNodeToArr(l2)
	if len(a2) != 3 {
		t.Errorf("Data was incorrect, got %v, want %v", len(a2), 3)
	}
	if a2[0] != 5 {
		t.Errorf("Data was incorrect, got %v, want %v", a2[0], 5)
	}
	if a2[1] != 6 {
		t.Errorf("Data was incorrect, got %v, want %v", a2[1], 6)
	}
	if a2[2] != 4 {
		t.Errorf("Data was incorrect, got %v, want %v", a2[0], 4)
	}
}

func TestAddTwoNumbersCase1(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	r1 := addTwoNumbers(l1, l2)
	if r1 == nil {
		t.Errorf("Data was incorrect, got %v, want reference of ListNode", r1)
		return
	}
	if r1.Val != 7 {
		t.Errorf("Data was incorrect, got %v, want %v", r1.Val, 7)
	}

	if r1.Next == nil {
		t.Errorf("Data was incorrect, got %v, want reference of ListNode", r1.Next)
		return
	}
	r1 = r1.Next
	if r1.Val != 0 {
		t.Errorf("Data was incorrect, got %v, want %v", r1.Val, 0)
	}

	if r1.Next == nil {
		t.Errorf("Data was incorrect, got %v, want reference of ListNode", r1.Next)
		return
	}
	r1 = r1.Next
	if r1.Val != 8 {
		t.Errorf("Data was incorrect, got %v, want %v", r1.Val, 8)
	}

	if r1.Next != nil {
		t.Errorf("Data was incorrect, got %v, want %v", r1.Next, nil)
	}
}

func TestAddTwoNumbersCase2(t *testing.T) {
	// Prepare data
	t1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	l1 := arrToListNode(t1)

	t2 := []int{5, 6, 4}
	l2 := arrToListNode(t2)

	// Transfer to slice
	r1 := addTwoNumbers(l1, l2)
	if r1 == nil {
		t.Errorf("Data was incorrect, got %v, want reference of ListNode", r1)
		return
	}
	if r1.Val != 6 {
		t.Errorf("Data was incorrect, got %v, want %v", r1.Val, 6)
	}

	r1 = r1.Next
	if r1 == nil {
		t.Errorf("Data was incorrect, got %v, want reference of ListNode", r1)
		return
	}
	if r1.Val != 6 {
		t.Errorf("Data was incorrect, got %v, want %v", r1.Val, 6)
	}

	r1 = r1.Next
	if r1 == nil {
		t.Errorf("Data was incorrect, got %v, want reference of ListNode", r1)
		return
	}
	if r1.Val != 4 {
		t.Errorf("Data was incorrect, got %v, want %v", r1.Val, 4)
	}

	for i := 1; i <= 27; i++ {
		r1 = r1.Next
		if r1 == nil {
			t.Errorf("Data was incorrect, got %v, want reference of ListNode", r1)
			return
		}
		if r1.Val != 0 {
			t.Errorf("Data was incorrect, got %v, want %v", r1.Val, 0)
		}
	}

	r1 = r1.Next
	if r1 == nil {
		t.Errorf("Data was incorrect, got %v, want reference of ListNode", r1)
		return
	}
	if r1.Val != 1 {
		t.Errorf("Data was incorrect, got %v, want %v", r1.Val, 1)
	}
}
