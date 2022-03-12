/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"a2aa/pkg/dir"
	"fmt"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"os"
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
		appFs := afero.NewOsFs()
		fileSystem := os.DirFS("./")

		if allFlag {
			files, err := dir.All(fileSystem, args[0])
			if err != nil {
				fmt.Println(err)
			}
			for _, file := range files {
				fmt.Println(file)
			}
		} else {
			files, err := dir.Current(appFs, args[0])
			if err != nil {
				fmt.Println(err)
			}
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
