package cmd

import (
	"fmt"
	"os"

	"github.com/TheCoy/razor/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var demoApp app.DemoApp

var demoCmd = &cobra.Command{
	Use:              "demo",
	Short:            "Demo Command of Razor toolkit",
	Long:             `Demo Command shows how to add ohter commands with different functions`,
	Args:             cobra.ArbitraryArgs,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("flag args:", args)
		fmt.Println("viper:author=", viper.GetString("logFile"))
		os.Exit(app.CLI(args, &demoApp))
	},
}

func init() {
	rootCmd.AddCommand(demoCmd)

	demoCmd.Flags().StringVar(&demoApp.LogFileName, "logFile", "./log/demoApp.log", "file path to save logFile")
	demoCmd.Flags().IntVar(&demoApp.MaxWorker, "worker", 1, "max worker num")
	demoCmd.Flags().Int64VarP(&demoApp.Times, "times", "t", 10, "times for task cycle")
}
