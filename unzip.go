package zip_files

import (
	"archive/zip"
	"bytes"
)

// Unzip unzips a zip archive and returns a map of the file names to their contents
func Unzip(zipBytes []byte) (map[string][]byte, error) {
	files := make(map[string][]byte)

	zipReader, err := zip.NewReader(bytes.NewReader(zipBytes), int64(len(zipBytes)))
	if err != nil {
		return nil, err
	}

	// Read each file from the zip
	for _, file := range zipReader.File {
		// Open the file inside the zip
		rc, err := file.Open()
		if err != nil {
			return nil, err
		}

		// Read the file contents
		content := new(bytes.Buffer)
		_, err = content.ReadFrom(rc)
		rc.Close()
		if err != nil {
			return nil, err
		}

		// Store the file content
		files[file.Name] = content.Bytes()
	}

	return files, nil
}
