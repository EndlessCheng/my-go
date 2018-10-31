package main

import (
	"os"
	"fmt"
	"image/jpeg"
	"image/png"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) <= 1 {
		os.Exit(1)
	}

	fileName := os.Args[1]
	fileNameWithExt := fileName
	ext := filepath.Ext(fileName)
	if ext == "" {
		fileNameWithExt += ".png"
	}

	pngFile, err := os.Open(fileNameWithExt)
	checkError(err)
	defer pngFile.Close()

	pngSrc, err := png.Decode(pngFile)
	checkError(err)

	jpgFile, err := os.Create(fileName + ".jpg")
	checkError(err)
	defer jpgFile.Close()

	err = jpeg.Encode(jpgFile, pngSrc, &jpeg.Options{Quality: 100})
	checkError(err)
}
