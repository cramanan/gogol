package main

import (
	"github.com/cramanan/gogol/cmd"
	_ "github.com/cramanan/gogol/cmd/languages"
	_ "github.com/cramanan/gogol/cmd/misc"
)

func main() { cmd.Execute() }
