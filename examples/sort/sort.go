package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cheebz/algos"
)

func main() {
	rand.Seed(time.Now().Unix())
	list := rand.Perm(20)
	fmt.Println("unsorted:	", list)
	fmt.Println()

	start := time.Now()
	algos.BubbleSort(list)
	fmt.Println("sorted!		", list)
	fmt.Println("bubble:		", time.Since(start))
	fmt.Println()

	start = time.Now()
	algos.QuickSort(list)
	fmt.Println("sorted!		", list)
	fmt.Println("quick:		", time.Since(start))
}
