/*
Copyright © 2022 Daiki Narushima daikinarushima@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"a2aa/pkg/id3"
	"fmt"
	"github.com/spf13/afero"
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
		appFs := afero.NewOsFs()
		fileSystem := afero.NewIOFS(appFs)

		if allFlag {
			if err := id3.SetTagsAll(fileSystem, args[0]); err != nil {
				fmt.Println(err)
			}
		} else {
			if err := id3.SetTags(appFs, args[0]); err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.Flags().BoolP("all", "a", false, "get all files")
}
