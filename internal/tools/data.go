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
	Path    string
	Content []byte `yaml:"content"`
}

type Directory struct {
	Name        string       `yaml:"name"`
	Directories []*Directory `yaml:"directories"`
	Files       []*File      `yaml:"files"`
}

func RetrieveYAMLdir(url string) (dir Directory, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if err = yaml.NewDecoder(resp.Body).Decode(&dir); err != nil {
		return
	}
	return
}

func CreateDirAndFiles(dir Directory) (err error) {
	if err = os.Mkdir(dir.Name, os.ModePerm); err != nil {
		return
	}
	if err = os.Chdir(dir.Name); err != nil {
		return
	}

	for _, file := range dir.Files {
		f, err := os.Create(file.Name)
		if err != nil {
			return err
		}

		_, err = f.Write(file.Content)
		if err != nil {
			return err
		}
	}

	for _, subdir := range dir.Directories {
		err = CreateDirAndFiles(*subdir)
		if err != nil {
			return
		}
	}
	if err = os.Chdir(".."); err != nil {
		return
	}
	return
}

func (root Directory) String() string {
	var sb strings.Builder
	var f func(string, Directory)
	f = func(path string, dir Directory) {
		for _, file := range dir.Files {
			sb.WriteString(filepath.Join(path, file.Name) + "\n")
		}
		for _, subdir := range dir.Directories {
			newPath := filepath.Join(path, subdir.Name)
			f(newPath, *subdir)
		}
	}
	f(root.Name, root)
	return sb.String()
}
