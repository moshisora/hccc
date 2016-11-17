package hccc

import (
	"log"
	"os"
	"path/filepath"
)

func generatePaths(str string) []string {
	var plist []string

	if b, _ := isDirectory(str); !b {
		plist = append(plist, str)
		return plist
	}

	err := filepath.Walk(str, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		plist = append(plist, path)

		return nil
	})

	if err != nil {
		log.Fatal("[error] file path walking", err)
	}

	return plist
}

func isDirectory(name string) (isDir bool, err error) {
	fInfo, err := os.Stat(name)
	if err != nil {
		return false, err
	}
	return fInfo.IsDir(), nil
}
