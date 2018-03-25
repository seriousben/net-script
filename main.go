package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"regexp"
	"strings"

	"github.com/seriousben/is-it-binary/binary"
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

func runCommands(cmds []types.Command) {
	for _, cmd := range cmds {
		fmt.Printf("# %s\n%s %s\n", cmd.Comment, cmd.Method, cmd.URL)
		out, err := executor.Execute(cmd)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		bodyStr := ""
		if resp, ok := out.Response.(*http.Response); ok {
			db, err := httputil.DumpResponse(resp, false)
			if err != nil {
				fmt.Println("Error dumping response")
			} else {
				rawS := strings.Trim(string(db), " \n\r")
				re := regexp.MustCompile("\n")
				s := re.ReplaceAllString(rawS, "\n\\\\ ")
				fmt.Println("\\\\ ", s)
			}
			isBin, err := binary.IsBinaryBuffer(out.Body)
			if err == nil && !isBin {
				bodyStr = string(out.Body)
			} else {
				bodyStr = fmt.Sprintf("<No Text Representation (type: %s)>\n", resp.Header.Get("Content-Type"))
			}
		}
		fmt.Printf("\n%s\n", bodyStr)
	}
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
			runCommands(cmds)
			return nil
		},
	}

	var rootCmd = &cobra.Command{
		Use: "netscript",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmds, err := parseCommands(args[0])
			if err != nil {
				return err
			}
			runCommands(cmds)
			return nil
		},
		Args: cobra.MinimumNArgs(1),
	}
	rootCmd.AddCommand(cmdLint, cmdRun)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
