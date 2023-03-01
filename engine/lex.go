package engine

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type item struct {
	typ  itemType
	pos  Pos
	val  string
	line int
}

func (i item) String() string {
	switch {
	case i.typ == itemEOF:
		return "EOF"
	case i.typ == itemErr:
		return i.val
	case len(i.val) > 10:
		// %q打印字面值，比如\n会被打印出来，而不是换行效果
		return fmt.Sprintf("%.10q...", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}

type itemType int
type Pos int

const (
	itemErr itemType = iota
	itemEOF
	itemSpace
	itemText
	itemTable
	itemTableField
	itemRepeat
	itemLeftDelim
	itemRightDelim
)

type stateFn func(*lexer) stateFn

const eof = -1
const spaceChars = " \t\r\n"

type lexer struct {
	name      string
	input     string
	start     Pos
	pos       Pos
	atEOF     bool
	items     chan item
	line      int
	startLine int
}

const (
	TableKeyword      = '#'
	TableFieldKeyword = '-'
	repeatKeyword     = '*'
	LeftDelim         = "{{"
	RightDelim        = "}}"
)

const keyWord = string(TableKeyword) + string(TableFieldKeyword) + string(repeatKeyword) +
	string(LeftDelim) + string(RightDelim)

func lex(name, input string) *lexer {
	l := &lexer{
		name:      name,
		input:     input,
		items:     make(chan item),
		line:      1,
		startLine: 1,
	}
	go l.run()
	return l
}

func (l *lexer) run() {
	for s := lexText; s != nil; {
		s = s(l)
	}
	close(l.items)
}

func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.atEOF = true
		return eof
	}
	// 返回输入字符串中的第一个utf8编码的rune字符
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += Pos(w)
	if r == '\n' {
		l.line++
	}
	return r
}

func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *lexer) backup() {
	if !l.atEOF && l.pos > 0 {
		r, w := utf8.DecodeLastRuneInString(l.input[:l.pos])
		l.pos -= Pos(w)
		// Correct newline count.
		if r == '\n' {
			l.line--
		}
	}
}

func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.start, l.input[l.start:l.pos], l.startLine}
	l.start = l.pos
	l.startLine = l.line
}

func (l *lexer) nextItem() item {
	return <-l.items
}

// 获取下一个Item并删除Item两侧的itemSpace
func (l *lexer) nextItemTrimSpace() item {
	i := <-l.items
	i.val = strings.TrimSpace(i.val)
	return i
}

func (l *lexer) errorf(format string, args ...any) stateFn {
	l.items <- item{itemErr, l.start, fmt.Sprintf(format, args...), l.line}
	return nil
}

func (l *lexer) ignore() {
	l.line += strings.Count(l.input[l.start:l.pos], "\n")
	l.start = l.pos
	l.startLine = l.line
}

func lexText(l *lexer) stateFn {
	for {
		switch r := l.peek(); {
		case r == eof:
			if l.pos > l.start {
				l.emit(itemText)
			}
			l.ignore()
			l.emit(itemEOF)
			return nil
		case strings.ContainsRune(keyWord, r):
			if l.pos > l.start {
				l.emit(itemText)
			}
			l.ignore()
			return lexKeyword
		default:
			l.next()
		}
	}
}

func lexKeyword(l *lexer) stateFn {
	switch r := l.next(); {
	case r == TableKeyword:
		l.emit(itemTable)
		return lexText
	case r == TableFieldKeyword:
		l.emit(itemTableField)
		return lexText
	case r == repeatKeyword:
		l.emit(itemRepeat)
		return lexText
	case r == rune(LeftDelim[0]):
		if l.peek() == rune(LeftDelim[1]) {
			l.next()
			l.emit(itemLeftDelim)
		}
		return lexText
	case r == rune(RightDelim[0]):
		if l.peek() == rune(RightDelim[1]) {
			l.next()
			l.emit(itemRightDelim)
		}
		return lexText
	}
	return l.errorf("unterminated keyword")
}
