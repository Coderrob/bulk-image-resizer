package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFileList(path string) (fileNames []string, err error) {
	defer handleFilePanic(path)
	folder, err := os.Open(path)
	if err != nil {
		return fileNames, err
	}
	files, err := folder.Readdir(-1)
	defer folder.Close()
	if err != nil {
		return fileNames, err
	}
	fileNames = make([]string, 0)
	for _, file := range files {
		name := file.Name()
		filepath.Ext(name)
		if filepath.Ext(name) == ".png" {
			fileNames = append(fileNames, filepath.Join(path, name))
		}
	}
	return fileNames, nil
}

func handleFilePanic(path string) {
	if err := recover(); err != nil {
		fmt.Printf("Failed reading folder '%s'. Check folder permission.\n", path)
	}
}
