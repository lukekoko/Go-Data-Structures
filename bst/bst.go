package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type bst struct {
	root *Node
}

func (tree *bst) new() {
	tree.root = nil
}

func (tree *bst) insert(val int) {
	node := Node{data: val, left: nil, right: nil}
	if tree.root == nil {
		tree.root = &node
	} else {
		current := tree.root
		parent := current
		for {
			parent = current
			if val < parent.data {
				current = current.left
				if current == nil {
					parent.left = &node
					return
				}
			} else {
				current = current.right
				if current == nil {
					parent.right = &node
					return
				}
			}
		}
	}
}

func (tree *bst) search(val int, node *Node) *Node {
	if node == nil || node.data == val {
		return node
	}
	if val < node.data {
		return tree.search(val, node.left)
	} else {
		return tree.search(val, node.right)
	}
}

func (tree *bst) delete(val int, root *Node) *Node {
	if root == nil {
		return root
	}
	if val < root.data {
		root.left = tree.delete(val, root.left)
	} else if val > root.data {
		root.right = tree.delete(val, root.right)
	} else {
		if root.left == nil && root.right == nil {
			root = nil
			return root
		} else if root.left != nil && root.right != nil {
			root.data = tree.min(root.right).data
			root.right = tree.delete(val, root.right)
			return root
		} else if root.left == nil {
			root = root.right
			return root
		} else if root.right == nil {
			root = root.left
			return root
		}
	}
	return root
}

func (tree *bst) inorder(node *Node) {
	if node == nil {
		return
	}
	tree.inorder(node.left)
	fmt.Println(node.data)
	tree.inorder(node.right)
}

func (tree *bst) preorder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.data)
	tree.inorder(node.left)
	tree.inorder(node.right)
}

func (tree *bst) min(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return tree.min(node.left)
}

func main() {
	tree := bst{}
	tree.new()
	tree.insert(7)
	tree.insert(1)
	tree.insert(3)
	tree.insert(8)
	tree.delete(1, tree.root)
	tree.delete(3, tree.root)
	tree.delete(8, tree.root)
	tree.delete(7, tree.root)
	tree.preorder(tree.root)

}
