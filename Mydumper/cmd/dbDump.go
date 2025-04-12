/*
Copyright © 2025 NAME HERE <nareshvaithi4@gmail.com>
*/
package cmd

import (
	"fmt"
	"strings"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)


var dbDumpCmd = &cobra.Command{
	Use:   "dbDump",
	Short: "A brief description of your command",
	DisableFlagParsing: true,
	Long:  `dbDump used `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide flags for dbDump, e.g., --user=root --password=pass --host=localhost")
			return
		}

		userFlag, err := cmd.Flags().GetString("user")
		hostFlag, err := cmd.Flags().GetString("host")
		outputdirFlag, err := cmd.Flags().GetString("outputdir")
		if err != nil {
			fmt.Println("Error fetching user flag:", err)
			return
		}

		hasUser := false
		hasHost := false
		hasOutputdir := false

		for _, v := range args {
			if strings.HasPrefix(v, "--user=") || strings.HasPrefix(v, "-u=") {
				hasUser = true
			}
			if strings.HasPrefix(v,"--host=") || strings.HasPrefix(v,"-h=") {
				hasHost = true
			}

			if strings.HasPrefix(v,"--outputdir=") || strings.HasPrefix(v,"-o=") {
				hasOutputdir = true
			}
			
		}

		if !hasUser {
			args = append(args, "--user="+userFlag)
		}

		if !hasHost {
			args = append(args, "--host="+hostFlag)
		}

		if !hasOutputdir {
			args = append(args,"--outputdir="+outputdirFlag)
		}

		dbDump := exec.Command("mydumper", args...)
		dbDump.Stdout = os.Stdout
		dbDump.Stderr = os.Stderr
		if err := dbDump.Run(); err != nil {
			fmt.Println("Error executing command:", err)
		}
	},
}

func init() {
	dbDumpCmd.Flags().StringP("user", "u", "root", "Database user")
	dbDumpCmd.Flags().StringP("host", "H", "localhost", "Database server host name")
	dbDumpCmd.Flags().StringP("outputdir", "o", "/home/mydbops/newBackup", "backup files store path")
	rootCmd.AddCommand(dbDumpCmd)
}

