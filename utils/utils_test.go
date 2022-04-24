package utils

import (
	"binaryTreeMaxPath/constants"
	"binaryTreeMaxPath/model"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InputToBinaryTree(t *testing.T) {
	assert := assert.New(t)

	input := model.Input{}
	json.Unmarshal([]byte(constants.TestCase1), &input)
	tree := InputToBinaryTree(input)

	actual := tree.Value
	expected := float64(1)
	assert.Equal(expected, actual, "Root node's value must be 1")

	actual = tree.Right.Left.Value
	expected = float64(6)
	assert.Equal(expected, actual, "Value of the root node right child's left child must be 6")

	json.Unmarshal([]byte(constants.TestCase3), &input)
	tree = InputToBinaryTree(input)

	actual = tree.Left.Right.Left.Value
	expected = float64(3)
	assert.Equal(expected, actual, "Value of the root node left child's right child's left child must be 3")

	json.Unmarshal([]byte(constants.UnOrderedTestCase3), &input)
	tree = InputToBinaryTree(input)

	actual = tree.Left.Right.Left.Value
	expected = float64(3)
	assert.Equal(expected, actual, "It must handle unordered inputs")
}

func Test_FindMaxPath(t *testing.T) {
	assert := assert.New(t)

	bn := model.BinaryNode{
		Left:    nil,
		Right:   nil,
		Value:   5,
		MaxPath: 0,
	}

	actual := FindMaxPath(&bn)
	expected := bn.Value

	assert.Equal(expected, actual, "leaf's MaxPathValue must be equal to it's value")

	bn.Left = &model.BinaryNode{
		Left:    nil,
		Right:   nil,
		Value:   3,
		MaxPath: 0,
	}
	bn.Right = &model.BinaryNode{
		Left:    nil,
		Right:   nil,
		Value:   2,
		MaxPath: 0,
	}
	actual = FindMaxPath(&bn)
	expected = float64(8)

	assert.Equal(expected, actual, "Node's MaxPathValue must be equal to sum of the bigger value child's and it's own value")

	bn.Left = nil
	actual = FindMaxPath(&bn)

	expected = float64(7)
	assert.Equal(expected, actual, "node's child can be nil. It can not change the result")
}

func Test_Print(t *testing.T) {
	assert := assert.New(t)

	bn := model.BinaryNode{
		Left:    nil,
		Right:   nil,
		Value:   5,
		MaxPath: 0,
	}
	actual := PrintTreeByValue(&bn)
	expected := "5 / "
	assert.Equal(expected, actual)

	actual = PrintTreeByMaxPath(&bn)
	expected = "0 / "
	assert.Equal(expected, actual)

	bnLeft := model.BinaryNode{
		Left:    nil,
		Right:   nil,
		Value:   1,
		MaxPath: 0,
	}
	bnRight := model.BinaryNode{
		Left:    nil,
		Right:   nil,
		Value:   2,
		MaxPath: 0,
	}
	bn.Right = &bnRight
	bn.Left = &bnLeft
	bn.MaxPath = 2

	actual = PrintTreeByValue(&bn)
	expected = "1 / 5 / 2 / "
	assert.Equal(expected, actual)

	actual = PrintTreeByMaxPath(&bn)
	expected = "0 / 2 / 0 / "
	assert.Equal(expected, actual)

}
