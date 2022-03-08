package dir

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func match(file string) bool {
	exts := []string{".mp3", ".mp4", ".aac", ".wav"}
	ext := filepath.Ext(file)
	return contains(exts, ext)
}

func Current(path string) []string {
	pattern := path + "/*"
	files, err := filepath.Glob(pattern)
	if err != nil {
		panic(err)
	}
	r := make([]string, 0, len(files))
	for _, file := range files {
		if match(file) {
			r = append(r, file)
		}
	}
	return r
}

func All(path string) []string {
	r := make([]string, 0)

	err := filepath.WalkDir(path, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}

		// is not dir
		if !info.IsDir() && match(path) {
			r = append(r, path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("not found")
	}

	return r

}
