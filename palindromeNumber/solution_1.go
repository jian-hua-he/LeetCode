package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	arr := make([]int, 0)

	for n := x; n > 0; n /= 10 {
		remainer := n % 10
		arr = append(arr, remainer)
	}

	for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}

	return true
}
