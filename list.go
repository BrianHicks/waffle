package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "list projects",
	Run: func(cmd *cobra.Command, args []string) {
		conf := loadConfig()

		files, err := ioutil.ReadDir(conf.Dir)
		if err != nil {
			fmt.Printf("Error reading %s:\n\n%s\n", conf.Dir, err)
			os.Exit(1)
		}

		for _, file := range files {
			name := file.Name()
			if strings.HasSuffix(name, ".md") {
				fmt.Println(name[:len(name)-3])
			}
		}
	},
}
