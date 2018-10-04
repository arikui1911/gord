package gord

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Parser struct {
	src       *bufio.Scanner
	line      uint
	toplevels []TopLevelElement
}

func NewParser(src io.Reader) *Parser {
	return &Parser{
		src:       bufio.NewScanner(src),
		line:      1,
		toplevels: make([]TopLevelElement, 0),
	}
}

func (p *Parser) Parse() error {
	var current TopLevelElement

	current = NewOverview()
	p.toplevels = append(p.toplevels, current)

	for p.src.Scan() {
		if isComment(p.src.Text()) {
			continue
		}
		if h := tryParseHeadline(p.src.Text()); h != nil {
			current = h
			p.toplevels = append(p.toplevels, current)
			continue
		}
		if i := tryParseInclude(p.src.Text()); i != nil {
			current = i
			p.toplevels = append(p.toplevels, current)
			continue
		}
		current.addLine(p.src.Text())
	}

	return p.src.Err()
}

func (p *Parser) Dump(dest io.Writer) {
	fmt.Fprintf(dest, "Document\n")
	for _, e := range p.toplevels {
		e.dump(dest)
	}
}

func isComment(line string) bool {
	return strings.HasPrefix(line, "#")
}

func tryParseHeadline(line string) *Headline {
	n := countSameRunePrefix(line, '=')
	if n > 0 && n < 5 {
		return NewHeadline(n, line[n:])
	}
	n = countSameRunePrefix(line, '+')
	if n > 0 && n < 3 {
		return NewHeadline(n+4, line[n:])
	}
	return nil
}

func tryParseInclude(line string) *Include {
	n := countSameRunePrefix(line, '<')
	if n != 3 {
		return nil
	}
	return NewInclude(strings.TrimSpace(line[n:]))
}

func countSameRunePrefix(str string, prefixRune rune) uint {
	n := 0
	for i, r := range str {
		n = i
		if r != prefixRune {
			break
		}
	}
	return uint(n)
}
