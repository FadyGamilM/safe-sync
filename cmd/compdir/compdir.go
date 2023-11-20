/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package compdir

import (
	"fmt"

	"github.com/spf13/cobra"
)

// compdirCmd represents the compdir command
var CompdirCmd = &cobra.Command{
	Use:   "compdir",
	Short: "compdir stands for compress-directory",
	Long:  `Compress folder`,
	// we expect 2 args :- the src and the destination paths
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("compdir called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compdirCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compdirCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
