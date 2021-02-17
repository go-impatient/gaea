package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var sources []string

type File struct {
	path string
}

// NewFile new a file source.
func NewFile(path string) *File {
	return &File{path: path}
}

func (f *File)loadFile(path string) (string, error)  {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return "", err
	}
	return info.Name(), nil
}

func (f *File) loadDir(path string) (sources []string,  err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		// ignore hidden files
		if file.IsDir() || strings.HasPrefix(file.Name(), ".") {
			continue
		}

		p := filepath.Join(path, file.Name())
		_, err = f.loadFile(p)
		if err != nil {
			return nil, err
		}
		sources = append(sources, p)
	}
	return
}

func (f *File) Load() (sources []string,  err error){
	fi, err := os.Stat(f.path)
	if err != nil {
		return nil, err
	}
	if fi.IsDir() {
		return f.loadDir(f.path)
	}

	_, err = f.loadFile(f.path)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("根路径: %s\n", f.path)
	return []string{f.path}, nil
}