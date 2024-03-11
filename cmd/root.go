/*
Copyright © 2024 MATHIAS MARCHETTI aquemaati@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Mkdir(name string) (err error) {
	err = os.Mkdir(name, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Chdir(name)
	if err != nil {
		return err
	}
	return
}

func InternalError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

var rootCmd = &cobra.Command{
	Use:     "gogol",
	Short:   "Create projects faster than ever.",
	Long:    "\r\ngogol is a library that helps you create projects using multiple languages.\r\nTo create a project, run:\r\n  $ gogol [language/command] [flags...]\r\n\r\n",
	Example: "gogol go\r\ngogol html -rlg",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var README, LICENSE, GIT, DOCKERFILE bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&README, "readme", "r", false, "Add a README.md to your project.")
	rootCmd.PersistentFlags().BoolVarP(&LICENSE, "license", "l", false, "Add a LICENSE.md to your project.")
	rootCmd.PersistentFlags().BoolVarP(&GIT, "github", "g", false, "Change your project into a Git repository.")
	rootCmd.PersistentFlags().BoolVarP(&DOCKERFILE, "dockerfile", "d", false, "Add a Dockerfile to your project.")
}
