package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configPath = expandPath("~/.waffle.json")

type config struct {
	Dir string
}

var (
	cmdConfig = &cobra.Command{
		Use: "config",
	}

	cmdInit = &cobra.Command{
		Use:   "init",
		Short: "initialize a configuration (or overwrite)",
		Run: func(cmd *cobra.Command, args []string) {
			var conf config

			conf.Dir = expandPath(questionStr("Where should the projects be kept?", "~/waffle"))

			if err := saveConfig(&conf); err != nil {
				fmt.Printf("shoot, we couldn't write!\n\n%s\n", err)
				os.Exit(1)
			}

			fmt.Println("Thanks, you're all good to go!")
		},
	}

	cmdShow = &cobra.Command{
		Use:   "show",
		Short: "show the current config",
		Run: func(cmd *cobra.Command, args []string) {
			conf := loadConfig()

			fmt.Printf("Dir: %s\n", conf.Dir)
		},
	}
)
