package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	manifest, _ := OpenManifest("./manifest.json")
	images := Resizer(manifest)
	CreateCSV(manifest.OutputDir, images)
}
