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
