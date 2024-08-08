/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var red = color.New(color.FgRed).SprintFunc()

func shellCommand(name string, args []string) {
	cmd := exec.Command(name, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe:", err)
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("Error creating StderrPipe:", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if n > 0 {
				fmt.Print(string(buf[:n]))
			}
			if err != nil {
				break
			}
		}
	}()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stderr.Read(buf)
			if n > 0 {
				fmt.Print(red(string(buf[:n])))
			}
			if err != nil {
				break
			}
		}
	}()

	// Wait for the command to complete
	if err := cmd.Wait(); err != nil {
		fmt.Println("Command finished with error:", err)
	}
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
			shellCommand(cmdName, cmdArgs)
		}
	},
}

func init() {
	rootCmd.AddCommand(pCmd)
}
