package engine

import (
	"fmt"
	"testing"
)

const Title = "Test"
const Template = "<a name=\"mQEDU\"></a>\n# Table1\n- name\n- age\n- job\n- home\n\n"
const RepeatTemplate = "<a name=\"mQEDU\"></a>\n# {{*}}Table1\n- name\n- age\n- job\n- home\n\n"
const Content = "<a name=\"mQEDU\"></a>\n# Table1\n- 张三\n- 18\n- 程序员\n- 杭州\n"
const RepeatContent = "<a name=\"mQEDU\"></a>\n# Table1\n- 张三\n- 18\n- 程序员\n- 杭州\n\n- 李四\n- 30\n- 教师\n- 台州\n"

func TestParseTemplate(t *testing.T) {
	templateTree := NewTree(Title, Template)
	templateTree.Parse()
	rs := execute(templateTree, Title, Content)
	fmt.Println(rs)
}

func TestParseRepeatTemplate(t *testing.T) {
	templateTree := NewTree(Title, RepeatTemplate)
	templateTree.Parse()
	rs := execute(templateTree, Title, RepeatContent)
	fmt.Println(rs)
}
