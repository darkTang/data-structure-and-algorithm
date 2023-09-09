package single_linkedList

import (
	"fmt"
	"strconv"
)

var str = ""

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Size int
	Head *ListNode
}

func (l *LinkedList) Append(val int) {
	newNode := &ListNode{
		Val:  val,
		Next: nil,
	}
	if l.Size == 0 {
		l.Head = newNode
	} else {
		lastNode := l.GetIndexNode(l.Size - 1)
		lastNode.Next = newNode
	}
	l.Size++
}

func (l *LinkedList) InsertAtIndex(i, val int) {
	newNode := &ListNode{
		Val:  val,
		Next: nil,
	}
	if i < 0 || i > l.Size {
		fmt.Println("index out of range")
		return
	}
	if i == 0 {
		newNode.Next = l.Head
		l.Head = newNode
	} else if i == l.Size {
		l.Append(val)
	} else {
		lastNode := l.GetIndexNode(i - 1)
		newNode.Next = lastNode.Next
		lastNode.Next = newNode
	}
	l.Size++
}

func (l *LinkedList) RemoveAtIndex(i int) {
	if i < 0 || i >= l.Size {
		fmt.Println("index out of range")
		return
	}
	if i == 0 {
		l.Head = l.Head.Next
	} else {
		lastNode := l.GetIndexNode(i - 1)
		lastNode.Next = lastNode.Next.Next
	}
}

func (l *LinkedList) GetIndexNode(index int) *ListNode {
	tempNode := l.Head
	for i := 0; i < index; i++ {
		tempNode = tempNode.Next
	}
	return tempNode
}

func (l *LinkedList) String() string {
	tempNode := l.Head
	for tempNode != nil {
		str = str + "->" + strconv.Itoa(tempNode.Val)
		tempNode = tempNode.Next
	}
	return str[2:]
}
