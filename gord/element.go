package gord

import (
	"fmt"
	"io"
)

type BlockElement interface {
}

type TopLevelElement interface {
	dump(dest io.Writer)
	addLine(line string)
}

type Overview struct {
	lines []string
}

func NewOverview() *Overview {
	return &Overview{
		lines: make([]string, 0),
	}
}

func (o *Overview) dump(f io.Writer) {
	fmt.Fprintf(f, "  Overview\n")
}

func (o *Overview) addLine(line string) {
	o.lines = append(o.lines, line)
}

type Headline struct {
	level   uint
	caption string
	lines   []string
}

func NewHeadline(level uint, caption string) *Headline {
	return &Headline{
		level:   level,
		caption: caption,
		lines:   make([]string, 0),
	}
}

func (h *Headline) dump(f io.Writer) {
	fmt.Fprintf(f, "  Headline(%d,%s)\n", h.level, h.caption)
}

func (h *Headline) addLine(line string) {
	h.lines = append(h.lines, line)
}

type Include struct {
	feature string
	lines   []string
}

func NewInclude(feature string) *Include {
	return &Include{
		feature: feature,
		lines:   make([]string, 0),
	}
}

func (i *Include) dump(f io.Writer) {
	fmt.Fprintf(f, "  Include(%s)\n", i.feature)
}

func (i *Include) addLine(line string) {
	i.lines = append(i.lines, line)
}
