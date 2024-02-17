/*
Copyright © 2024 Kian Musser
Freely available under the MIT license
*/
package cmd

import (
	"os"

	"github.com/kianmusser/unjn/server"
	"github.com/spf13/cobra"
)

var ntfyTopic string
var upworkRssUrls []string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "unjn",
	Short: "Upwork New Jobs Notifier",
	Long:  `unjn is a simple program to notify you of any new jobs on Upwork via ntfy.sh`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewServer(ntfyTopic, upworkRssUrls)
		s.Run()
	},
}

func init() {
	rootCmd.Flags().StringVarP(&ntfyTopic, "topic", "t", "", "the ntfy topic")
	rootCmd.Flags().StringArrayVarP(&upworkRssUrls, "urls", "u", []string{}, "a comma separated list of Upwork RSS urls")
	rootCmd.MarkFlagRequired("topic")
	rootCmd.MarkFlagRequired("urls")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
