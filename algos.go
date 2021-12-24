package algos

import (
	"math/rand"
)

func FindIndex(list []string, term string) int {
	for index, item := range list {
		if item == term {
			return index
		}
	}
	return -1
}

// Bubble sort a list of integers
func BubbleSort(list []int) []int {
	for {
		hasChanged := false
		for i, item := range list[0 : len(list)-1] {
			if item > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				hasChanged = true
			}
		}
		if !hasChanged {
			break
		}
	}
	return list
}

// Quick sort a list of integers
func QuickSort(list []int) []int {
	quickSortHelper(list, 0, len(list)-1)
	return list
}

func quickSortHelper(a []int, first int, last int) {
	if first >= last {
		return
	}
	pivotIndex := partition(a, first, last, rand.Intn(last-first+1)+first)
	quickSortHelper(a, first, pivotIndex-1)
	quickSortHelper(a, pivotIndex+1, last)
}

func partition(a []int, first int, last int, pivotIndex int) int {
	a[first], a[pivotIndex] = a[pivotIndex], a[first]
	left := first + 1
	right := last
	for left <= right {
		for left <= last && a[left] < a[first] {
			left++
		}
		for right >= first && a[first] < a[right] {
			right--
		}
		if left <= right {
			a[left], a[right] = a[right], a[left]
			left++
			right--
		}
	}
	a[first], a[right] = a[right], a[first]
	return right
}