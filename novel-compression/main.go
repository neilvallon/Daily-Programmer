package main

import (
	"flag"
	"fmt"
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
		panic("Decompression not implemented")
	}
}
