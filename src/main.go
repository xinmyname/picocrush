package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {

	if len(os.Args) < 4 {
		println("Usage: picocrush -i <filename> -o <filename>")
		println()
		println("  -i Specify source filename.")
		println("  -o Specify destination filename.")
		os.Exit(1)
	}

	var srcPath string
	var dstPath string

	args := os.Args[1:]

	for ; len(args) > 0; args = args[1:] {
		if args[0] == "-i" {
			args = args[1:]
			srcPath = args[0]
		} else if args[0] == "-o" {
			args = args[1:]
			dstPath = args[0]
		}
	}

	srcFile, err := os.Open(srcPath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", srcPath, err)
		os.Exit(1)
	}

	defer srcFile.Close()

	srcImage, _, err := image.Decode(srcFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", srcPath, err)
		os.Exit(1)
	}

	fmt.Printf("Source image size: %v", srcImage.Bounds())
}
