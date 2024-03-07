package tools

import (
	"fmt"
	"os"
)

// This function use recursivity to create subfiles and subdirectories
type File struct {
	Name    string
	Content []byte
}

type Directory struct {
	Name           string
	SubDirectories []Directory
	Files          []File
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
	for _, subdir := range dir.SubDirectories {
		CreateFileStructure(subdir)
	}
	return nil
}
