package main

import "fmt"

// Tree : Create a new tree, starting from the Root, which is a type Node
type Tree struct {
	Root *Node
}

// Node : create a type node
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// InsertRoot : If there's no root on the tree, this function adds one
func (tree *Tree) InsertRoot(key int) {
	tree.Root.Key = key
	return
}

// InsertNode : Insert a new node on the tree
func (tree *Tree) InsertNode(key int) {
	if tree.Root == nil {

	}
}

func printNode(node *Node) {
	fmt.Print("Key: ", node.Key)
	if node.Left != nil {
		fmt.Print(" Left: ", node.Left.Key)
	} else {
		fmt.Print(" Left is empty ")
	}
	if node.Right != nil {
		fmt.Print(" Right: ", node.Right.Key)
	} else {
		fmt.Print(" Right is empty ")
	}
	fmt.Println()
}

func read() []Node {
	var input int
	fmt.Scanf("%d", &input)
	var nodes = make([]Node, input)
	for i := 0; i < input; i++ {
		var key, indexLeft, indexRight int
		fmt.Scanf("%d %d %d", &key, &indexLeft, &indexRight)
		nodes[i].Key = key
		if indexLeft >= 0 {
			nodes[i].Left = &nodes[indexLeft]
		}
		if indexRight >= 0 {
			nodes[i].Right = &nodes[indexRight]
		}
	}
	return nodes
}

func main() {
	tree := new(Tree)
	fmt.Println(tree.Root == nil)
}
