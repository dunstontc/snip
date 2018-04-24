// Package snippet provides snippet models.
package snippet

import (
	"io/ioutil"
	"log"
	"regexp"
)

// ParseFile returns an array of snippets in a file at the given path.
func ParseFile(path string) []string {
	bytez := getBytes(path)
	re := regexp.MustCompile(`(?msU:((^snippet.+$\n)(.+)(?:endsnippet)))`)
	matches := re.FindAllString(string(bytez), -1) // TODO: Find out if FindAll() is faster than FindAllString()
	return matches
}

// Gets a slice of bytes from a file.
func getBytes(file string) []byte {
	theBytes, err := ioutil.ReadFile(file) // TODO: use `func ReadAll(r io.Reader) ([]byte, error)``
	if err != nil {
		log.Fatalf("Demonic Invasion In Progress: %s", err.Error())
		return []byte{}
	}
	return theBytes
}

// File models a file and comes with some handy goodness.
// type File struct {
// 	Path     string
// 	Contents []byte
// }

// NewFile news up a new File.
// func NewFile(path string) *File {
// 	theFile := File{
// 		Path:     path,
// 		Contents: getBytes(path),
// 	}
// 	return &theFile
// }

// // Return a File's contents as a string.
// func (f *File) String() string {
// 	return string(f.Contents)
// }
