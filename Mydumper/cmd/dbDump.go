package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

var dbDumpCmd = &cobra.Command{
	Use:                "dbDump",
	Short:              "Dump MySQL using mydumper",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
	
		if configFile == "" {
			configFile = "./cmd/config.yaml"
		}
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("unable to read config file:", err)
		}


		passed := make(map[string]bool)
		for _, arg := range args {
		
			if strings.HasPrefix(arg, "--") {
				key := strings.SplitN(arg[2:], "=", 2)[0]
				passed[key] = true
			}
		}

		
		for _, key := range viper.AllKeys() {
			if !passed[key] {
				val := viper.GetString(key)
				if val != "" {
					args = append(args, fmt.Sprintf("--%s=%s", key, val))
				}
			}
		}

	
		cmdExec := exec.Command("mydumper", args...)
		cmdExec.Stdout = os.Stdout
		cmdExec.Stderr = os.Stderr
		if err := cmdExec.Run(); err != nil {
			fmt.Println("Error executing command dbDump:", err)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "Custom config file path")
	rootCmd.AddCommand(dbDumpCmd)
}
