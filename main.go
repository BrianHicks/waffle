package main

import "github.com/spf13/cobra"

var root = &cobra.Command{
	Use: "waffle",
}

func init() {
	cmdConfig.AddCommand(cmdConfigInit, cmdConfigShow)

	root.AddCommand(cmdConfig, cmdOpen, cmdShow)
}

func main() {
	root.Execute()
}
