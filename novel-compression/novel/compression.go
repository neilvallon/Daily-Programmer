package novel

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type dictionary map[string]int

func (d dictionary) String() string {
	sarr := make([]string, len(d))
	for k, v := range d {
		sarr[v] = k
	}
	return strings.Join(sarr, " ")
}

func Compress(in string) (str string, err error) {
	c := &compressor{
		data: in,
		dict: make(dictionary),
	}

	return c.encode()
}

type compressor struct {
	dict     dictionary
	dictSize int

	data  string
	start int
	pos   int
}

const (
	LOWER = iota
	FIRST
	UPPER
)

func (c *compressor) String() string {
	return strconv.Itoa(c.dictSize) + c.dict.String()
}

func (c *compressor) encode() (str string, err error) {
	archbuff := bytes.NewBufferString("")

	lastword := false
	for {
		c.ignoreWhitespace()
		switch b := c.next(); b {
		case '\n':
			archbuff.WriteByte('R')
			lastword = false
		case '!':
			archbuff.WriteByte(' ')
			archbuff.WriteByte(b)
			lastword = false
		case '.', ',', '?', ';', ':', '-':
			archbuff.WriteByte(b)
			lastword = false
		case '\x00':
			return c.String() + archbuff.String() + "E", nil
		default:
			c.back()
			w, lc, err := c.readWord()
			if err != nil {
				return "", err
			}

			if lastword {
				archbuff.WriteByte(' ')
			}

			i := c.addWord(strings.ToLower(w))
			archbuff.WriteString(strconv.Itoa(i))

			lastword = false
			switch lc {
			case FIRST:
				archbuff.WriteByte('^')
			case UPPER:
				archbuff.WriteByte('!')
			default:
				lastword = true
			}
		}
	}
}

func (c *compressor) addWord(w string) (i int) {
	i, ok := c.dict[w]
	if !ok {
		i = c.dictSize
		c.dict[w] = i
		c.dictSize++
	}
	return
}

func (c *compressor) next() (b byte) {
	if c.pos < len(c.data) {
		b = c.data[c.pos]
	}

	c.pos++
	return
}

func (c *compressor) back() {
	c.pos--
}

func (c *compressor) ignoreWhitespace() {
	for b := c.next(); b == ' ' || b == '\t'; b = c.next() {
	}
	c.back()
	c.start = c.pos
	return
}

func (c *compressor) readWord() (str string, lcase int, err error) {
	if b := c.next(); 'A' <= b && b <= 'Z' {
		lcase = UPPER
	}

LOOP:
	for {
		switch b := c.next(); {
		case 'A' <= b && b <= 'Z':
		case 'a' <= b && b <= 'z':
			if lcase == UPPER {
				lcase = FIRST
			}
		default:
			c.back()
			break LOOP
		}
	}

	str = c.data[c.start:c.pos]
	if str == "" {
		err = fmt.Errorf("Could not read String at position: %d", c.start)
	}
	return
}
