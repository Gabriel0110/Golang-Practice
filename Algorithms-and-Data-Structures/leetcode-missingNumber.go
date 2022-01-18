import "sort"

func missingNumber(nums []int) int {
	sort.Ints(nums)
	i := 0
	missing := -1
	for _, x := range nums {
		if x == i {
			i++
		} else {
			missing = i
			break
		}
	}

	if missing == -1 {
		missing = nums[len(nums)-1] + 1
	}

	return missing
}