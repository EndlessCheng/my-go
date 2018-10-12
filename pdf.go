package main

import (
	"os"
	"gopkg.in/gographics/imagick.v2/imagick"
	"fmt"
)

func init() {
	imagick.Initialize()
}

func main() {
	if len(os.Args) <= 1 {
		os.Exit(1)
	}

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	if err := mw.ReadImage(fmt.Sprintf("%s[0]", os.Args[1])); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(mw.GetImageWidth(), mw.GetImageHeight())
}
