/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"safesync/cmd/compdir"

	cliFigures "github.com/common-nighthawk/go-figure"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "safesync",
	Short: "A CLI tool for working with files, folders, getting backups of them easily, working with compression and decompression of folders",
	Long: `Safe-Sync is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func registerCommands() {
	// rootCmd.AddCommand(backup.BackupCmd)
	rootCmd.AddCommand(compdir.CompdirCmd)
}

// --> runs after we run the main.go
func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.safesync.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//! display nice figure
	fig := cliFigures.NewFigure("Safe-Sync", "doom", true)
	fig.Print()

	// ! call the register commands method
	registerCommands()
}
