package engine

import (
	"fmt"
	"testing"
)

func TestTraverse(t *testing.T) {
	l := lex(Title, Template)
	fmt.Printf("template: %q\n", Template)
	for {
		i := l.nextItem()
		fmt.Println("item: ", i)
		if i.typ == itemEOF {
			break
		}
	}

	fmt.Println("=====================================")

	fmt.Printf("content: %q\n", Content)
	l = lex(Title, Content)
	for {
		i := l.nextItem()
		fmt.Println("item: ", i)
		if i.typ == itemEOF {
			break
		}
	}

}
