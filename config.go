package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var configPath = expandPath("~/.waffle.json")

type config struct {
	Dir    string
	Editor string
}

var (
	cmdInit = &cobra.Command{
		Use:   "init",
		Short: "initialize a configuration (or overwrite)",
		Run: func(cmd *cobra.Command, args []string) {
			conf := new(config)

			conf.Dir = expandPath(questionStr("Where should the projects be kept?", "~/waffle"))
			// check if that directory exists, and create it
			if _, err := os.Stat(conf.Dir); err != nil {
				if os.IsNotExist(err) {
					fmt.Println("making that directory real quick...")
					err := os.MkdirAll(conf.Dir, 0755)
					if err != nil {
						fmt.Printf("Something went wrong making that directory.\n\n%s\n", err)
						os.Exit(1)

					}

					out, err := git(conf, "init", ".")
					if err != nil {
						fmt.Printf("Something went wrong initializing the git repo...\n\n%s\n%s\n", out, err)
						os.Exit(1)
					}
				}
			}

			conf.Editor = expandPath(questionStr(
				"What editor do you want to use?",
				os.Getenv("EDITOR"),
			))
			fullPath, err := exec.LookPath(conf.Editor)
			if err != nil {
				fmt.Printf("I couldn't find that in the PATH, what's up with that?\n\n%s\n", err)
			}
			conf.Editor = fullPath

			if err := saveConfig(conf); err != nil {
				fmt.Printf("shoot, we couldn't write!\n\n%s\n", err)
				os.Exit(1)
			}

			fmt.Println("Thanks, you're all good to go!")
		},
	}

	cmdConfig = &cobra.Command{
		Use:   "config",
		Short: "show the current config",
		Run: func(cmd *cobra.Command, args []string) {
			conf := loadConfig()

			fmt.Printf("Dir: %s\n", conf.Dir)
			fmt.Printf("Editor: %s\n", conf.Editor)
		},
	}
)
