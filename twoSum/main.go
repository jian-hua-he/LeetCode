package main

type numArr []int

func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		n := target - num
		if j := numArr(nums).indexOf(n); j >= 0 && i != j {
			return []int{i, j}
		}
	}

	return []int{}
}

func (nums numArr) indexOf(num int) int {
	for i, n := range nums {
		if n == num {
			return i
		}
	}

	return -1
}
