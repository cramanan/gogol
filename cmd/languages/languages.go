package languages

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cramanan/gogol/cmd"
	"github.com/cramanan/gogol/filesystem"
	"github.com/spf13/cobra"
)

const GROUP_LANG = "LANG"

func init() {
	cmd.RootCmd.AddGroup(&cobra.Group{
		ID:    GROUP_LANG,
		Title: "Languages",
	})
}

func LanguagePreRunE(command *cobra.Command, _ []string) error {
	// TODO: use args to determine destination
	HasBoolFlag := command.Root().PersistentFlags().GetBool
	for flag, filename := range cmd.FILES_FLAGS {
		if ok, _ := HasBoolFlag(flag); ok {
			filesystem.RootDirectory.NewFile(filename)
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Project name: ")
	name, err := reader.ReadString('\n')
	if strings.Contains(name, "/") || strings.Contains(name, ".") {
		return errors.New("project name cannot contain '/' or '.'")
	}
	if err != nil {
		return err
	}
	name = name[:len(name)-1]
	if name == "" {
		name = "untitled"
	}
	filesystem.RootDirectory.Name = name
	return nil
}

func LanguagePostRunE(command *cobra.Command, args []string) (err error) {
	// TODO: use args to determine destination

	if ok, _ := cmd.RootCmd.PersistentFlags().GetBool(cmd.FLAG_GITHUB); ok {
		root := filesystem.RootDirectory
		root.NewFile(".gitignore")
		for _, directory := range root.Directories {
			if len(directory.Files) == 0 {
				directory.NewFile(".gitkeep")
			}
		}

	}

	return filesystem.RootDirectory.Create(".")
}
