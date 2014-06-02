package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var cmdShow = &cobra.Command{
	Use:   "show [project-name]",
	Short: "show the specified project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("I need a project name!")
			os.Exit(1)
		}
		conf := loadConfig()
		if err := os.Chdir(conf.Dir); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, name := range args {
			name = name + ".md"

			out, err := ioutil.ReadFile(name)
			if err != nil {
				fmt.Printf("Error reading %s!\n\n%s\n", name, err)
				os.Exit(1)
			}

			fmt.Print(string(out))
		}
	},
}
