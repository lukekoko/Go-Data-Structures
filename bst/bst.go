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

func (tree *bst) delete() {
	tree.root = nil
}

func (tree *bst) remove(val int) {
	tree.removeNode(val, tree.root)
}

func (tree *bst) removeNode(val int, root *Node) *Node {
	if root == nil {
		return nil
	}
	if val < root.data {
		root.left = tree.removeNode(val, root.left)
	} else if val > root.data {
		root.right = tree.removeNode(val, root.right)
	} else {
		if root.left == nil && root.right == nil {
			root = nil
		} else if root.left != nil && root.right != nil {
			min := tree.min(root.right)
			root.data = min.data
			root.right = tree.removeNode(min.data, root.right)
		} else if root.left == nil {
			*root = *root.right
		} else if root.right == nil {
			*root = *root.left
		}
	}
	return root
}

func (tree *bst) print(traversal string) {
	if traversal == "inorder" {
		tree.inorder(tree.root)
	}
	if traversal == "preorder" {
		tree.preorder(tree.root)
	}
	if traversal == "postorder" {
		tree.postorder(tree.root)
	}
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

func (tree *bst) postorder(node *Node) {
	if node == nil {
		return
	}
	tree.inorder(node.left)
	tree.inorder(node.right)
	fmt.Println(node.data)
}

func (tree *bst) min(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return tree.min(node.left)
}

func (tree *bst) max(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return tree.max(node.right)
}

func (tree *bst) isBST(node *Node, min int, max int) bool {
	if node == nil {
		return true
	}
	if node.data < min || node.data > max {
		return false
	}
	return tree.isBST(node.left, min, node.data-1) && tree.isBST(node.right, node.data+1, max)
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
	tree.print("inorder")
	tree.delete()
	tree.print("postorder")
	tree.insert(2)
	
}
