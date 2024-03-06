/*
Copyright © 2024 MATHIAS MARCHETTI aquemaati@gmail.com
*/
package cmd

import (
	"bufio"
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

func YesNo(question string) (string, error) {
	fmt.Print(question + " [y/n]: ")
	reader := bufio.NewReader(os.Stdin)
	choice, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return choice[:len(choice)-1], nil
}

func InternalError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "GOGOL",
	Short: "Create projects faster than ever.",
	Long:  `GOGOL Copyright © 2024 MATHIAS MARCHETTI aquemaati@gmail.com`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {},
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pre-master-v2.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
