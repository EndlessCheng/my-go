package main

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"fmt"
	"image"
)

func main() {
	if len(os.Args) <= 1 {
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(img.Width, img.Height)
}
