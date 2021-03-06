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
	"github.com/fsnotify/fsnotify"

	"github.com/TheCoy/razor/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	robot app.RobotApp
)

// sizeCmd represents the size command
var sizeCmd = &cobra.Command{
	Use:   "size",
	Short: "A Command which can read config dynamically",
	Long: `A Command which can read config dynamically, For example:

razor size --config ./size.yml`,
	PreRun: func(cmd *cobra.Command, args []string) {
		_ = viper.BindPFlag("worker", cmd.Flags().Lookup("worker"))
		_ = viper.BindPFlag("qps", cmd.Flags().Lookup("qps"))
		_ = viper.BindPFlag("times", cmd.Flags().Lookup("times"))

		//在这里监听各自感兴趣的参数变化
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("config changed:", e.Name)
			newQps := viper.GetInt64("qps")
			robot.SetNewLimiter(int(newQps))
		})
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("size called")
		fmt.Println("input args:", cmd.PersistentFlags().Args())
		fmt.Println("flag args:", cmd.Flags().Args())

		robot = app.RobotApp{
			Times:  viper.GetInt64("times"),
			Worker: viper.GetInt("worker"),
			QPS:    viper.GetInt64("qps"),
		}

		app.CLI(args, &robot)
	},
}

func init() {
	rootCmd.AddCommand(sizeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sizeCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	sizeCmd.Flags().Int64("times", 100, "total times for request")
	sizeCmd.Flags().Int64("qps", 1, "qps for requet")
	sizeCmd.Flags().Int("worker", 1, "num for workers")
}
