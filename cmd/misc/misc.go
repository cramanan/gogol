package misc

import (
	"github.com/cramanan/gogol/cmd"
	"github.com/spf13/cobra"
)

const GROUP_MISC = "Miscellaneous"

func init() {
	cmd.RootCmd.AddGroup(&cobra.Group{
		ID:    GROUP_MISC,
		Title: "Miscellaneous",
	})
}
