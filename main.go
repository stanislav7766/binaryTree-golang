package main

import (
	"fmt"

	bst "github.com/stanislav7766/binaryTree-golang/tree"
)

func main() {
	myTree := &bst.Tree{}
	myTree.Insert(7, "7")
	myTree.Insert(2, "2")
	myTree.Insert(5, "5")
	myTree.Insert(1, "1")
	myTree.Insert(10, "10")
	myTree.Insert(15, "15")
	myTree.Insert(9, "9")
	myTree.Insert(3, "3")
	myTree.String()
	fmt.Printf("Min %s\n", myTree.Min())
	fmt.Printf("Max %s\n", myTree.Max())
	fmt.Printf("key 7 exist: %t\n", myTree.Search(7))
	fmt.Printf("key 40 exist: %t\n", myTree.Search(40))
	fmt.Printf("count: %d\n", myTree.GetCount())

	fmt.Println("after removing:")
	myTree.Remove(2)
	myTree.Remove(15)
	myTree.Remove(7)
	myTree.Remove(1)
	myTree.String()
	fmt.Printf("count: %d\n", myTree.GetCount())

}
