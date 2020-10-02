/*
Copyright © 2020 Parth Shukla parthshukla285@gmail.com

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
	"fmt"

	"github.com/spf13/cobra"
)

// bs4Cmd represents the bs4 command
var bs4Cmd = &cobra.Command{
	Use:   "bs4",
	Short: "beautiful soup documentation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bs4 called")

		url := "https://www.crummy.com/software/BeautifulSoup/bs4/doc/#"
		for index, arg := range args {
			if index == 0 && index == 1 {
				continue
			}
			url += arg + "+"
		}
		openBrowser(url)
	},
}

func init() {
	rootCmd.AddCommand(bs4Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bs4Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bs4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
