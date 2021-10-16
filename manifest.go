package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Dimensions struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type Manifest struct {
	Dimensions Dimensions `json:"dimensions"`
	InputDir   string     `json:"inputDir"`
	OutputDir  string     `json:"outputDir"`
}

func OpenManifest(name string) (*Manifest, error) {
	defer handleManifestPanic(name)
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var manifest Manifest
	err = json.Unmarshal(bytes, &manifest)
	if err != nil {
		return nil, err
	}
	return &manifest, nil
}

func handleManifestPanic(path string) {
	if err := recover(); err != nil {
		fmt.Printf("Failed reading manifest '%s'. Check folder permission.\n", path)
	}
}
