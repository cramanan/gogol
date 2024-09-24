package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

type File struct {
	name string
	body []byte
}

func (f *File) Write(b []byte) (int, error) {
	f.body = append(f.body, b...)
	return len(b), nil
}

type Directory struct {
	Name        string                `json:"name"`
	Directories map[string]*Directory `json:"directories"`
	Files       map[string]*File      `json:"files"`
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name,
		make(map[string]*Directory),
		make(map[string]*File),
	}
}

func NewFile(name string) *File {
	return &File{
		name,
		[]byte{},
	}
}

func (root *Directory) NewDirectory(name string) (d *Directory) {
	d = NewDirectory(name)
	root.Directories[name] = d
	return d
}

func (root Directory) Create(origin string) error {
	var f func(string, Directory) error
	f = func(path string, dir Directory) error {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory: %s", err.Error())
		}
		for name, file := range dir.Files {
			ff, err := os.Create(filepath.Join(path, name))
			if err != nil {
				return fmt.Errorf("error creating file: %s", err.Error())
			}
			_, err = ff.Write(file.body)
			if err != nil {
				return err
			}
			ff.Close()
		}
		for _, subdir := range dir.Directories {
			err = f(filepath.Join(path, subdir.Name), *subdir)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return f(filepath.Join(origin, root.Name), root)
}

func (root *Directory) NewFile(name string) (f *File) {
	f = NewFile(name)
	root.Files[name] = f
	return f
}
