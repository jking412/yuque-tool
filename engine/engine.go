package engine

func Parse(name, content, template string) map[string]any {
	templateTree := NewTree(name, template)
	templateTree.Parse()
	rs := execute(templateTree, name, content)
	return rs
}
