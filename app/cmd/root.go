package cmd

import (
	"log"

	"github.com/zYros90/go-boilerplate-v1/app/config"
	"github.com/zYros90/pkg/logger"

	"github.com/spf13/cobra"
)

var (
	cfgFile      string
	cfgMergeFile string
	conf         *config.Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "template",
	Short: "A brief description of your application",
	Long:  `A longer description `,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "app/config/base.yaml", "config file")
	rootCmd.PersistentFlags().StringVar(&cfgMergeFile, "mergeconfig", "", "config to merge base config with")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	conf = config.NewConfig()
	logger, err := logger.NewLogger("debug", true, false, false)
	if err != nil {
		log.Fatal("error creating new logger", err)
	}

	err = config.ReadConfig(conf, cfgFile, cfgMergeFile)
	if err != nil {
		logger.Sugar().Fatal(err)
	}
}
