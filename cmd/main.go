package main

import (
	"fmt"
	"os"

	"github.com/leonardodelira/go-git-clone/command"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fit",
	Short: "fit is a minimalist git cli tool",
	Long:  "fit is a minimalist git cli tool",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// TODO: fit status - step 6
func main() {
	rootCmd.AddCommand(command.Init)
	rootCmd.AddCommand(command.Add)
	rootCmd.AddCommand(command.Status)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
