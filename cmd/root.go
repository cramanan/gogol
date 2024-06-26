/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
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

type File struct {
	Name    string
	Content []byte
}

func NewFile(name string) *File {
	return &File{
		Name: name,
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gogol",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {},
}

type CobraFunc func(cmd *cobra.Command, args []string, root *Directory)

func GenerateFS(fn CobraFunc) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		root := NewDirectory("project")
		files := map[string]string{
			"readme":     ("README.md"),
			"license":    ("LICENSE.md"),
			"dockerfile": ("Dockerfile"),
		}

		for flag := range files {
			value, _ := rootCmd.PersistentFlags().GetBool(flag)
			if value {
				root.Files[flag] = NewFile(files[flag])
			}
		}

		var err error

		fn(cmd, args, root)
		err = root.Create(".")
		if err != nil {
			fmt.Println(err)
			return
		}

		root.repo, err = git.PlainInit(root.Name, false)
		if err != nil {
			return
		}
	}
}

func (root Directory) Create(origin string) (err error) {
	var f func(string, Directory)
	f = func(path string, dir Directory) {
		err = os.Mkdir(root.Name, os.ModePerm)
		if err != nil {
			return
		}
		for _, file := range dir.Files {
			ff, err := os.Create(filepath.Join(origin, path, file.Name))
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gogol.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolP("readme", "r", false, "Help message for readme")
	rootCmd.PersistentFlags().BoolP("license", "l", false, "Help message for license")
	rootCmd.PersistentFlags().BoolP("dockerfile", "d", false, "Help message for license")
}
