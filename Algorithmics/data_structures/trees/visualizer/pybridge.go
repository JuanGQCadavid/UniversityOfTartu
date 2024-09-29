package visualizer

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/internal/domain"
	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/trees/binarysearch"
	"github.com/gocarina/gocsv"
	"golang.org/x/exp/constraints"
)

var (
	treeToPNGPythonCode = "trees/visualizer/pycode/print_tree.py"
)

type NodeData struct {
	Info       string `csv:"info"`
	LeftIndex  int    `csv:"left"`
	RigthIndex int    `csv:"right"`
	Color      string `csv:"color"`
}

type PendingNode[K constraints.Ordered, D any] struct {
	Node  *domain.Node[K, D]
	Index int
}

func FromBSTToPNG[K constraints.Ordered, D any](bst *binarysearch.BSTrea[K, D], fileName string) {
	result := FromTreeToMatrix(bst.Root, bst.NodeCounter(bst.Root))
	FromMatrixToFile(result, fileName+".csv")
	FromCSVToImage(fileName + ".csv")

}

func FromCSVToImage(fileName string) {
	cmd := exec.Command("python3", treeToPNGPythonCode, fileName)

	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal("python3 response: ", string(out))
		log.Fatal("python3 Err: ", err.Error())
	}

	log.Println("python3 response: ", string(out))
}

func FromMatrixToFile(matrix []NodeData, a string) {
	gFile, err := os.OpenFile(a, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	if err = gocsv.MarshalFile(matrix, gFile); err != nil {
		panic(err.Error())
	}
}

func FromTreeToMatrix[K constraints.Ordered, D any](root *domain.Node[K, D], nodesCount int) []NodeData {
	matrix := make([]NodeData, nodesCount)
	actualIndex := 0
	pendingNodes := make([]PendingNode[K, D], 0, nodesCount)
	pendingNodes = append(pendingNodes, PendingNode[K, D]{
		Node:  root,
		Index: actualIndex,
	})

	for len(pendingNodes) > 0 {
		pendingNode := pendingNodes[0]
		index := pendingNode.Index
		log.Println(index)

		// Rezise
		pendingNodes = pendingNodes[1:]

		matrix[index] = NodeData{
			Info:  fmt.Sprintf("%+v", pendingNode.Node.SatalliteData),
			Color: "blue",
		}

		if pendingNode.Node.Left != nil {
			actualIndex += 1
			matrix[pendingNode.Index].LeftIndex = actualIndex
			pendingNodes = append(pendingNodes, PendingNode[K, D]{
				Node:  pendingNode.Node.Left,
				Index: actualIndex,
			})
		}

		if pendingNode.Node.Rigth != nil {
			actualIndex += 1
			matrix[pendingNode.Index].RigthIndex = actualIndex
			pendingNodes = append(pendingNodes, PendingNode[K, D]{
				Node:  pendingNode.Node.Rigth,
				Index: actualIndex,
			})
		}
	}

	return matrix
}
