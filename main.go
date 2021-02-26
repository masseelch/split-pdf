package main

import (
	"errors"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// Get working directory.
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// List all files.
	fs, err := ioutil.ReadDir(wd)
	if err != nil {
		panic(err)
	}

	// Count the number of pdf-files.
	var c int
	for _, f := range fs {
		if filepath.Ext(f.Name()) == ".pdf" {
			c++
		}
	}

	// There hs to be exactly one pdf file which is then used as split source.
	if c != 1 {
		panic(errors.New("could not determine correct source file"))
	}

	// Get the source file.
	var s os.FileInfo
	for _, f := range fs {
		if filepath.Ext(f.Name()) == ".pdf" {
			s = f
		}
	}

	// Make sure the output dir exists and it is sure its empty.
	o := filepath.Join(wd, "output")
	if err := os.RemoveAll(o); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(o, 0755); err != nil {
		panic(err)
	}

	// Split it.
	if err := api.ExtractPagesFile(s.Name(), o, nil, nil); err != nil {
		panic(err)
	}
}
