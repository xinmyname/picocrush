package main

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) < 4 {
		println("Usage: picocrush -i <filename> -o <filename>")
		println()
		println("  -i Specify source filename.")
		println("  -o Specify destination filename.")

		println("  -o Specify destination filename. Output can be .png or .bin")
		println()
		println("     .png generates an indexed PNG file")
		println("     .bin generates a raw binary file, 16bpp RGB 565")

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

	dstImage := image.NewRGBA(srcImage.Bounds())

	width := srcImage.Bounds().Dx()
	height := srcImage.Bounds().Dy()

	for y := 0; y < height; y += 1 {
		for x := 0; x < width; x += 1 {
			r, g, b, _ := srcImage.At(x, y).RGBA()
			cc := color.RGBA{uint8((r >> 11) << 3), uint8((g >> 10) << 2), uint8((b >> 11) << 3), 0xff}
			dstImage.SetRGBA(x, y, cc)
		}
	}

	dstFile, err := os.Create(dstPath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", dstPath, err)
		os.Exit(1)
	}

	defer dstFile.Close()

	dstExt := filepath.Ext(dstPath)

	if dstExt == ".png" {
		png.Encode(dstFile, dstImage)
	} else if dstExt == ".bin" {

		for y := 0; y < height; y += 1 {
			for x := 0; x < width; x += 1 {
				r, g, b, _ := dstImage.At(x, y).RGBA()
				c := uint16(((r >> 11) << 11) | ((g >> 10) << 5) | (b >> 11))
				binary.Write(dstFile, binary.LittleEndian, c)
			}
		}
	}
}
