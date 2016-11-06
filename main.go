package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	NumberOfNodeChildren = 26 // a through z
)

type Node struct {
	NumWords int
	Children []*Node
}

type Trie struct {
	Root *Node
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
	return n.Children[GetCharIndex(char)]
}

func (n *Node) SetNode(char string, node *Node) {
	n.Children[GetCharIndex(char)] = node
}

func (t *Trie) Insert(s string) {
	node := t.Root
	index := 0
	length := len(s)

	for index < length {
		if newNode := node.GetNode(string(s[index])); newNode != nil {
			node = newNode
			node.NumWords++
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
		node.NumWords++
		index++
	}
}

func (t *Trie) Find(s string) (*Node, bool) {
	node := t.Root

	for _, char := range s {
		n := node.GetNode(string(char))
		if n != nil {
			node = n
		} else {
			return nil, false
		}
	}

	return node, true
}

// While technically a correct way to count the number of 'complete' words -- it isn't a representative
// count of all words _added_ that a partial matches.

// func (t *Trie) CountPartialMatches(s string) int {
//   var count int
//   if node, ok := t.Find(s); ok {
//     count = node.NumWords
//     count += node.ChildCounts()
//   } else {
//     return count
//   }
//   return count
// }
//
// func (n *Node) ChildCounts() int {
//   var count int
//   for _, node := range(n.Children) {
//     if node != nil {
//       count += node.NumWords
//       count += node.ChildCounts()
//     }
//   }
//   return count
// }

func main() {
	trie := &Trie{NewNode()}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // let's see how many commands we have
	n_cmds, _ := strconv.Atoi(scanner.Text())
	commands := make([]string, n_cmds)

	for i, _ := range commands {
		scanner.Scan()
		commands[i] = scanner.Text()
	}

	for _, cmd := range commands {
		s := strings.Split(cmd, " ")
		switch s[0] {
		case "add":
			trie.Insert(s[1])
		case "find":
			if node, ok := trie.Find(s[1]); ok {
				fmt.Println(node.NumWords)
			} else {
				fmt.Println(0)
			}
		}
	}

}
