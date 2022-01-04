package main

func binarySearch(arr []int, l int, r int, x int) int {
	if r >= l {
		mid := l + (r-1)/2

		if arr[mid] == x {
			return mid
		} else if arr[mid] > x {
			return binarySearch(arr, l, mid-1, x)
		} else {
			return binarySearch(arr, mid+1, r, x)
		}
	} else {
		return -1
	}
}
