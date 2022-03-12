package dir

import (
	"github.com/spf13/afero"
	"io/fs"
	"path/filepath"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func Match(file string) bool {
	exts := []string{".mp3", ".aac"}
	ext := filepath.Ext(file)
	return Contains(exts, ext)
}

func Current(appFs afero.Fs, path string) ([]string, error) {
	pattern := path + "/*"
	files, err := afero.Glob(appFs, pattern)
	if err != nil {
		return nil, err
	}
	r := make([]string, 0, len(files))
	for _, file := range files {
		if Match(file) {
			r = append(r, file)
		}
	}
	return r, nil
}

func All(fileSystem fs.FS, path string) ([]string, error) {
	r := make([]string, 0)

	err := fs.WalkDir(fileSystem, path, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// is not dir
		if !info.IsDir() && Match(path) {
			r = append(r, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return r, nil

}
