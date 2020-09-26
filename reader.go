package id3

import (
	"io"
	"os"

	v1 "github.com/mikkyang/id3-go/v1"
	v2 "github.com/mikkyang/id3-go/v2"
)

const (
	LatestVersion = 3
)

// Tagger represents the metadata of a tag
type Tagger interface {
	Title() string
	Artist() string
	Album() string
	Year() string
	Genre() string
	Comments() []string
	AllFrames() []v2.Framer
	Frames(string) []v2.Framer
	Frame(string) v2.Framer
	Bytes() []byte
	Dirty() bool
	Padding() uint
	Size() int
	Version() string
}

// File represents the tagged file
type File struct {
	Tagger
	originalSize int
	file         io.ReadSeeker
}

// Parse /s an open file
func Parse(file io.ReadSeeker) (*File, error) {
	res := &File{file: file}

	if v2Tag := v2.ParseTag(file); v2Tag != nil {
		res.Tagger = v2Tag
		res.originalSize = v2Tag.Size()
	} else if v1Tag := v1.ParseTag(file); v1Tag != nil {
		res.Tagger = v1Tag
	} else {
		// Add a new tag if none exists
		res.Tagger = v2.NewTag(LatestVersion)
	}

	return res, nil
}

// Open /s a new tagged file
func Open(name string) (*File, error) {
	fi, err := os.OpenFile(name, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	file, err := Parse(fi)
	if err != nil {
		return nil, err
	}

	return file, nil
}
