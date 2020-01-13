package node

// Node a single node that composes the tree
type Node struct {
	Key   int
	Value string
	Left  *Node //left
	Right *Node //right
}

//InsertNode internal function to find the correct place for a node in a tree
func InsertNode(node *Node, newNode *Node) {
	if newNode.Key < node.Key {
		if node.Left == nil {
			node.Left = newNode
		} else {
			InsertNode(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			InsertNode(node.Right, newNode)
		}
	}
}

// InOrderTraverse internal recursive function to traverse in order
func InOrderTraverse(node *Node, f func(string)) {
	if node != nil {
		InOrderTraverse(node, f)
		f(node.Value)
		InOrderTraverse(node, f)
	}
}

// PreOrderTraverse internal recursive function to traverse pre order
func PreOrderTraverse(node *Node, f func(string)) {
	if node != nil {
		f(node.Value)
		PreOrderTraverse(node.Left, f)
		PreOrderTraverse(node.Right, f)
	}
}

// PostOrderTraverse internal recursive function to traverse post order
func PostOrderTraverse(node *Node, f func(string)) {
	if node != nil {
		PostOrderTraverse(node.Left, f)
		PostOrderTraverse(node.Right, f)
		f(node.Value)
	}
}

// Search internal recursive function to search an item in the tree
func Search(node *Node, key int) bool {
	if node == nil {
		return false
	}
	if key < node.Key {
		return Search(node.Left, key)
	}
	if key > node.Key {
		return Search(node.Right, key)
	}
	return true
}

//Remove internal recursive function to remove an item
func Remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.Key {
		node.Left = Remove(node.Left, key)
		return node
	}
	if key > node.Key {
		node.Right = Remove(node.Right, key)
		return node
	}
	// key == node.key
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
	leftmostrightside := node.Right
	for {
		//find smallest value on the right side
		if leftmostrightside != nil && leftmostrightside.Left != nil {
			leftmostrightside = leftmostrightside.Left
		} else {
			break
		}
	}
	node.Key, node.Value = leftmostrightside.Key, leftmostrightside.Value
	node.Right = Remove(node.Right, node.Key)
	return node
}
