/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/metruzanca/lz/cmd"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		// Handle cleanup here
		os.Exit(0)
	}()

	cmd.Execute()
}
