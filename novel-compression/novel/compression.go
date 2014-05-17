package novel

import (
	"bytes"
	"fmt"
	"strconv"
)

func Compress(in string) (str string, err error) {
	c := &compressor{
		data: in,
		dict: make(map[string]int),
	}

	return c.encode()
}

type compressor struct {
	dict     map[string]int
	dictSize int

	data  string
	start int
	pos   int
}

func (c *compressor) encode() (str string, err error) {
	archbuff := bytes.NewBufferString("")

	for {
		c.ignoreWhitespace()
		w, err := c.readWord()
		if err != nil {
			return archbuff.String(), nil
		}
		i := c.addWord(w)
		archbuff.WriteString(strconv.Itoa(i) + " ")
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
	for b := c.next(); b == ' ' || b == '\n' || b == '\t'; b = c.next() {
	}
	c.back()
	c.start = c.pos
	return
}

func (c *compressor) readWord() (str string, err error) {
	for b := c.next(); 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z'; b = c.next() {
	}
	c.back()

	str = c.data[c.start:c.pos]
	if str == "" {
		err = fmt.Errorf("Could not read String at position: %d", c.start)
	}
	return
}
