package main

func findDisappearedNumbers(nums []int) []int {
	var ans []int

	for _, num := range nums {
		nums[abs(num)-1] = -abs(nums[abs(num)-1])
	}

	for i, num := range nums {
		if num > 0 {
			ans = append(ans, i+1)
		}
	}

	return ans
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
