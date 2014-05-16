package novel

import (
	"fmt"
	"strconv"
)

func Decompress(in string) (str string, err error) {
	d := &decompressor{data: in}

	_, err = d.readInt()
	if err != nil {
		return
	}

	return
}

type decompressor struct {
	data  string
	start int
	pos   int
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
