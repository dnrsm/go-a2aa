/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"a2aa/pkg/id3"

	"github.com/spf13/cobra"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		allFlag, _ := cmd.Flags().GetBool("all")

		if allFlag {
			id3.SetTagsAll(args[0])
		} else {
			id3.SetTags(args[0])

		}
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.Flags().BoolP("all", "a", false, "get all files")
}
