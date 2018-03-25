package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/seriousben/net-script/internal/executor"
	"github.com/seriousben/net-script/internal/parser"
	"github.com/seriousben/net-script/internal/types"
	"github.com/spf13/cobra"
)

func parseCommands(filename string) ([]types.Command, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	cmds, err := parser.Parse(b)
	if err != nil {
		return nil, err
	}
	return cmds, nil
}

func main() {
	var cmdLint = &cobra.Command{
		Use:   "lint [file to lint]",
		Short: "Lint a netscript program",
		Long:  "Lint a netscript program",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Linting...")
			cmds, err := parseCommands(args[0])
			if err != nil {
				return err
			}
			fmt.Printf("%+v", cmds)
			return nil
		},
	}

	var cmdRun = &cobra.Command{
		Use:   "run [script to run]",
		Short: "Run a netscript program",
		Long:  "Run a netscript program",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmds, err := parseCommands(args[0])
			if err != nil {
				return err
			}
			return executor.ExecuteMany(cmds)
		},
	}

	var rootCmd = &cobra.Command{
		Use: "netscript",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmds, err := parseCommands(args[0])
			if err != nil {
				return err
			}
			return executor.ExecuteMany(cmds)
		},
		Args: cobra.MinimumNArgs(1),
	}
	rootCmd.AddCommand(cmdLint, cmdRun)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
