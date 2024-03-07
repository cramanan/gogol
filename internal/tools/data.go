package tools

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

// This function use recursivity to create subfiles and Directories
type File struct {
	Name    string `yaml:"name"`
	Content []byte `yaml:"content"`
}

type Directory struct {
	Name        string
	Directories []Directory `yaml:"directories"`
	Files       []File      `yaml:"files"`
}

func RetrieveYAMLdir(name string) (dir *Directory, err error) {
	resp, err := http.Get("https://raw.githubusercontent.com/cramanan/gogol/cramanan/api/python.yaml")
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

func CreateFileStructure(dir Directory) (err error) {
	fmt.Println("++> Creating directory")
	if err = os.Mkdir(dir.Name, 0777); err != nil {
		return err
	}
	if err = os.Chdir(dir.Name); err != nil {
		return err
	}
	for _, file := range dir.Files {
		fmt.Printf("+++> Creating file : %s\n", file.Name)
		// Create file
		fil, err := os.Create(file.Name)
		if err != nil {
			return fmt.Errorf("%s can't be created", file.Name)
		}
		fmt.Printf("%s has been created\n", file.Name)

		if _, err = fil.Write(file.Content); err != nil {
			return err
		}
		defer fil.Close()
	}
	for _, subdir := range dir.Directories {
		CreateFileStructure(subdir)
	}
	return
}
