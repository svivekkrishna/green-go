/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"green-go/cli/reporting"
	"green-go/lib"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "green-go",
	Short: "A tool to quickly run status checks for a set of applications you are interested in.",
	Long:  `This is a long description which I'll add later`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		yamlPath := cmd.Flag("hosts-from-file").Value.String()
		if yamlPath == "" {
			fmt.Println("Add logic to read from a default path on PWD if exists.")
		}

		outputReporter := reporting.GetByType(cmd.Flag("format").Value.String())

		content, err := ioutil.ReadFile(yamlPath)
		if err != nil {
			log.Fatal("Error occured while parsing YAML ", err.Error())
			return
		}

		var endpoints []lib.Endpoint

		err = yaml.Unmarshal(content, &endpoints)

		if err != nil {
			log.Fatal("Failed to parse file ", err)
		}

		outputReporter.Render(lib.PerformChecks(endpoints))

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("hosts-from-file", "f", "", "File to read the hosts info.")
	rootCmd.Flags().StringP("format", "o", "table", "Output format.")
}
