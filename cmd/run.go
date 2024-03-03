/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: Run,
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("config", "c", "config.yaml", "Pipeline configuration")
}

type RunConfig struct {
	Actions []Action `yaml:"actions"`
}

type Action struct {
	Cmd    string
	Data   string
	Repeat int
}

func Run(cmd *cobra.Command, args []string) {
	yamlFile, err := os.ReadFile(cmd.Flags().Lookup("config").Value.String())
	if err != nil {
		panic(err)
	}

	var config RunConfig

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalln(err)
	}

	err = RunActions(&config.Actions)

	if err != nil {
		log.Fatalln(err)
	}
}

func ExecuteCommand(action Action) ([]byte, error) {
	args := strings.Split(action.Cmd, " ")
	fmt.Println()
	cmd := exec.Command(args[0], args[1:]...)

	output, err := cmd.Output()
	if err != nil {
		return output, fmt.Errorf("issue with exec command %s", err)
	}
	return output, nil
}

func RunActions(actions *[]Action) error {
	for _, action := range *actions {
		numRepeats := 1
		if action.Repeat > 1 {
			numRepeats = action.Repeat
		}
		for i := 0; i < numRepeats; i++ {
			output, err := ExecuteCommand(action)
			if err != nil {
				return err
			}
			fmt.Println(string(output))
		}
	}
	return nil
}
