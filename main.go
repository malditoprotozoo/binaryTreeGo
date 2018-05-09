package main

import (
	"fmt"
)

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

// Insert : Inserts a new node on the tree, if there's already a Root, it calls InserNode
func (tree *Tree) Insert(key int, data string) func() {
	if tree.Root == nil {
		tree.Root = &Node{
			Key:  key,
			Data: data,
		}
		return nil
	}
	return tree.Root.InsertNode(key, data)
}

// InsertNode : Inserts a new children on the node
func (node *Node) InsertNode(key int, data string) func() {
	if node == nil {
		return nil
	}

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

// Find : Searches for a key and returns its Data
func Find(node *Node, key int) func() {
	if node.Key == key {
		fmt.Println(node.Data)
		return nil
	}
	if node.Key < key {
		return Find(node.Right, key)
	}
	return Find(node.Left, key)
}

func Remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.Key > key {
		node.Left = Remove(node.Left, key)
		return node
	}
	if node.Key < key {
		node.Right = Remove(node.Right, key)
		return node
	}
	if node.Left == nil && node.Right == nil {
		node = nil
		return nil
	}
	if node.Left == nil {
		node = node.Right
		return node
	}
	if node.Right == nil {
		node = node.Left
		return node
	}
	findSmall := node.Right
	for {
		if findSmall != nil && findSmall.Left != nil {
			findSmall = findSmall.Left
		} else {
			break
		}
	}
	node.Key, node.Data = findSmall.Key, findSmall.Data
	node.Right = Remove(node.Right, node.Key)
	return node
}

// COUNT : for display the tree
var COUNT = 10

func display(node *Node, space int) {
	if node != nil {
		space += COUNT
		display(node.Right, space)
		fmt.Print("\n")
		for i := COUNT; i < space; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("%d\n", node.Key)
		display(node.Left, space)
	}
	return
}

func main() {
	tree := new(Tree)
	keys := [7]int{10, 9, 122, 7, 345, 5, 4}
	datas := [7]string{"a", "b", "c", "d"}
	for i := 0; i < len(keys); i++ {
		tree.Insert(keys[i], datas[i])
	}
	display(tree.Root, 0)
}
