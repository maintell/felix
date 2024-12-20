package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/johntdyer/slackrus"

	"github.com/mojocn/felix/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "felix",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if isShowVersion {
			fmt.Println("Golang Env: %s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
			fmt.Println("UTC build time:%s", buildTime)
			fmt.Println("Build from Github repo version: https://github.com/mojocn/felix/commit/%s", gitHash)
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(bTime, gHash string) {
	buildTime = bTime
	gitHash = gHash
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var buildTime, gitHash string
var verbose, isShowVersion bool

func init() {
	cobra.OnInitialize(initFunc)
	rootCmd.Flags().BoolVarP(&isShowVersion, "version", "v", false, "show binary build information")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "V", false, "verbose")
}

func initFunc() {
	initViper()
	model.CreateSQLiteDb(verbose)
	initSlackLogrus()
}
func initSlackLogrus() {
	lvl := logrus.InfoLevel
	//钉钉群机器人API地址
	logrus.SetLevel(lvl)
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "06-01-02T15:04:05", PrettyPrint: true})
	logrus.SetReportCaller(true)
	mySlackApi := viper.GetString("felix.slack")
	logrus.AddHook(&slackrus.SlackrusHook{
		HookURL:        mySlackApi,
		AcceptedLevels: slackrus.LevelThreshold(logrus.ErrorLevel),
		Channel:        "#Felix",
		IconEmoji:      ":ghost:",
		Username:       "Felix",
	})

}

func initViper() {
	viper.SetConfigName("config")       // name of config file (without extension)
	viper.AddConfigPath(".")            // optionally look for config in the working directory
	viper.AddConfigPath("/etc/felix/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.felix") // call multiple times to add many search paths
	err := viper.ReadInConfig()         // Find and read the config file
	if err != nil {                     // Handle errors reading the config file
		logrus.Fatalf("Fatal error config file: %s \n", err)
	}
}
