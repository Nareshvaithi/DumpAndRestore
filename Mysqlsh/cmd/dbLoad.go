package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var dbRestoreCmd = &cobra.Command{
	Use:   "dbLoad",
	Short: "Restore MySQL dump using mysqlsh",
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		dumpDir, _ := cmd.Flags().GetString("dumpdir")
		schema, _ := cmd.Flags().GetString("schema")

		// JavaScript command
		jsScript := fmt.Sprintf("util.loadDump('%s', {schema: '%s'})", dumpDir,schema)

		// mysqlsh command
		restoreCmd := exec.Command("mysqlsh",
			"--user="+user,
			"--password="+password,
			"--host="+host,
			"--port="+port,
			"--dumpdir"+dumpDir,
			"--js",
			"-e", jsScript,
		)

		restoreCmd.Stdout = os.Stdout
		restoreCmd.Stderr = os.Stderr

		if err := restoreCmd.Run(); err != nil {
			fmt.Println("Restore failed:", err)
		}
	},
}

func init() {
	dbRestoreCmd.Flags().String("user", "", "Database user")
	dbRestoreCmd.Flags().String("password", "", "Database password")
	dbRestoreCmd.Flags().String("host", "localhost", "Database host")
	dbRestoreCmd.Flags().String("port", "3306", "Database port")
	dbRestoreCmd.Flags().String("dumpdir", "", "Path to the dump folder to restore")
	dbRestoreCmd.Flags().String("schema", "", "Schema")

	dbRestoreCmd.MarkFlagRequired("user")
	dbRestoreCmd.MarkFlagRequired("password")
	dbRestoreCmd.MarkFlagRequired("dumpdir")
	dbRestoreCmd.MarkFlagRequired("schema")

	rootCmd.AddCommand(dbRestoreCmd)
}