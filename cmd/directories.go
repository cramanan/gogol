package cmd

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

type Directory struct {
	Name        string
	repo        *git.Repository
	Directories map[string]*Directory
	Files       map[string]*File
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name,
		nil,
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

func (root *Directory) NewDirectory(name string, files ...*File) (d *Directory) {
	d = NewDirectory(name)
	for _, value := range files {
		d.NewFile(value.Name, value.Content)
	}
	root.Directories[name] = d
	return d
}

func (root Directory) Create(origin string) (err error) {
	var f func(string, Directory)
	f = func(path string, dir Directory) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return
		}
		for _, file := range dir.Files {
			ff, err := os.Create(filepath.Join(path, file.Name))
			if err != nil {
				return
			}
			defer ff.Close()
			_, err = ff.Write(file.Content)

			if err != nil {
				return
			}
		}
		for _, subdir := range dir.Directories {
			f(filepath.Join(path, subdir.Name), *subdir)
		}
	}
	f(filepath.Join(origin, root.Name), root)
	return err
}

func (root *Directory) NewFile(name string, content ...[]byte) (f *File) {
	f = &File{
		name,
		[]byte{},
	}
	for _, b := range content {
		f.Content = append(f.Content, b...)
	}
	root.Files[name] = f
	return f
}

type File struct {
	Name    string
	Content []byte
}

func (f *File) WriteString(s string) {
	f.Content = append(f.Content, s...)
}
