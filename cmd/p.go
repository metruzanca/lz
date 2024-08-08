/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func shellCommand(name string, args []string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	return cmd.CombinedOutput()
}

// pCmd represents the p command
var pCmd = &cobra.Command{
	Use:   "p",
	Short: "Runs commands in parallel",
	Long: `Commands specified will have all their outputs to stdout/stderr combined into a single stdout/stderr with a prefix.
	
Great for running multiple build tools e.g. templ, tailwind and go from a single terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("p called")
		for _, item := range args {
			// TODO learn goroutine
			segments := strings.Fields(item)
			cmdName := segments[0]
			cmdArgs := segments[1:]
			stdout, stderr := shellCommand(cmdName, cmdArgs)
			if stderr != nil {
				fmt.Println("error:", stderr)
			} else {
				fmt.Println(string(stdout))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pCmd)
}
