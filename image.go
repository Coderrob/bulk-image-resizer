package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

type Image struct {
	From   string
	To     string
	Height int
	Width  int
}

func (i *Image) Create() error {
	defer handleImagePanic(i.From)
	file, err := os.Open(i.From)
	if err != nil {
		log.Fatal(err)
	}
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	_ = file.Close()
	m := resize.Resize(uint(i.Width), 0, img, resize.Lanczos3)
	out, err := os.Create(i.To)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	return jpeg.Encode(out, m, nil)
}

func handleImagePanic(image string) {
	if err := recover(); err != nil {
		fmt.Printf("Failed resizing '%s'. Check folder permission.\n", image)
	}
}
