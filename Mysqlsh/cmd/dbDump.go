/*
Copyright © 2025 NAME HERE <nareshvaithi4@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)


var dbDumpCmd = &cobra.Command{
	Use:   "dbDump",
	Short: "Backup MySQL using mysqlsh",
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		outputDir, _ := cmd.Flags().GetString("outputdir")
		schema, _ := cmd.Flags().GetString("schema")
	
	
		jsScript := fmt.Sprintf("util.dumpSchemas(['%s'], '%s', {ocimds: false})", schema, outputDir)
	
		// mysqlsh command
		dumpCmd := exec.Command("mysqlsh",
			"--user="+user,
			"--password="+password,
			"--host="+host,
			"--port="+port,
			"--js",
			"-e", jsScript,
		)
	
		dumpCmd.Stdout = os.Stdout
		dumpCmd.Stderr = os.Stderr
	
		if err := dumpCmd.Run(); err != nil {
			fmt.Println("Backup failed:", err)
		}
	},
	
}


func init() {
	dbDumpCmd.Flags().String("user", "", "Database user")
	dbDumpCmd.Flags().String("password", "", "Database password")
	dbDumpCmd.Flags().String("host", "localhost", "Database host")
	dbDumpCmd.Flags().String("port", "3306", "Database port")
	dbDumpCmd.Flags().String("outputdir", "", "Backup destination path")
	dbDumpCmd.Flags().String("util", "dumpSchema", "mysqlsh util function")
	dbDumpCmd.Flags().String("schema", "", "Schema name (required for dumpSchema)")

	dbDumpCmd.MarkFlagRequired("user")
	dbDumpCmd.MarkFlagRequired("password")
	dbDumpCmd.MarkFlagRequired("outputdir")
	dbDumpCmd.MarkFlagRequired("schema")

	rootCmd.AddCommand(dbDumpCmd)
}




