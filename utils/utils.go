package utils

import (
	"binaryTreeMaxPath/model"
	"fmt"
	"math"
)

func InputToBinaryTree(input model.Input) *model.BinaryNode {
	itemMap := map[string]*model.BinaryNode{}
	for _, node := range input.Tree.Nodes {
		itemMap[node.Id] = &model.BinaryNode{Value: float64(node.Value)}
	}

	for _, node := range input.Tree.Nodes {
		itemMap[node.Id].Left = itemMap[node.Left]
		itemMap[node.Id].Right = itemMap[node.Right]
	}
	return itemMap[input.Tree.Root]
}

func FindMaxPath(node *model.BinaryNode) float64 {
	if node == nil {
		return 0
	}
	left := math.Max(FindMaxPath(node.Left), 0)
	right := math.Max(FindMaxPath(node.Right), 0)
	node.MaxPath = math.Max(left, right) + node.Value
	return node.MaxPath
}

func PrintTreeByValue(node *model.BinaryNode) string {
	result := ""
	if node == nil {
		return ""
	}
	result += PrintTreeByValue(node.Left)
	result += fmt.Sprintf("%v / ", int(node.Value))
	result += PrintTreeByValue(node.Right)
	return result
}

func PrintTreeByMaxPath(node *model.BinaryNode) string {
	result := ""
	if node == nil {
		return ""
	}
	result += PrintTreeByMaxPath(node.Left)
	result += fmt.Sprintf("%v / ", int(node.MaxPath))
	result += PrintTreeByMaxPath(node.Right)

	return result
}
