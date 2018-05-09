package utils

import "strconv"

// ContainsInt64 tells whether a slice contains x.
func ContainsInt64(a []int64, x int64) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// SliceAtoi -> convert array of string to array of integer
func SliceAtoi(s []string) ([]int, bool) {
	var status = true
	var arr []int

	for _, val := range s {
		if i, err := strconv.Atoi(val); err != nil {
			status = false
		} else {
			arr = append(arr, i)
		}
	}
	return arr, status
}
