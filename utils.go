package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"
)

func expandPath(in string) (out string) {
	usr, _ := user.Current()
	dir := usr.HomeDir

	if in[:2] == "~/" {
		out = strings.Replace(in, "~/", dir+"/", 1)
	} else {
		out = in
	}

	return
}

func questionStr(question string, def string) (response string) {
	var out string

	fmt.Printf("%s [%s]: ", question, def)
	fmt.Scanf("%s", &out)

	if out == "" {
		return def
	} else {
		return out
	}
}

func loadConfig() *config {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Trouble reading %s, have you run \"waffle config init\"?\n\n%s\n", configPath, err)
		os.Exit(1)
	}

	conf := new(config)
	err = json.Unmarshal(file, conf)
	if err != nil {
		fmt.Printf("Trouble parsing %s, is it valid JSON?\n\n%s\n", configPath, err)
		os.Exit(1)
	}

	return conf
}

func saveConfig(conf *config) error {
	out, _ := json.Marshal(conf)
	err := ioutil.WriteFile(configPath, out, 0755)

	return err
}

func exists(conf *config, filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func edit(conf *config, filename string) error {
	cmd := exec.Cmd{
		Path: conf.Editor,
		Args: append(
			[]string{conf.Editor},
			path.Join(conf.Dir, filename),
		),
		Dir: conf.Dir,

		// pass these through!
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	return cmd.Run()
}
