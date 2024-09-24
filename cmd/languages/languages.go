package languages

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cramanan/gogol/cmd"
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
	HasBoolFlag := command.PersistentFlags().GetBool
	for flag, filename := range cmd.FILES_FLAGS {
		if value, _ := HasBoolFlag(flag); value {
			cmd.RootDirectory.NewFile(filename)
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
	cmd.RootDirectory.Name = name
	return nil
}

func LanguagePostRunE(_ *cobra.Command, args []string) (err error) {
	// TODO: use args to determine destination
	return cmd.RootDirectory.Create(".")
}
