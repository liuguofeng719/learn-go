package main

import (
	"fmt"
	"os"
)

type consoleWriter struct {
}

func (c *consoleWriter) Write(data interface{}) error {
	str := fmt.Sprintf("%v", data)
	_, err := os.Stdout.Write([]byte(str))
	return err
}

func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}
