/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package backup

import (
	"fmt"

	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var BackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "A brief description of your command",
	Long: `Safe-Sync is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("backup called")
	},
}

func init() {
	// rootCmd.AddCommand(backupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
