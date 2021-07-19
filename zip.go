package gek

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Unzip unzip("/tmp/report-2015.zip", "/tmp/reports/")
func Unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(path, file.Mode())
			if err != nil {
				return err
			}
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer func(fileReader io.ReadCloser) {
			err := fileReader.Close()
			if err != nil {
				return
			}
		}(fileReader)

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer func(targetFile *os.File) {
			err := targetFile.Close()
			if err != nil {
				return
			}
		}(targetFile)

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

// Zipit zipit("/tmp/documents", "/tmp/backup.zip") or zipit("/tmp/report.txt", "/tmp/report-2015.zip")
func Zipit(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer func(zipfile *os.File) {
		err := zipfile.Close()
		if err != nil {
			return
		}
	}(zipfile)

	archive := zip.NewWriter(zipfile)
	defer func(archive *zip.Writer) {
		err := archive.Close()
		if err != nil {
			return
		}
	}(archive)

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				return
			}
		}(file)
		_, err = io.Copy(writer, file)
		return err
	})
	if err != nil {
		return err
	}

	return err
}
