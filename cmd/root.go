package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command {
	Use: "go-task",
	Short: "go-task is a CLI task manager written in Go",
}