package engine

import (
	"fmt"
	"strings"
)

type Tree struct {
	Name      string
	Root      *ListNode
	text      string
	lex       *lexer
	token     item
	peekCount int
}

func NewTree(name, text string) *Tree {
	return &Tree{
		Name: name,
		text: text,
	}
}

func (t *Tree) Parse() error {
	l := lex(t.Name, t.text)
	t.lex = l
	t.parse()
	return nil
}

func (t *Tree) parse() {
	t.Root = t.newList(t.peek().pos)
	for t.peek().typ != itemEOF {
		switch t.peek().typ {
		case itemText:
			textNode := t.newText(t.peek().pos, t.peek().val)
			t.Root.Nodes = append(t.Root.Nodes, textNode)
			t.next()
		case itemTable:
			tableNode := t.newTable(t.peek().pos)
			t.Root.Nodes = append(t.Root.Nodes, tableNode)
			t.next()
			t.parseKeyWord()
			tableNode.Name = t.nextTrimSpace().val
			t.parseTable(tableNode)
		default:
			// 跳过不被处理的 item，如 itemSpace
			t.next()
		}
	}
}

func (t *Tree) parseTable(node *TableNode) {
	for t.peek().typ != itemEOF {
		switch t.peek().typ {
		case itemTableField:
			tableFieldNode := t.newTableField(t.peek().pos)
			t.next()
			tableFieldNode.Val = t.nextTrimSpace().val
			node.Fields = append(node.Fields, tableFieldNode)
		default:
			break
		}
	}
}

func (t *Tree) parseKeyWord() {
	leftDelim := t.peekNonSpace()
	if leftDelim.typ != itemLeftDelim {
		return
	}
	t.nextNonSpace()
	i := t.nextNonSpace()
	rightDelim := t.nextNonSpace()
	if rightDelim.typ != itemRightDelim {
		fmt.Println(rightDelim.typ)
		panic("not right delim")
	}
	switch i.typ {
	case itemRepeat:
		t.Root.Nodes[len(t.Root.Nodes)-1].(*TableNode).Repeated = true
	}
}

func (t *Tree) next() item {
	if t.peekCount > 0 {
		t.peekCount--
	} else {
		t.token = <-t.lex.items
	}
	return t.token
}

func (t *Tree) nextTrimSpace() item {
	token := t.next()
	token.val = strings.TrimSpace(token.val)
	return token
}

func (t *Tree) backup() {
	t.peekCount++
}

func (t *Tree) peek() item {
	if t.peekCount > 0 {
		return t.token
	}
	t.token = <-t.lex.items
	t.peekCount = 1
	return t.token
}

func (t *Tree) nextNonSpace() (token item) {
	for {
		token = t.next()
		if strings.TrimSpace(token.val) != "" {
			break
		}
	}
	return
}

func (t *Tree) peekNonSpace() (token item) {
	token = t.nextNonSpace()
	t.backup()
	return token
}
