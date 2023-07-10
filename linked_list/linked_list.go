package linked_list

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type List struct {
	Head *Node
}

func ListLen(l *List) int {
	current := l.Head
	count := 0
	for current != nil {
		count += 1
		current = current.Next
	}
	return count
}

func InsertAt(l *List, value int, position int) int {
	newNode := Node{Val: value, Next: nil}
	if position == 0 {
		if l.Head == nil {
			l.Head = &newNode
		} else {
			newNode.Next = l.Head
			l.Head = &newNode
		}
		return value
	}

	currentNode := l.Head
	for i := 1; i < position && currentNode.Next != nil; i += 1 {
		currentNode = currentNode.Next
	}

	if currentNode.Next == nil {
		currentNode.Next = &newNode
	} else {
		newNode.Next = currentNode.Next
		currentNode.Next = &newNode
	}
	return value
}

func Remove(l *List, value int) int {
	currentNode := l.Head
	if currentNode != nil && currentNode.Val == value {
		l.Head = currentNode.Next
		return value
	}

	for currentNode.Next != nil {
		if currentNode.Next.Val == value {
			currentNode.Next = currentNode.Next.Next
			return value
		}
		currentNode = currentNode.Next
	}
	return -1
}

func RemoveAt(l *List, position int) int {
	if position == 0 {
		if l.Head == nil {
			return -1
		} else {
			l.Head = l.Head.Next
		}
		return position
	}

	currentNode := l.Head
	for i := 1; i <= position && currentNode != nil; i += 1 {
		if i == position {
			currentNode.Next = currentNode.Next.Next
			return position
		}
		currentNode = currentNode.Next
	}
	return -1
}

func Append(l *List, value int) int {
	ln := ListLen(l)
	InsertAt(l, value, ln)
	return value
}

func Prepend(l *List, value int) int {
	InsertAt(l, value, 0)
	return value
}

func GetItem(l *List, position int) int {
	currentNode := l.Head
	for i := 0; i < position; i += 1 {
		currentNode = currentNode.Next
	}
	return currentNode.Val
}

func PrintList(l *List) {
	currentNode := l.Head
	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.Val)
		currentNode = currentNode.Next
	}
	fmt.Println()
}
