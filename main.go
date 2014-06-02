package main

import "github.com/spf13/cobra"

var root = &cobra.Command{
	Use: "waffle",
}

func init() {
	cmdConfig.AddCommand(cmdInit, cmdShow)

	root.AddCommand(cmdConfig)
}

func main() {
	root.Execute()
}
