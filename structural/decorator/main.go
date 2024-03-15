package main

import (
	"fmt"
)

type Reader interface {
	Read() string
}

type FileReader struct{}

func (fr *FileReader) Read() string {
	return "data from file"
}

type BufferedReader struct {
	reader Reader
}

func (br *BufferedReader) Read() string {
	data := br.reader.Read()
	return "buffered " + data
}

type DecryptReader struct {
	reader Reader
}

func (dr *DecryptReader) Read() string {
	data := dr.reader.Read()
	return "decrypted " + data
}

func main() {
	fileReader := &FileReader{}

	bufferedReader := &BufferedReader{reader: fileReader}
	fmt.Println(bufferedReader.Read())

	decryptBufferedReader := &DecryptReader{reader: bufferedReader}
	fmt.Println(decryptBufferedReader.Read())
}