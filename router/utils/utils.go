package utils

import (
	"fmt"
	"net/http"
)

type EndpointNode struct {
	Next  *EndpointNode
	Value func(http.ResponseWriter, *http.Request)
	Path  string
}

type EndpointLinkedList struct {
	Tail *EndpointNode
	Head *EndpointNode
}

func (LL *EndpointLinkedList) Add(path string, v func(http.ResponseWriter, *http.Request)) {
	NewNode := &EndpointNode{Path: path, Value: v}
	if LL.Head == nil {
		LL.Head = NewNode

	} else {
		cur := LL.Head

		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = NewNode
	}
}

func (LL *EndpointLinkedList) Print() {
	if LL.Head != nil {
		cur := LL.Head
		Index := 1
		fmt.Println(Index)
		for cur.Next != nil {
			Index++
			fmt.Println(Index)

			cur = cur.Next
		}
	}
}

type Node struct {
	Next  *Node
	Value func(http.ResponseWriter, *http.Request)
}

type LinkedList struct {
	Tail *Node
	Head *Node
}

func (LL *LinkedList) Add(v func(http.ResponseWriter, *http.Request)) {
	NewNode := &Node{Value: v}
	if LL.Head == nil {
		LL.Head = NewNode
	} else {
		cur := LL.Head

		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = NewNode
	}
}

func (LL *LinkedList) Print() {
	if LL.Head != nil {
		cur := LL.Head
		Index := 1
		fmt.Println(Index)
		for cur.Next != nil {
			Index++
			fmt.Println(Index)

			cur = cur.Next
		}
	}
}
