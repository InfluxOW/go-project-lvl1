package printer

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

const (
	lineLn = 120
)

var (
	infoColor    = color.New(color.FgHiWhite).Add(color.BgHiBlue)
	successColor = color.New(color.FgHiWhite).Add(color.BgHiGreen)
	failureColor = color.New(color.FgHiWhite).Add(color.BgHiRed)
	h1color      = color.New(color.FgHiBlue).Add(color.BgHiWhite).Add(color.Bold)
	h2color      = color.New(color.FgHiWhite).Add(color.BgHiMagenta)
	h3color      = color.New(color.FgHiWhite).Add(color.BgHiBlue)
)

func copyColor(color *color.Color) *color.Color {
	cp := *color

	return &cp
}

func PrintH1(h string) {
	newline(h1color.PrintFunc(), h).newlined().filled().headered().centered().print()
}

func PrintH2(h string, prefix string) {
	basePrinter := h2color.PrintFunc()
	prefixPrinter := copyColor(h2color).Add(color.Bold).PrintFunc()

	newline(basePrinter, h).prefixed(newline(prefixPrinter, fmt.Sprintf(" %s ", prefix))).newlined().filled().headered().print()
}

func PrintH3(h string, prefix string) {
	basePrinter := h3color.PrintFunc()
	prefixPrinter := copyColor(h3color).Add(color.Bold).PrintFunc()

	newline(basePrinter, h).prefixed(newline(prefixPrinter, fmt.Sprintf(" %s ", prefix))).newlined().filled().print()
}

func PrintInfo(msg string) {
	printMessage(infoColor.PrintFunc(), msg, false)
}

func PrintlnInfo(msg string) {
	printMessage(infoColor.PrintFunc(), msg, true)
}

func PrintSuccess(msg string) {
	printMessage(successColor.PrintFunc(), msg, false)
}

func PrintlnSuccess(msg string) {
	printMessage(successColor.PrintFunc(), msg, true)
}

func PrintFailure(msg string) {
	printMessage(failureColor.PrintFunc(), msg, false)
}

func PrintlnFailure(msg string) {
	printMessage(failureColor.PrintFunc(), msg, true)
}

func printMessage(printer func(a ...interface{}), msg string, fill bool) {
	line := newline(printer, fmt.Sprintf(" %s ", msg))

	if fill {
		line.filled()
	}

	line.newlined().print()
}

type line struct {
	printer func(a ...interface{})
	str     string
	prefix  *line
	fill    bool
	center  bool
	newline bool
	header  bool
}

func newline(printer func(a ...interface{}), str string) *line {
	return &line{
		printer: printer,
		str:     str,
		prefix:  nil,
		fill:    false,
		center:  false,
		newline: false,
		header:  false,
	}
}

func emptyline(printer func(a ...interface{})) {
	newline(printer, "").centered().newlined().filled().print()
}

func (l *line) prefixed(prefix *line) *line {
	l.prefix = prefix

	return l
}

func (l *line) filled() *line {
	l.fill = true

	return l
}

func (l *line) centered() *line {
	l.center = true

	return l
}

func (l *line) newlined() *line {
	l.newline = true

	return l
}

func (l *line) headered() *line {
	l.header = true

	return l
}

func (l *line) pad() *line {
	strLn := len(l.str)
	if l.fill && strLn < lineLn {
		remainedChars := lineLn - strLn
		if l.prefix != nil {
			remainedChars -= len(l.prefix.str)
		}

		if l.center {
			leftSpaces := remainedChars / 2
			rightSpaces := remainedChars - leftSpaces
			l.str = strings.Repeat(" ", leftSpaces) + l.str + strings.Repeat(" ", rightSpaces)
		} else {
			l.str = l.str + strings.Repeat(" ", remainedChars)
		}
	}

	return l
}

func (l *line) print() {
	l.pad()

	if l.header {
		emptyline(l.printer)
	}

	if l.prefix != nil {
		l.prefix.print()
	}

	l.printer(l.str)

	if l.header {
		fmt.Println()
		emptyline(l.printer)
	}

	if l.newline {
		fmt.Println()
	}
}
