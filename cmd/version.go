package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	//ApplicationName displays name of the app
	ApplicationName = "yessir"
)

var (
	//ApplicationVersion displays version of the app
	ApplicationVersion = ""
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display yessir's version",
	Long:  `Display API Server version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v-%v\n", ApplicationName, ApplicationVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
