package main

func binary_search(s []int, item int) int {
	down := 0
	end := len(s)
	for down <= end {
		mid := (down + end) / 2
		value := mid
		if value == item {
			return mid
		}
		if value > item {
			end = mid - 1

		} else {
			down = mid + 1
		}
		return mid
	}
	return 0
}

//the search binary plus the initial index with end index,
