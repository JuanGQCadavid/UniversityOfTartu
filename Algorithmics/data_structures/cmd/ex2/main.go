package main

import (
	"fmt"
	"log"

	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/internal/domain"
	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/trees/binarysearch"
	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/trees/search"
	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/trees/visualizer"
	"golang.org/x/exp/constraints"
)

func nextCollatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func main() {
	// start := 7
	// visualize(start)
	// printString(start)

	wouldItWork()
}

func wouldItWork() {
	nodesToAppend := []int{
		-10, 8, -1, -20, -30, 9, 1, 5,
	}
	bst := binarysearch.NewBSTree[int, int]()
	for _, node := range nodesToAppend {
		bst.Append(node, node)
	}

	sumL, nodesL := search.MaximunPathSum(bst.Root.Left)
	sumR, nodesR := search.MaximunPathSum(bst.Root.Rigth)
	log.Println("sumL:", sumL)
	log.Println("sumR:", sumR)
	log.Println("Total Sum: ", sumL+sumR+bst.Root.SatalliteData)

	totalNodes := make([]*domain.Node[int, int], 0, len(nodesL)+len(nodesR)+1)

	totalNodes = append(totalNodes, nodesL...)
	totalNodes = append(totalNodes, bst.Root)
	totalNodes = append(totalNodes, nodesR...)

	root := BuildNewTreeFromSubtree(totalNodes)

	fileName := "ex2.csv"
	result := visualizer.FromTreeToMatrix(root, len(totalNodes))
	visualizer.FromMatrixToFile(result, fileName)
	visualizer.FromCSVToImage(fileName)

}

func visualize(start int) {
	results := make([]int, 10)
	bst := binarysearch.NewBSTree[int, int]()
	for n := start; n > 1; n = nextCollatz(n) {
		results = append(results, n)
		log.Println(n)
		bst.Append(n, n)
	}

	sumL, nodesL := search.MaximunPathSum(bst.Root.Left)
	sumR, nodesR := search.MaximunPathSum(bst.Root.Rigth)
	log.Println("sumL:", sumL)
	log.Println("sumR:", sumR)
	log.Println("Total Sum: ", sumL+sumR+bst.Root.SatalliteData)

	totalNodes := make([]*domain.Node[int, int], 0, len(nodesL)+len(nodesR)+1)

	totalNodes = append(totalNodes, nodesL...)
	totalNodes = append(totalNodes, bst.Root)
	totalNodes = append(totalNodes, nodesR...)

	root := BuildNewTreeFromSubtree(totalNodes)

	fileName := "ex2.csv"
	result := visualizer.FromTreeToMatrix(root, len(totalNodes))
	visualizer.FromMatrixToFile(result, fileName)
	visualizer.FromCSVToImage(fileName)
}

func printString(start int) {
	results := make([]int, 10)
	bst := binarysearch.NewBSTree[int, int]()
	for n := start; n > 1; n = nextCollatz(n) {
		results = append(results, n)
		log.Println(n)
		bst.Append(n, n)
	}

	log.Println("Nodes: ", results)

	sumL, nodesL := search.MaximunPathSum(bst.Root.Left)
	sumR, nodesR := search.MaximunPathSum(bst.Root.Rigth)
	log.Println("sumL:", sumL)
	log.Println("sumR:", sumR)
	log.Println("Total Sum: ", sumL+sumR+bst.Root.SatalliteData)

	totalNodes := make([]*domain.Node[int, int], 0, len(nodesL)+len(nodesR)+1)

	totalNodes = append(totalNodes, nodesL...)
	totalNodes = append(totalNodes, bst.Root)
	totalNodes = append(totalNodes, nodesR...)
	newPointer := len(totalNodes) - 1
	for i, node := range totalNodes {
		if i == len(nodesL) {
			fmt.Print("(", node.SatalliteData, "), ")
			continue
		}

		if i > len(nodesL) {
			// pos := (len(nodesL) - 1 + len(nodesR) - 1) - (len(nodesL) - 1) + i
			fmt.Print(totalNodes[newPointer].SatalliteData, ", ")
			newPointer--
			continue
		}

		fmt.Print(node.SatalliteData, ", ")
	}
}

func BuildNewTreeFromSubtree[K constraints.Ordered, D any](nodes []*domain.Node[K, D]) *domain.Node[K, D] {
	// Create a map to track which nodes are in the list
	nodeMap := make(map[*domain.Node[K, D]]bool)

	// Populate the map with nodes in the list
	for _, node := range nodes {
		nodeMap[node] = true
	}

	var root *domain.Node[K, D] = nil

	// Iterate through nodes and clean up the pointers
	for _, node := range nodes {
		// Remove Parent, Left, and Right pointers that are not in the nodeMap
		if node.Parent != nil && !nodeMap[node.Parent] {
			node.Parent = nil
		}
		if node.Left != nil && !nodeMap[node.Left] {
			node.Left = nil
		}
		if node.Rigth != nil && !nodeMap[node.Rigth] {
			node.Rigth = nil
		}
		// If the node has no parent, it's the root node
		if node.Parent == nil {
			root = node
		}
	}

	return root
}
