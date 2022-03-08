/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-a2aa/pkg/dir"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		allFlag, _ := cmd.Flags().GetBool("all")

		if allFlag {
			files := dir.All(args[0])
			for _, file := range files {
				fmt.Println(file)
			}
		} else {
			files := dir.Current(args[0])
			for _, file := range files {
				fmt.Println(file)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().BoolP("all", "a", false, "get all files")
}
