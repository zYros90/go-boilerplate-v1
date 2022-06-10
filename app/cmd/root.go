package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile      string // nolint:gochecknoglobals
	cfgMergeFile string // nolint:gochecknoglobals
)

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{ // nolint:exhaustruct,gochecknoglobals
	Use:   "template",
	Short: "A brief description of your application",
	Long:  `A longer description `,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() { // nolint:gochecknoinits
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "app/config/base.yaml", "config file")
	RootCmd.PersistentFlags().StringVar(&cfgMergeFile, "mergeconfig", "", "config to merge base config with")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
