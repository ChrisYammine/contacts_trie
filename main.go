package main

import (
	"fmt"
)

const (
	NumberOfNodeChildren = 26 // a through z
)

type Node struct {
	num_words int
	children  []*Node
}

func NewNode() *Node {
	return &Node{0, make([]*Node, NumberOfNodeChildren)}
}

func GetCharIndex(char string) int {
	c := []byte(char)
	return int(c[0] - 'a')
}

func CharAt(s string, i int) string {
	return string(s[i])
}

func (n *Node) GetNode(char string) *Node {
	return n.children[GetCharIndex(char)]
}

func (n *Node) SetNode(char string, node *Node) {
	n.children[GetCharIndex(char)] = node
}

func (n *Node) Insert(s string) {
	node := n
	index := 0
	length := len(s)

	for index < length {
		if newNode := node.GetNode(string(s[index])); newNode != nil {
			node = newNode
			index++
		} else {
			break
		}
	}

	for index < length {
		val := string(s[index])
		newNode := NewNode()
		node.SetNode(val, newNode)
		node = newNode
		index++
	}
}

func main() {
	trie := NewNode()
	fmt.Println(trie)
	trie.Insert("animal")
	fmt.Println(trie)
}
