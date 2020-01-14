package tree

import (
	"fmt"

	node "github.com/stanislav7766/binaryTree-golang/node"
)

//Tree struct
type Tree struct {
	Root  *node.Node
	Count int
}

// Insert inserts the data in the tree
func (bst *Tree) Insert(key int, value string) {
	n := &node.Node{key, value, nil, nil}
	if bst.Root == nil {
		bst.Root = n
	} else {
		node.InsertNode(bst.Root, n)
	}
	bst.Count++
}

// InOrderTraverse visits all nodes with in-order traversing dfs
func (bst *Tree) InOrderTraverse(f func(n *node.Node)) {
	if bst.Root == nil {
		return
	}
	node.InOrderTraverse(bst.Root, f)
}

//GetCount return count of nodes
func (bst *Tree) GetCount() int {
	return bst.Count
}

// PreOrderTraverse visits all nodes with pre-order traversing dfs
func (bst *Tree) PreOrderTraverse(f func(n *node.Node)) {
	if bst.Root == nil {
		return
	}
	node.PreOrderTraverse(bst.Root, f)
}

//Clear remove all nodes
func (bst *Tree) Clear() {
	bst.PostOrderTraverse(func(n *node.Node) {
		if n == bst.Root {
			bst.Root = nil
		}
		bst.Count--
		n.Left = nil
		n.Right = nil
		n = nil
	})
}

// PostOrderTraverse visits all nodes with post-order traversing dfs
func (bst *Tree) PostOrderTraverse(f func(n *node.Node)) {
	if bst.Root == nil {
		return
	}
	node.PostOrderTraverse(bst.Root, f)
}

// Min returns the Item with min value stored in the tree
func (bst *Tree) Min() string {
	n := bst.Root
	if n == nil {
		return ""
	}
	for {
		if n.Left == nil {
			return n.Value
		}
		n = n.Left
	}
}

// Max returns the Item with max value stored in the tree
func (bst *Tree) Max() string {
	n := bst.Root
	if n == nil {
		return ""
	}
	for {
		if n.Right == nil {
			return n.Value
		}
		n = n.Right
	}
}

// Search returns true if the Item t exists in the tree
func (bst *Tree) Search(key int) bool {
	if bst.Root == nil {
		return false
	}
	return node.Search(bst.Root, key)
}

// Remove removes the Item with key `key` from the tree
func (bst *Tree) Remove(key int) {
	if bst.Root == nil {
		return
	}
	if node.Remove(bst.Root, key) == nil {
		bst.Root = nil
	}
	bst.Count--

}
func (bst *Tree) String() {
	fmt.Println("------------------------------------------------")
	if bst.Root == nil {
		fmt.Println("The tree is Empty")
	}
	stringify(bst.Root, 0)
	fmt.Println("------------------------------------------------")
}
func stringify(n *node.Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.Right, level)
		fmt.Printf(format+"%s\n", n.Value)
		stringify(n.Left, level)
	}

}
