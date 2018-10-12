package main

import (
	"os"
	"gopkg.in/gographics/imagick.v2/imagick"
	"fmt"
	"strconv"
)

func init() {
	imagick.Initialize()
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "<pdf-path>")
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "<pdf-path>", "<page-index>")
		os.Exit(1)
	}

	pdfPath := os.Args[1]

	index := 1
	var err error
	if len(os.Args) >= 3 {
		index, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		if index < 1 {
			fmt.Fprintln(os.Stderr, "index must > 0")
			os.Exit(1)
		}
	}

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	if err := mw.ReadImage(fmt.Sprintf("%s[%d]", pdfPath, index-1)); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(mw.GetImageWidth(), mw.GetImageHeight())
}
