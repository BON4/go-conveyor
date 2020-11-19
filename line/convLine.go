package line

import items "conveyor/items"

type LineNode struct {
	Item     items.Item
	NextNode *Node
	PrevNode *Node
}

func newLineNode(item items.Item) Node {
	return &LineNode{Item: item}
}

func (ln *LineNode) GetItem() items.Item {
	return ln.Item
}

func (ln *LineNode) Next() *Node {
	return ln.NextNode
}

func (ln *LineNode) Prev() *Node {
	return ln.PrevNode
}

func (ln *LineNode) SetNext(n Node) {
	ln.NextNode = &n
}

func (ln *LineNode) SetPrev(n Node) {
	ln.PrevNode = &n
}

type ConveyorLine struct {
	First *Node
	Last  *Node
}

func NewConveyorLine() Line {
	return &ConveyorLine{}
}

func (cl *ConveyorLine) Front() *Node {
	return cl.First
}

func (cl *ConveyorLine) Back() *Node {
	return cl.Last
}

func (cl *ConveyorLine) append(n Node) {
	if cl.First == nil && cl.Last == nil {
		cl.First = &n
		cl.Last = &n
	} else {
		var temp Node = *cl.Last
		cl.Last = &n

		temp.SetNext(*cl.Last)
		(*cl.Last).SetPrev(temp)
	}
}

func (cl *ConveyorLine) AppendItem(item items.Item) {
	n := newLineNode(item)
	cl.append(n)
}
