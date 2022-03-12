package id3

import (
	"a2aa/pkg/dir"
	"github.com/bogem/id3v2"
	"github.com/spf13/afero"
	"io/fs"
	"log"
)

func SetTags(appFs afero.Fs, path string) error {
	files, err := dir.Current(appFs, path)
	if err != nil {
		return err
	}
	for _, file := range files {
		if err := setTag(file); err != nil {
			return err
		}
	}
	return nil
}

func SetTagsAll(dirFS fs.FS, path string) error {
	files, err := dir.All(dirFS, path)
	if err != nil {
		return err
	}
	for _, file := range files {
		if err := setTag(file); err != nil {
			return err
		}
	}
	return nil
}

func setTag(file string) error {
	tag, err := id3v2.Open(file, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("Error while opening mp3 file: ", err)
	}
	defer tag.Close()

	// Read tag
	artist := tag.Artist()
	// Set tag
	tag.SetAlbumArtist(artist)

	// Write tag to file.mp3.
	if err = tag.Save(); err != nil {
		log.Fatal("Error while saving a tag: ", err)
		return err
	}
	return nil
}
