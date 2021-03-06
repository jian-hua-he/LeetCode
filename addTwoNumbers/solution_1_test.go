package main

import "testing"

// Real Test
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

	r := addTwoNumbers(l1, l2)
	if r == nil {
		t.Errorf("Data was incorrect, got %v, want reference of ListNode", r)
		return
	}
	if r.Val != 7 {
		t.Errorf("Data was incorrect, got %v, want %v", r.Val, 7)
	}

	if r.Next == nil {
		t.Errorf("Data was incorrect, got %v, want reference of ListNode", r.Next)
		return
	}
	r = r.Next
	if r.Val != 0 {
		t.Errorf("Data was incorrect, got %v, want %v", r.Val, 0)
	}

	if r.Next == nil {
		t.Errorf("Data was incorrect, got %v, want reference of ListNode", r.Next)
		return
	}
	r = r.Next
	if r.Val != 8 {
		t.Errorf("Data was incorrect, got %v, want %v", r.Val, 8)
	}

	if r.Next != nil {
		t.Errorf("Data was incorrect, got %v, want %v", r.Next, nil)
	}
}

func TestAddTwoNumbersCase2(t *testing.T) {
	// Prepare data
	t1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	l1 := arrToListNode(t1)

	t2 := []int{5, 6, 4}
	l2 := arrToListNode(t2)

	// Transfer to slice
	r := addTwoNumbers(l1, l2)
	if r == nil {
		t.Errorf("Data was incorrect, got %v, want reference of ListNode", r)
		return
	}
	if r.Val != 6 {
		t.Errorf("Data was incorrect, got %v, want %v", r.Val, 6)
	}

	r = r.Next
	if r == nil {
		t.Errorf("Data was incorrect, got %v, want reference of ListNode", r)
		return
	}
	if r.Val != 6 {
		t.Errorf("Data was incorrect, got %v, want %v", r.Val, 6)
	}

	r = r.Next
	if r == nil {
		t.Errorf("Data was incorrect, got %v, want reference of ListNode", r)
		return
	}
	if r.Val != 4 {
		t.Errorf("Data was incorrect, got %v, want %v", r.Val, 4)
	}

	for i := 0; i < 27; i++ {
		r = r.Next
		if r == nil {
			t.Logf("Index is %v, List Node is %v", i, r)
			t.Errorf("Data was incorrect, got %v, want reference of ListNode", r)
			return
		}
		if r.Val != 0 {
			t.Logf("Index is %v, List Node is %v", i, r)
			t.Errorf("Data was incorrect, got %v, want %v", r.Val, 0)
		}
	}

	r = r.Next
	if r == nil {
		t.Errorf("Data was incorrect, got %v, want reference of ListNode", r)
		return
	}
	if r.Val != 1 {
		t.Errorf("Data was incorrect, got %v, want %v", r.Val, 1)
	}
}

func TestAddTwoNumbersCase3(t *testing.T) {
	l1 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  5,
		Next: nil,
	}

	r := addTwoNumbers(l1, l2)
	if r == nil {
		t.Errorf("Data was incorrect, got %v, want reference of ListNode", r)
		return
	}
	if r.Val != 0 {
		t.Errorf("Data was incorrect, got %v, want %v", r.Val, 0)
	}

	r = r.Next
	if r == nil {
		t.Errorf("Data was incorrect, got %v, want reference of ListNode", r)
		return
	}
	if r.Val != 1 {
		t.Errorf("Data was incorrect, got %v, want %v", r.Val, 1)
	}
	if r.Next != nil {
		t.Errorf("Data was incorrect, got %v, want %v", r.Next, nil)
	}
}

// Test Helper
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
