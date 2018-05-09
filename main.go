package main

import "fmt"

// Node : create a type node
type Node struct {
	Key   int
	Data  string
	Left  *Node
	Right *Node
}

// Tree : Create a new tree, starting from the Root, which is a type Node
type Tree struct {
	Root *Node
}

// InsertRoot : If there's no root on the tree, this function adds one
func (tree *Tree) InsertRoot(key int, data string) {
	tree.Root = &Node{
		Key:  key,
		Data: data,
	}
	return
}

// Insert : Inserts a new node on the tree, if there's already a Root, it calls InserNode
func (tree *Tree) Insert(key int, data string) {
	if tree.Root == nil {
		tree.InsertRoot(key, data)
		return
	}
	tree.Root.InsertNode(key, data)
}

// InsertNode : Inserts a new children on the node
func (node *Node) InsertNode(key int, data string) func() {
	switch {
	case key == node.Key:
		return nil
	case key < node.Key:
		if node.Left == nil {
			node.Left = &Node{
				Key:  key,
				Data: data,
			}
			return nil
		}
		return node.Left.InsertNode(key, data)
	case key > node.Key:
		if node.Right == nil {
			node.Right = &Node{
				Key:  key,
				Data: data,
			}
			return nil
		}
		return node.Right.InsertNode(key, data)
	}
	return nil
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

func display(tree *Tree) {
	if tree.Root != nil {
		fmt.Println("Root Key: ", tree.Root.Key)
		fmt.Println("Title Tree: ", tree.Root.Data)
	}
}

func main() {
	tree := new(Tree)
	arr := [6]int{1, 2, 3, 4, 5, 6}
	arrD := [6]string{"Pedro", "Juan", "Diego", "Nutella", "Rain", "Names"}
	for i := 0; i < len(arr); i++ {
		tree.InsertRoot(arr[i], arrD[i])
	}
	display(tree)
}
