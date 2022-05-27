package utils

import "sort"

func MinInt(array []int) int {
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	return array[0]
}
