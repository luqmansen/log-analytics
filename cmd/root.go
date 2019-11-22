/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var directory, fileName string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "analytics",
	Short: "cli-tool for display log file analytics",
	Long: `cli-tool for display log file analytics, specifically to read log file that
follow Common Log Format (https://en.wikipedia.org/wiki/Common_Log_Format).
Example provided is the default log file for access log on apache server`,
	Run: analytics,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.analytics.yaml)")
	rootCmd.Flags().StringVarP(&directory, "directory", "d", ".", "Directory where the log file lies")
	rootCmd.Flags().StringVarP(&fileName, "file-name-pattern", "f", "access.log.", "file name pattern to look for, follow apache log file naming format, just omit the number ")
	rootCmd.Flags().IntP("time", "t", 30, "Log file for the last \"t\" minute")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".analytics" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".analytics")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
