package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

var RootDirectory = NewDirectory("untitled")

type File struct {
	body []byte
}

func (f *File) Write(b []byte) (int, error) {
	f.body = append(f.body, b...)
	return len(b), nil
}

type Directory struct {
	Name        string
	Directories map[string]*Directory
	Files       map[string]*File
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
		[]byte{},
	}
}

func (root *Directory) NewDirectory(name string) (d *Directory) {
	d = NewDirectory(name)
	root.Directories[name] = d
	return d
}

func (root *Directory) Create(origin string) (err error) {
	var f func(string, *Directory) error
	f = func(path string, dir *Directory) error {
		os.Mkdir(path, os.ModePerm) // TODO: Should warn
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
			err = f(filepath.Join(path, subdir.Name), subdir)
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
