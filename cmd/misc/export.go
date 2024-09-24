// Copyright © 2024 cramanan cramananjaonapro@gmail.com
package misc

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/cramanan/gogol/cmd"
	"github.com/cramanan/gogol/filesystem"
	"github.com/spf13/cobra"
)

func parseEntry(dir *filesystem.Directory, entry fs.DirEntry, path string) (err error) {
	name := entry.Name()
	if entry.Type() == fs.ModeDir {
		subdir := dir.NewDirectory(name)
		subentries, err := os.ReadDir(filepath.Join(path, name))
		if err != nil {
			return err
		}

		for _, subentry := range subentries {
			err = parseEntry(subdir, subentry, filepath.Join(path, name))
			if err != nil {
				return err
			}
		}
	} else {
		dir.NewFile(name)
	}
	return err
}

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:     "export",
	Short:   "",
	Long:    ``,
	GroupID: GROUP_MISC,
	RunE: func(_ *cobra.Command, args []string) (err error) {
		args = append(args, ".")
		entries, err := os.ReadDir(args[0])
		if err != nil {
			return err
		}

		for _, entry := range entries {
			err = parseEntry(cmd.RootDirectory, entry, args[0])
			if err != nil {
				return err
			}
		}

		delete(cmd.RootDirectory.Directories, ".git")

		export, err := os.Create("export.json")
		if err != nil {
			return err
		}
		defer export.Close()

		encoder := json.NewEncoder(export)
		encoder.SetIndent("", "  ")
		encoder.Encode(cmd.RootDirectory)

		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(exportCmd)
}
