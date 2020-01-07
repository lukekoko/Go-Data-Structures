package binary_search_tree

type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

type bst struct {
	root *Node
}

func (tree *bst) new() {
	tree.root = nil
}

func (tree *bst) insert(val interface{}) {
	node := Node{data: val, left: nil, right: nil}
	if tree.root == nil {
		tree.root = &node
	} else {
		root := tree.root
	}
}
