package main

import (
	"errors"
	"fmt"
	"os"
)

type FileWriter struct {
	file *os.File
}

func (f *FileWriter) SetFile(fileName string) (err error) {
	if f.file != nil {
		f.file.Close()
	}
	f.file, err = os.Create(fileName)
	return err
}

func (f *FileWriter) Write(data interface{}) error {
	if f.file == nil {
		return errors.New("file not created")
	}
	str := fmt.Sprintf("%v\n", data)
	_, err := f.file.Write([]byte(str))
	return err
}

func newFileWriter() *FileWriter {
	return &FileWriter{}
}
