package linkedList

type Node struct {
	Word   string
	DocIds map[string]int // key - document's name, value - count of appearance
	Next   *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}
