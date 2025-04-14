/*
Copyright © 2025 NAME HERE <nareshvaithi4@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)


var dbLoadCmd = &cobra.Command{
	Use:   "dbLoad",
	Short: "A brief description of your command",
	DisableFlagParsing: true,
	Long: `dbLoad`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("Please provide dbLoad flags eg,.. --user=root --password=pass --host=localhost")
			return
		}

		flagUser,err := cmd.Flags().GetString("user")
		if err != nil {
			fmt.Println("Error to fetch user: ",err)
			return
		}
		flaghost,err := cmd.Flags().GetString("host")
		if err != nil {
			fmt.Println("Error to fetch host: ",err)
			return
		}
		flagDirectory,err := cmd.Flags().GetString("directory")
		if err != nil {
			fmt.Println("Error to fetch directory: ",err)
			return
		}
		
		var hasUser = false
		var hasHost = false
		var hasDirectory = false

		for _,v := range args {
			if strings.HasPrefix(v,"--user=") || strings.HasPrefix(v,"-u=") {
				hasUser = true
			}
			if strings.HasPrefix(v,"--host=") || strings.HasPrefix(v,"-H=") {
				hasHost = true
			}
			if strings.HasPrefix(v,"--directory=") || strings.HasPrefix(v,"--D=") {
				hasDirectory = true
			}
		}

		if !hasUser {
			args = append(args, "--user="+flagUser)
		}
		if !hasHost {
			args = append(args, "--host="+flaghost)
		}
		if !hasDirectory {
			args = append(args, "--directory="+flagDirectory)
		}

		dbLoad := exec.Command("myloader",args...)
		dbLoad.Stdout = os.Stdout
		dbLoad.Stderr = os.Stderr

		err = dbLoad.Run()
		if err != nil {
			fmt.Println("Error executing command dbLoad\n", err)
		}
	},
}

func init() {
	
	dbLoadCmd.Flags().StringP("user","u","root","Database user")
	dbLoadCmd.Flags().StringP("host","H","localhost","Database server host name")
	dbLoadCmd.Flags().StringP("directory","d","/home/mydbops/hi","Restore file contains directory")
	rootCmd.AddCommand(dbLoadCmd)
}