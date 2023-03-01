package engine

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	tr := NewTree(Title, Template)
	tr.Parse()
	fmt.Println(tr.Root)

	tr1 := NewTree(Title, Content)
	tr1.Parse()
	fmt.Println(tr1.Root)
}
