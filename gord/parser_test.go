package gord

import (
	"os"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	p := NewParser(strings.NewReader(`
= headline

text
block

== headline2

<<< include.rd

+ headline3

# comment

not comment

`))
	p.Parse()
	p.Dump(os.Stdout)
}
