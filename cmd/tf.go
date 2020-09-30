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

// tfCmd represents the tf command
var tfCmd = &cobra.Command{
	Use:   "tf",
	Short: "Tensorflow documentation. Note: please call keras for keras documentation",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tf called")
		// TODO: clean argument entered
		url := "https://www.tensorflow.org/api_docs/python/tf/" + args[len(args)-1]

		openBrowser(url)
	},
}

func init() {
	rootCmd.AddCommand(tfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
