package main

import (
	"github.com/cramanan/gogol/cmd"
	_ "github.com/cramanan/gogol/cmd/languages"
	_ "github.com/cramanan/gogol/cmd/scripts"
	_ "github.com/cramanan/gogol/cmd/utilities"
)

func main() {
	cmd.Execute()
}
