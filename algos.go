/*
Package algos provides some common utility algorithms
suited for general use in applications.
*/
package algos

import (
	"fmt"
	"math/rand"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) PreOrderTraversal() {
	if root != nil {
		fmt.Println(root.Val)
		root.Left.PreOrderTraversal()
		root.Right.PreOrderTraversal()
	}
}

// ListNode represents a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseLinkedList reverses a linked list
func ReverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

// unlinkList converts a linked list into a slice
func unlinkList(head *ListNode) (result []int) {
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return
}

// linkList converts a slice into a linked list
func linkList(s []int) (head *ListNode) {
	for i := len(s) - 1; i >= 0; i-- {
		head = &ListNode{
			Val:  s[i],
			Next: head,
		}
	}
	return
}

// FindIndexLinear returns the index of a string in a list
// or returns -1 if the string is not present in the list.
func FindIndexLinear(list []string, term string) int {
	for index, item := range list {
		if item == term {
			return index
		}
	}
	return -1
}

// FindIndexBinary finds the index of an integer in a sorted list
// or returns -1 if the integer is not present in the list
func FindIndexBinary(list []int, item int) int {
	lower, upper := 0, len(list)-1

	for lower <= upper {
		mid := (upper + lower) / 2
		guess := list[mid]

		if guess == item {
			return mid
		}

		if guess > item {
			upper = mid - 1
		} else {
			lower = mid + 1
		}
	}

	return -1
}

// InsertAtIndex inserts the item into the array at index
// and returns the modified array.
func InsertAtIndex(list []string, item string, index int) []string {
	list = append(list[:index+1], list[index:]...)
	list[index] = item
	return list
}

// RemoveAtIndex removes the item at index from the array
// and returns the modified array.
func RemoveAtIndex(list []string, index int) []string {
	return append(list[:index], list[index+1:]...)
}

// ReverseString reverses the letters in a string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ReverseInt reverses a 32 bit integer.
// Returns 0 if reversed integer overflows.
func ReverseInt(x int) int {
	var res int32
	for x != 0 {
		tail := int32(x) % 10
		newRes := res*10 + tail
		if (newRes-tail)/10 != res {
			return 0
		}
		res = newRes
		x = x / 10
	}
	return int(res)
}

// BubbleSort sorts a list of integers
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

// QuickSort sorts a list of integers
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

// Fibonacci returns the resulting fibonacci number n
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
