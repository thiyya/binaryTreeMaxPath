package model

type Node struct {
	Id    string `json: "id"`
	Left  string `json: "left"`
	Right string `json: "right"`
	Value int    `json: "value"`
}

type Tree struct {
	Nodes []Node `json: "nodes"`
	Root  string `json: "root"`
}

type Input struct {
	Tree Tree `json: "tree"`
}

type BinaryNode struct {
	Left    *BinaryNode
	Right   *BinaryNode
	Value   float64
	MaxPath float64
}
type Result struct {
	MaxPathValue float64 `json: "maxPathValue"`
}
