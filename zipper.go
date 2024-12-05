package zip_files

import (
	"archive/zip"
	"bytes"
)

// New creates a new Zipper
func New() *Zipper {
	return &Zipper{
		files: make(map[string][]byte),
	}
}

// Zipper is a struct that can add files to a zip archive and zip them
type Zipper struct {
	files map[string][]byte
}

// AddFile adds a file to the zip archive
func (z *Zipper) AddFile(filename string, file []byte) {
	z.files[filename] = file
}

// CountFiles returns the number of files in the zip archive
func (z *Zipper) CountFiles() int {
	return len(z.files)
}

// Files returns the names of the files in the zip archive
func (z *Zipper) Files() []string {
	files := make([]string, 0)
	for file := range z.files {
		files = append(files, file)
	}
	return files
}

// Zip returns the zip archive as a byte slice
func (z *Zipper) Zip() ([]byte, error) {
	buf := new(bytes.Buffer)

	zipWriter := zip.NewWriter(buf)

	for filename, file := range z.files {
		zipFile, err := zipWriter.Create(filename)
		if err != nil {
			return nil, err
		}

		_, err = zipFile.Write(file)
		if err != nil {
			return nil, err
		}
	}

	err := zipWriter.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
