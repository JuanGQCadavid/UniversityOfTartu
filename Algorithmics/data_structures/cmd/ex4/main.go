package main

import (
	"log"

	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/trees/btree"
)

func nextCollatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func main() {

	var (
		maxValuesPerNode = 4
		start            = 51
		btree            = btree.NewBTree[int](maxValuesPerNode)
	)

	for n := start; n > 1; n = nextCollatz(n) {
		log.Println(n)
		btree.InsertValue(n, btree.Root)
		btree.PrintTree()
	}
	log.Println("Result: ")
	log.Println()

	btree.PrintTree()

}
