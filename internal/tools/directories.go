package tools

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type File struct {
	Name    string `yaml:"name"`
	Content []byte `yaml:"content"`
}

// File.Write() write bytes onto the File.Content
func (f *File) Write(b []byte) {
	f.Content = append(f.Content, b...)
}

// File.WriteString() converts the string input into bytes then write with File.Write()
func (f *File) WriteString(s string) {
	f.Write([]byte(s))
}

type Directory struct {
	Name        string       `yaml:"name"`
	Directories []*Directory `yaml:"directories"`
	Files       []*File      `yaml:"files"`
}

func RetrieveYAMLdir(url string) (dir *Directory, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if err = yaml.NewDecoder(resp.Body).Decode(&dir); err != nil {
		return
	}
	return
}

// Directory.Create() opens every directory and files. It will also write the Files.Content into the files.
func (root Directory) Create(origin string) (err error) {
	var f func(string, Directory)
	f = func(path string, dir Directory) {
		if err = os.MkdirAll(filepath.Join(origin, path), os.ModePerm); err != nil {
			return
		}
		for _, file := range dir.Files {
			ff, err := os.Create(filepath.Join(origin, path, file.Name))
			if err != nil {
				return
			}
			_, err = ff.Write(file.Content)
			if err != nil {
				return
			}
		}
		for _, subdir := range dir.Directories {
			newPath := filepath.Join(path, subdir.Name)
			if err != nil {
				return
			}
			f(newPath, *subdir)
		}

	}
	f(root.Name, root)
	return err
}

func (root Directory) String() string {
	var sb strings.Builder
	var f func(string, Directory)
	f = func(path string, dir Directory) {
		sb.WriteString(filepath.Join(path, "\n"))
		for _, file := range dir.Files {
			sb.WriteString(filepath.Join("  "+path, file.Name) + "\n")
		}
		for _, subdir := range dir.Directories {
			newPath := filepath.Join(path, subdir.Name)
			f(newPath, *subdir)
		}
	}
	f(root.Name, root)
	return sb.String()
}

func (root *Directory) AddFile(f File) {
	root.Files = append(root.Files, &f)
}

func (root *Directory) Mkdir(d Directory) {
	root.Directories = append(root.Directories, &d)
}

func (root *Directory) Search(filename string) (fptr *File) {
	var f func(string, *Directory)
	f = func(path string, dir *Directory) {
		for _, file := range dir.Files {
			if filepath.Join(path, file.Name) == filename {
				fptr = file
				return
			}
		}
		for _, subdir := range dir.Directories {
			newPath := filepath.Join(path, subdir.Name)
			f(newPath, subdir)
		}
	}
	f(root.Name, root)
	return fptr
}

func (root *Directory) PopFile(filename string) (fptr *File) {
	var f func(string, *Directory)
	f = func(path string, dir *Directory) {
		for index, file := range dir.Files {
			if filepath.Join(path, file.Name) == filename {
				fptr = file
				dir.Files = append(dir.Files[:index], dir.Files[index+1:]...)
				return
			}
		}
		for _, subdir := range dir.Directories {
			newPath := filepath.Join(path, subdir.Name)
			f(newPath, subdir)
		}
	}
	f(root.Name, root)
	return fptr
}
