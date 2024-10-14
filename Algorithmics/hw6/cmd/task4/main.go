package main

import (
	"log"

	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/trees/binarysearch"
	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/trees/search"
)

func nextCollatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func main() {
	var (
		start = 17
		btree = binarysearch.NewBSTree[int, int]()
	)

	for n := start; n > 1; n = nextCollatz(n) {
		log.Println(n)
		btree.Append(n, n)
	}

	representative_vals := search.BreadthFirstSearchAll(btree.Root, btree.NodeCounter(btree.Root))
	log.Println(representative_vals)
	// visualizer.FromBSTToPNG(btree, "ex1")

	bit_vector := make([]byte, len(representative_vals))

	for index := range bit_vector {
		node := btree.TreeSearch(btree.Root, representative_vals[index])

		if node == nil {
			panic("Node not found")
		}

		if node.Rigth != nil || node.Left != nil {
			bit_vector[index] = 1
		}

	}

	log.Println(bit_vector)

}
