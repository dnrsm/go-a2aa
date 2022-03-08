package id3

import (
	"github.com/bogem/id3v2"
	"go-a2aa/pkg/dir"
	"log"
)

func SetTags(path string) {
	files := dir.Current(path)
	for _, file := range files {
		setTag(file)
	}
}

func SetTagsAll(path string) {
	files := dir.All(path)
	for _, file := range files {
		setTag(file)
	}
}

func setTag(file string) {
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
	}
}
