package main

import (
	"fmt"
	"os"

	"github.com/cheebz/algos"
)

func main() {
	var s string
	if len(os.Args) > 1 {
		s = os.Args[1]
	} else {
		s = "!dlrow ,olleh"
	}
	reversed := algos.ReverseString(s)
	fmt.Println(reversed)
}
