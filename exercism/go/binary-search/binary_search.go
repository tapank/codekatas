package binarysearch

func SearchInts(list []int, key int) int {
	for low, high := 0, len(list)-1; low <= high; {
		mid := (low + high) / 2
		if list[mid] > key {
			high = mid - 1
		} else if list[mid] < key {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
