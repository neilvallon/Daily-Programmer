package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/NeilVallon/dailyProgrammer/novel-compression/novel"
)

func main() {
	var comp = flag.Bool("c", false, "Compress file")
	var decomp = flag.Bool("d", false, "Decompress file")

	flag.Parse()

	if nf := flag.NFlag(); nf < 1 {
		fmt.Println("Error: No flag set")
		return
	} else if nf > 1 {
		fmt.Println("Error: Too many options")
		return
	}

	// Ignores arguments after first
	if na := flag.NArg(); na < 1 {
		fmt.Println("Error: No input file Specified")
		return
	} else if na != 2 {
		fmt.Println("Error: No output file Specified")
		return
	}

	switch {
	case *comp:
		panic("Compression not implemented")
	case *decomp:
		err := decompress(flag.Arg(0), flag.Arg(1))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func decompress(inf string, outf string) (err error) {
	cd, err := ioutil.ReadFile(inf)
	if err != nil {
		return
	}

	data, err := novel.Decompress(string(cd))
	if err != nil {
		return
	}

	err = ioutil.WriteFile(outf, []byte(data), 0755)
	return
}
