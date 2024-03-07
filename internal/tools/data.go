package tools

import (
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type File struct {
	Name    string `yaml:"name"`
	Content []byte `yaml:"content"`
}

type Directory struct {
	Name        string      `yaml:"name"`
	Directories []Directory `yaml:"directories"`
	Files       []File      `yaml:"files"`
}

func RetrieveYAMLdir(url string) (dir Directory, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if err = yaml.Unmarshal(buf, &dir); err != nil {
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
		os.Create(file.Name)
	}

	for _, subdir := range dir.Directories {
		err = CreateDirAndFiles(subdir)
		if err != nil {
			return
		}
	}
	if err = os.Chdir(".."); err != nil {
		return
	}
	return
}
