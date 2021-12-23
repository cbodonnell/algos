package algos

import "fmt"

func SearchList(list []string, terms ...string) {
	for _, term := range terms {
		index := -1
		for i, s := range list {
			if s == term {
				index = i
				break
			}
		}
		if index > -1 {
			fmt.Printf("%s found at index %d", term, index)
		} else {
			fmt.Printf("%s is not in list", term)
		}
	}
}
