package novel

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func Decompress(in string) (str string, err error) {
	d := &decompressor{data: in}

	i, err := d.readInt()
	if err != nil {
		return
	}

	d.dictionary = make([]string, i)
	for j := 0; j < i; j++ {
		d.ignoreWhitespace()

		d.dictionary[j], err = d.readWord()
		if err != nil {
			return
		}
	}

	return d.decode()
}

type decompressor struct {
	data  string
	start int
	pos   int

	dictionary []string
}

func (d *decompressor) decode() (str string, err error) {
	sbuff := bytes.NewBufferString("")

	for {
		d.ignoreWhitespace()
		switch b := d.next(); {
		case b == 'E' || b == 'e':
			return sbuff.String(), nil
		case b == 'R' || b == 'r':
			sbuff.WriteByte('\n')
		case b == '-':
			// ignore
		case '0' <= b && b <= '9':
			d.back()
			i, _ := d.readInt()
			word := d.dictionary[i]
			switch d.next() {
			case '^':
				word = strings.Title(word)
			case '!':
				word = strings.ToUpper(word)
			default:
				d.back()
			}
			sbuff.WriteString(word + d.nextSeparator())
		default:
			fmt.Printf("Skipping byte: %q\n", b)
		}
	}

	return
}

func (d *decompressor) nextSeparator() (c string) {
	start, pos := d.start, d.pos

	d.ignoreWhitespace()
	if b := d.next(); b == '-' {
		c = "-"
	} else if '0' <= b && b <= '9' {
		c = " "
	}

	d.start, d.pos = start, pos
	return
}

func (d *decompressor) next() (b byte) {
	b = d.data[d.pos]
	d.pos++
	return
}

func (d *decompressor) back() {
	d.pos--
}

func (d *decompressor) readInt() (i int, err error) {
	for b := d.next(); '0' <= b && b <= '9'; b = d.next() {
	}
	d.back()

	intS := d.data[d.start:d.pos]
	if intS == "" {
		err = fmt.Errorf("Could not read Int at position: %d", d.start)
		return
	}

	i, err = strconv.Atoi(intS)
	if err == nil {
		d.start = d.pos
	}

	return
}

func (d *decompressor) readWord() (str string, err error) {
	for b := d.next(); 'a' <= b && b <= 'z'; b = d.next() {
	}
	d.back()

	str = d.data[d.start:d.pos]
	if str == "" {
		err = fmt.Errorf("Could not read String at position: %d", d.start)
	}
	return
}

func (d *decompressor) ignoreWhitespace() {
	for b := d.next(); b == ' ' || b == '\n' || b == '\t'; b = d.next() {
	}
	d.back()
	d.start = d.pos
	return
}
