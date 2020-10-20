/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/TheCoy/razor/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// detectCmd represents the detect command
var detectCmd = &cobra.Command{
	Use:   "detect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		detectApp := app.RegexDetectApp{
			InputFileName:  viper.GetString("inputFile"),
			OutputFileName: viper.GetString("outputFile"),
			LogFileName:    viper.GetString("logFile"),
			DetectType:     viper.GetString("detectType"),
			Worker:         viper.GetInt("worker"),
		}
		fmt.Printf("【DetectApp】%s", detectApp.String())
		os.Exit(app.CLI(args, &detectApp))
	},
}

func init() {
	rootCmd.AddCommand(detectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// detectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// detectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	detectCmd.Flags().StringP("inputFile", "i", "./data/input.txt", "file name for inputFile")
	detectCmd.Flags().StringP("outputFile", "o", "./data/output.txt", "file name for outputFile")
	detectCmd.Flags().StringP("logFile", "l", "./log/detect.log", "file name for logFile")
	detectCmd.Flags().String("type", "html", "type for text detect")
	detectCmd.Flags().Int("worker", 1, "num for workers")

	viper.BindPFlag("inputFile", detectCmd.Flags().Lookup("inputFile"))
	viper.BindPFlag("outputFile", detectCmd.Flags().Lookup("outputFile"))
	viper.BindPFlag("logFile", detectCmd.Flags().Lookup("logFile"))
	viper.BindPFlag("detectType", detectCmd.Flags().Lookup("type"))
	viper.BindPFlag("worker", detectCmd.Flags().Lookup("worker"))
}
