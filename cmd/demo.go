package cmd

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
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
	PreRun: func(cmd *cobra.Command, args []string) {
		_ = viper.BindPFlag("qps", cmd.Flags().Lookup("qps"))

		//在这里监听各自感兴趣的参数变化
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("config changed:", e.Name)
			newQps := viper.GetInt64("qps")
			_ = demoApp.SetQps(int(newQps))
			fmt.Println("new qps:", demoApp.Qps)
		})
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("flag args:", args)
		fmt.Println("viper:author=", viper.GetString("logFile"))
		demoApp.Qps = viper.GetInt("qps")
		os.Exit(app.CLI(args, &demoApp))
	},
}

func init() {
	rootCmd.AddCommand(demoCmd)

	demoCmd.Flags().StringVar(&demoApp.LogFileName, "logFile", "./log/demoApp.log", "file path to save logFile")
	demoCmd.Flags().IntVar(&demoApp.MaxWorker, "worker", 1, "max worker num")
	demoCmd.Flags().Int64VarP(&demoApp.Times, "times", "t", 10, "times for task cycle")
	demoCmd.Flags().IntVar(&demoApp.Qps, "qps", 10, "rate to execute")
}
