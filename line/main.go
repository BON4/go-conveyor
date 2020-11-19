package line

import (
	items "conveyor/items"
)

type Node interface {
	GetItem() items.Item

	Next() *Node
	Prev() *Node

	SetNext(Node)
	SetPrev(Node)
}

type Line interface {
	Front() *Node
	Back() *Node
	AppendItem(items.Item)
}
