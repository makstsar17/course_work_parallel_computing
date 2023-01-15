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

func (l *LinkedList) Insert(word string, docId string) {
	if l.Head == nil {
		l.Head = &Node{Word: word, DocIds: make(map[string]int)}
		l.Head.DocIds[docId] = 1
		l.Length += 1
		return

	} else {
		last := l.Head
		for true {
			if last.Word == word {
				last.DocIds[docId] += 1
				return
			}
			if last.Next == nil {
				break
			}
			last = last.Next
		}
		last.Next = &Node{Word: word, DocIds: make(map[string]int)}
		last.Next.DocIds[docId] = 1
		l.Length += 1
	}
}
