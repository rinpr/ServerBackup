package archive

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func EzCompress(fileName string) {
	start := time.Now()

	if err := zipSource(fileName, fileName+".rar"); err != nil {
		log.Fatal(err)
	}

	end := time.Now()
	duration := end.Sub(start)

	fmt.Printf("Function took %v milliseconds to execute\n", duration.Milliseconds())
}

func zipSource(source, target string) error {
	// 1. Create a ZIP file and zip.Writer
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	writer := zip.NewWriter(f)
	defer func(writer *zip.Writer) {
		err := writer.Close()
		if err != nil {

		}
	}(writer)

	// 2. Go through all the files of the source
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 3. Create a local file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// set compression
		header.Method = zip.Deflate

		// 4. Set relative path of a file as the header name
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		// 5. Create writer for the file header and save content of the file
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(f)

		_, err = io.Copy(headerWriter, f)
		return err
	})
}
