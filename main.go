package main

import "github.com/spf13/cobra"

var root = &cobra.Command{
	Use: "waffle",
}

func init() {
	root.AddCommand(cmdOpen, cmdShow, cmdInit, cmdConfig)
}

func main() {
	root.Execute()
}
