package main

import (
	"cruncher/pkg/cruncher"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if isFromPipe() {
		e := crunchData(os.Stdin)
		if e != nil {
			panic(e)
		}
	} else {
		file, e := getFile()
		if e != nil {
			panic(e)
		}
		defer file.Close()
		e = crunchData(file)
		if e != nil {
			panic(e)
		}
	}
}

func isFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode() & os.ModeCharDevice == 0
}

func getFile() (*os.File, error){
	if os.Args[0] == "" {
		return nil, errors.New("please input a file")
	}
	file, e := os.Open(os.Args[1])
	if e != nil {
		return nil, e
	}
	return file, nil
}

func crunchData(r io.Reader) error {
	bytes, err:= ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	result, err := cruncher.Crunch(string(bytes))
	if err != nil {
		return err
	}

	fmt.Printf("%s", result)
	return nil
}