package engine

import "fmt"

type Node interface {
	Type() NodeType
	String() string
	Position() Pos
	tree() *Tree
}

type NodeType int

const (
	NodeText NodeType = iota
	NodeTable
	NodeTableField
	NodeList
)

// 验证TextNode是否实现Node接口
var _ Node = (*TextNode)(nil)

type TextNode struct {
	NodeType
	Pos
	tr   *Tree
	Text string
}

func (t *Tree) newText(pos Pos, text string) *TextNode {
	return &TextNode{
		NodeType: NodeText,
		Pos:      pos,
		Text:     text,
		tr:       t,
	}
}

func (t *TextNode) Type() NodeType {
	return t.NodeType
}

func (t *TextNode) Position() Pos {
	return t.Pos
}

func (t *TextNode) String() string {
	return fmt.Sprintf("%s", t.Text)
}

func (t *TextNode) tree() *Tree {
	return t.tr
}

type TableNode struct {
	NodeType
	Pos
	tr       *Tree
	Fields   []*TableFieldNode
	Name     string
	Repeated bool
}

func (t *Tree) newTable(pos Pos, name string) *TableNode {
	return &TableNode{
		NodeType: NodeTable,
		Pos:      pos,
		tr:       t,
		Name:     name,
		Repeated: false,
		Fields:   make([]*TableFieldNode, 0),
	}
}

func (t *TableNode) Type() NodeType {
	return t.NodeType
}

func (t *TableNode) Position() Pos {
	return t.Pos
}

func (t *TableNode) String() string {
	return fmt.Sprintf("%s", t.Name)
}

func (t *TableNode) tree() *Tree {
	return t.tr
}

type TableFieldNode struct {
	NodeType
	Pos
	tr  *Tree
	Val string
}

func (t *Tree) newTableField(pos Pos, val string) *TableFieldNode {
	return &TableFieldNode{
		NodeType: NodeTableField,
		Pos:      pos,
		tr:       t,
		Val:      val,
	}
}

func (t *TableFieldNode) Type() NodeType {
	return t.NodeType
}

func (t *TableFieldNode) Position() Pos {
	return t.Pos
}

func (t *TableFieldNode) String() string {
	return fmt.Sprintf("%s", t.Val)
}

func (t *TableFieldNode) tree() *Tree {
	return t.tr
}

type ListNode struct {
	NodeType
	Pos
	tr    *Tree
	Nodes []Node
}

func (t *Tree) newList(pos Pos) *ListNode {
	return &ListNode{
		NodeType: NodeList,
		Pos:      pos,
		tr:       t,
		Nodes:    make([]Node, 0),
	}
}

func (t *ListNode) Type() NodeType {
	return t.NodeType
}

func (t *ListNode) Position() Pos {
	return t.Pos
}

func (t *ListNode) String() string {
	return fmt.Sprintf("%s", t.Nodes)
}

func (t *ListNode) tree() *Tree {
	return t.tr
}
