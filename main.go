package main

import (
	"fmt"
	"log"
)

func main() {
	tree := &Tree{}
	err := tree.Insert("s", "stas")
	err = tree.Insert("b", "banna")
	err = tree.Insert("a", "anna")
	if err != nil {
		log.Fatal("Error inserting value '", "s", "': ", err)
	}
	fmt.Println(tree.Root.Value)

}
