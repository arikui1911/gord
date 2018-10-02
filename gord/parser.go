package gord

import (
	"bufio"
	"io"
)

type Parser struct {
	src    *bufio.Scanner
	lineno uint
}

func NewParser(src io.Reader) *Parser {
	return &Parser{
		src:    bufio.NewScanner(src),
		lineno: 1,
	}
}

func (p *Parser) Parse() error {
	for p.src.Scan() {
		p.src.Text()
	}
	if err := p.src.Err(); err != nil {
		return err
	}

	return nil
}
