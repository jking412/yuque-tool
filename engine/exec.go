package engine

import "strings"

type state struct {
	tr   *Tree
	node Node
	text string
	*lexer
	peekCount int
	token     item
	rs        map[string]any
}

func execute(tr *Tree, name, text string) map[string]any {
	s := &state{
		tr:   tr,
		text: text,
		node: tr.Root,
		rs:   make(map[string]any),
	}
	s.lexer = lex(name, s.text)
	s.walk(s.node)
	return s.rs
}

func (s *state) walk(node Node) {
	s.node = node
	switch node.(type) {
	case *TextNode:
		s.walkTextNode(node.(*TextNode))
	case *TableNode:
		s.walkTableNode(node.(*TableNode))
	case *ListNode:
		for _, n := range node.(*ListNode).Nodes {
			s.walk(n)
		}
	}
}

func (s *state) walkTextNode(node *TextNode) {
	i := s.peekItem()
	if i.typ == itemText {
		s.nextItem()
	}
}

func (s *state) walkTableNode(node *TableNode) {
	i := s.nextItem()
	if i.typ != itemTable {
		panic("not table")
	}
	s.nextItem()
	if !node.Repeated {
		fields := make(map[string]any)
		for _, n := range node.Fields {
			i2 := s.nextItem()
			if i2.typ != itemTableField {
				panic("not table field")
			}
			i2 = s.nextItemTrimSpace()
			fields[n.Val] = i2.val
		}
		s.rs[node.Name] = fields
	} else {
		fields := make([]map[string]any, 0)
		for {
			i := s.peekItem()
			if i.typ != itemTableField {
				break
			}
			f := make(map[string]any)
			for _, n := range node.Fields {
				i2 := s.nextItem()
				if i2.typ != itemTableField {
					panic("not table field")
				}
				i2 = s.nextItemTrimSpace()
				f[n.Val] = i2.val
			}
			fields = append(fields, f)
		}
		s.rs[node.Name] = fields
	}
}

func (s *state) peekItem() item {
	if s.peekCount > 0 {
		return s.token
	} else {
		s.token = s.nextItem()
		s.peekCount++
	}
	return s.token
}

func (s *state) nextItem() item {
	if s.peekCount > 0 {
		s.peekCount--
	} else {
		s.token = s.lexer.nextItem()
	}
	return s.token
}

func (s *state) nextItemTrimSpace() item {
	token := s.nextItem()
	token.val = strings.TrimSpace(token.val)
	return token

}