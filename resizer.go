package main

import (
	"fmt"
	"path/filepath"
)

func Resizer(manifest *Manifest) []*Image {
	files, _ := GetFileList(manifest.InputDir)
	var images []*Image
	for _, file := range files {
		name := filepath.Base(file)
		img := &Image{
			From: file,
			To:   filepath.Join(manifest.OutputDir, name),
			Height: manifest.Dimensions.Height,
			Width: manifest.Dimensions.Width,
		}
		if err := img.Create(); err != nil {
			fmt.Printf("Failed to create image '%s'. Error: %v\r\n", name, err)
		} else {
			images = append(images, img)
			fmt.Printf("Image '%s' created\r\n", name)
		}
	}
	return images
}
