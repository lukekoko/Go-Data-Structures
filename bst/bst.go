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

func (tree *bst) delete(val int) {
	tree.deleteNode(val, tree.root)
}

func (tree *bst) deleteNode(val int, root *Node) *Node {
	if root == nil {
		return nil
	}
	if val < root.data {
		root.left = tree.deleteNode(val, root.left)
	} else if val > root.data {
		root.right = tree.deleteNode(val, root.right)
	} else {
		if root.left == nil && root.right == nil {
			root = nil
		} else if root.left != nil && root.right != nil {
			min := tree.min(root.right)
			root.data = min.data
			root.right = tree.deleteNode(min.data, root.right)
		} else if root.left == nil {
			*root = *root.right
		} else if root.right == nil {
			*root = *root.left
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
	tree.insert(14)
	tree.insert(9)
	tree.insert(8)
	tree.delete(7)
	tree.delete(8)
	tree.delete(9)
	tree.delete(14)
	tree.delete(1)
	tree.delete(3)
	tree.preorder(tree.root)

}
