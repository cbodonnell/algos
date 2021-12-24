package algos

import (
	"reflect"
	"testing"
)

// TestFindIndex calls algos.TestFindIndex with a list of strings
// and returns the index of the search term.
// If the term is not in the list, -1 is returned.
func TestFindIndex(t *testing.T) {
	var haystack = []string{"Zig", "Zag", "Wally", "Ronald", "Bush", "Krusty",
		"Charlie", "Bush", "Bozo", "Zag", "mouse", "hat", "cup", "deodorant",
		"television", "soap", "methamphetamine", "severed cat heads", "foo",
		"bar", "baz", "quux", "quuux", "quuuux", "bazola", "ztesch", "foo",
		"bar", "thud", "grunt", "foo", "bar", "bletch", "foo", "bar", "fum",
		"fred", "jim", "sheila", "barney", "flarp", "zxc", "spqr", ";wombat",
		"shme", "foo", "bar", "baz", "bongo", "spam", "eggs", "snork", "foo",
		"bar", "zot", "blarg", "needle", "toto", "titi", "tata", "tutu", "pippo",
		"pluto", "paperino", "aap", "noot", "mies", "oogle", "foogle", "boogle",
		"zork", "gork", "bork", "sodium", "phosphorous", "californium",
		"copernicium", "gold", "thallium", "carbon", "silver", "gold", "copper",
		"helium", "sulfur"}

	term := "needle"
	index := 56
	testIndex := FindIndex(haystack, term)
	if testIndex != 56 {
		t.Fatalf("%s index %d is not %d\n", term, testIndex, index)
	}

	term = "nothing"
	index = -1
	testIndex = FindIndex(haystack, term)
	if testIndex != -1 {
		t.Fatalf("%s index %d is not %d\n", term, testIndex, index)
	}
}

// TestBubbleSort calls algos.BubbleSort with a list of integers,
// and checks that it was sorted correctly
func TestBubbleSort(t *testing.T) {
	unsorted := []int{5, 3, 7, 2, 9}
	sorted := []int{2, 3, 5, 7, 9}
	testSort := BubbleSort(unsorted)
	if !reflect.DeepEqual(testSort, sorted) {
		t.Fatalf("%v and %v are not equal", testSort, sorted)
	}
}

// TestQuickSort calls algos.QuickSort with a list of integers,
// and checks that it was sorted correctly
func TestQuickSort(t *testing.T) {
	unsorted := []int{5, 3, 7, 2, 9}
	sorted := []int{2, 3, 5, 7, 9}
	testSort := QuickSort(unsorted)
	if !reflect.DeepEqual(testSort, sorted) {
		t.Fatalf("%v and %v are not equal", testSort, sorted)
	}
}
