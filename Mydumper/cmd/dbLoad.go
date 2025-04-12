/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dbLoadCmd represents the dbLoad command
var dbLoadCmd = &cobra.Command{
	Use:   "dbLoad",
	Short: "A brief description of your command",
	Long: `dbLoad`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dbLoad called")
	},
}

func init() {
	rootCmd.AddCommand(dbLoadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbLoadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbLoadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
