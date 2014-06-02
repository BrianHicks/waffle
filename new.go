package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var cmdOpen = &cobra.Command{
	Use:   "open [project-name]",
	Short: "create or edit a project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("I need a project name!")
		}
		conf := loadConfig()
		if err := os.Chdir(conf.Dir); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, name := range args {
			name = name + ".md"

			present := exists(conf, name)
			var pre []byte
			if present {
				pre, _ = ioutil.ReadFile(name)
			} else {
				pre = []byte{}
			}

			err := edit(conf, name)
			if err != nil {
				fmt.Println("It looks like your editor exited with an error. I'm not going to trust this and fail.")
				os.Exit(1)
			}

			saved := exists(conf, name)
			var post []byte
			if saved {
				post, _ = ioutil.ReadFile(name)
			} else {
				post = []byte{}
			}

			if !present && saved {
				gitAdd(conf, name, true)
				fmt.Printf("Added %s\n", name)
			} else if present && !saved {
				gitRm(conf, name, true)
				fmt.Printf("Removed %s\n", name)
			} else if present && saved {
				if bytes.Equal(pre, post) {
					fmt.Printf("No change detected.")
				} else {
					gitSave(conf, name, true)
					fmt.Printf("Saved %s\n", name)
				}
			} else {
				fmt.Println("No change detected.")
			}
		}
	},
}
