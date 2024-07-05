package scripts

import (
	"os/exec"

	"github.com/cramanan/gogol/cmd"
	"github.com/spf13/cobra"
)

const GROUP_SCRIPTS = "SCRIPTS"

func init() {
	cmd.RootCmd.AddGroup(&cobra.Group{
		ID:    GROUP_SCRIPTS,
		Title: "Scripts",
	})
	cmd.RootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use: "update",
	RunE: func(command *cobra.Command, args []string) error {
		return exec.Command("go", "install", "github.com/cramanan/gogol@latest").Run()
	},
}
