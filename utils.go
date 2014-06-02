package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
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

func git(conf *config, args ...string) (stdout []byte, err error) {
	path, err := exec.LookPath("git")
	if err != nil {
		fmt.Printf("Couldn't find git! Is it installed?\n\n%s\n", err)
		os.Exit(1)
	}

	cmd := exec.Cmd{
		Path: path,
		Args: append([]string{path}, args...),
		Dir:  conf.Dir,
	}

	stdout, err = cmd.CombinedOutput()

	return
}
