package main 

import "fmt"

const AlphabetSize = 26

type Node struct {
	children [26]*Node
	isEnd bool 
}

type Trie struct{
	root *Node 
}

func InitTrie() *Trie{
	result := &Trie{root:&Node{}}
	return result 
}

func (t *Trie) Insert(w string){
	wordLength := len(w)
	currentNode := t.root 
	for i := 0; i < wordLength; i++{
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil{
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(w string) bool{
	wordLength := len(w)
	currentNode := t.root 
	for i := 0; i < wordLength; i++{
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil{
			return false 
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd
}

func (t *Trie) Delete(word string) bool {
	if len(word) == 0 {
		return false
	}
	
	return deleteHelper(t.root, word, 0)
}

func deleteHelper(current *Node, word string, index int) bool {
	if index == len(word) {
		if !current.isEnd {
			return false
		}
		
		current.isEnd = false
		
		return hasNoChildren(current)
	}
	
	charIndex := word[index] - 'a'
	
	if current.children[charIndex] == nil {
		return false
	}
	
	shouldDeleteChild := deleteHelper(current.children[charIndex], word, index+1)
	
	if shouldDeleteChild {
		current.children[charIndex] = nil
		
		return !current.isEnd && hasNoChildren(current)
	}
	
	return false
}

func hasNoChildren(node *Node) bool {
	for i := 0; i < AlphabetSize; i++ {
		if node.children[i] != nil {
			return false
		}
	}
	return true
}

func main() {
	myTrie := InitTrie()
	
	toAdd := []string{
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oreo",
	}

	for _, v := range toAdd {
		myTrie.Insert(v)
	}
	
	fmt.Println("Search 'area':", myTrie.Search("area"))
	fmt.Println("Search 'argon':", myTrie.Search("argon"))
	
	fmt.Println("Deleting 'argon':", myTrie.Delete("argon"))
	fmt.Println("Search 'argon' after deletion:", myTrie.Search("argon"))
	
}
