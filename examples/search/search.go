package main

import (
	"fmt"

	"github.com/cheebz/algos"
)

func main() {
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
	index := algos.FindIndexLinear(haystack, term)
	if index > -1 {
		fmt.Printf("Found %s at index %d\n", term, index)
	} else {
		fmt.Printf("%s not in list\n", term)
	}
}
