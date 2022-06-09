package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile      string
	cfgMergeFile string
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "app/config/base.yaml", "config file")
	rootCmd.PersistentFlags().StringVar(&cfgMergeFile, "mergeconfig", "", "config to merge base config with")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
