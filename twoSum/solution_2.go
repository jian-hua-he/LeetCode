package main

import "sort"

type Num struct {
	Index int
	Data  int
}

type ByNum []Num

func (nums ByNum) Len() int {
	return len(nums)
}

func (nums ByNum) Swap(i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func (nums ByNum) Less(i, j int) bool {
	return nums[i].Data < nums[j].Data
}

func twoSum2(nums []int, target int) []int {
	arr := make([]Num, 0)
	for i, v := range nums {
		arr = append(arr, Num{
			Index: i,
			Data:  v,
		})
	}

	sort.Sort(ByNum(arr))

	i := 0
	j := len(arr) - 1

	for i <= j {
		n := arr[i].Data + arr[j].Data

		if n == target {
			break
		}

		if n > target {
			j -= 1
			continue
		}

		if n < target {
			i += 1
			continue
		}
	}

	return []int{arr[j].Index, arr[i].Index}
}
