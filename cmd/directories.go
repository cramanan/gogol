package cmd

import (
	"os"
	"path/filepath"
)

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

func (root *Directory) NewDirectory(name string, files ...*File) (d *Directory) {
	d = NewDirectory(name)
	for _, value := range files {
		d.NewFile(value.name, value.content)
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
			ff, err := os.Create(filepath.Join(path, file.name))
			if err != nil {
				return
			}
			defer ff.Close()
			_, err = ff.Write(file.content)

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

func (root *Directory) Read(origin string) error {
	var f func(string, *Directory) error
	f = func(path string, dir *Directory) error {
		entries, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			if entry.IsDir() {
				if entry.Name() == ".git" {
					continue
				}
				subdir := dir.NewDirectory(entry.Name())
				err = f(filepath.Join(path, entry.Name()), subdir)
				if err != nil {
					return err
				}
			} else {
				file := dir.NewFile(entry.Name())
				file.content, err = os.ReadFile(filepath.Join(path, entry.Name()))
				if err != nil {
					return err
				}
			}
		}

		return nil
	}
	return f(origin, root)
}

func (root *Directory) NewFile(name string, content ...[]byte) (f *File) {
	f = &File{
		name,
		[]byte{},
	}
	for _, b := range content {
		f.content = append(f.content, b...)
	}
	root.Files[name] = f
	return f
}

type File struct {
	name    string
	content []byte
}

func (f *File) Write(b []byte) {
	f.content = append(f.content, b...)
}

func (f *File) WriteString(s string) {
	f.content = append(f.content, s...)
}

func (f File) Create(origin string) error {
	osf, err := os.Create(filepath.Join(origin, f.name))
	if err != nil {
		return err
	}
	defer osf.Close()

	_, err = osf.Write(f.content)
	return err
}
