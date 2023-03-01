package engine

type Node interface {
}

type NodeType int

const (
	NodeText NodeType = iota
	NodeTable
	NodeTableField
	NodeList
	NodeRepeat
	NodeLeftDelim
	NodeRightDelim
)

type TextNode struct {
	NodeType
	Pos
	tr   *Tree
	Text []byte
}

type TableNode struct {
	NodeType
	Pos
	tr       *Tree
	Fields   []*TableFieldNode
	Name     string
	Repeated bool
}

type TableFieldNode struct {
	NodeType
	Pos
	tr  *Tree
	Val string
}

type ListNode struct {
	NodeType
	Pos
	tr    *Tree
	Nodes []Node
}
