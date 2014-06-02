package main

import (
	"fmt"
	"os"
	"os/exec"
)

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

func gitMust(conf *config, exitMessage string, args ...string) {
	out, err := git(conf, args...)
	if err != nil {
		fmt.Printf("%s\n\n%s\n", exitMessage, out)
		os.Exit(1)
	}
}

func gitAdd(conf *config, filename string, commit bool) {
	gitMust(
		conf, fmt.Sprintf("Couldn't add %s, does it exist?", filename),
		"add", filename,
	)
	if commit {
		gitCommit(conf, fmt.Sprintf("add %s", filename))
	}
}

func gitRm(conf *config, filename string, commit bool) {
	gitMust(
		conf, fmt.Sprintf("Couldn't remove %s.", filename),
		"rm", filename,
	)
	if commit {
		gitCommit(conf, fmt.Sprintf("remove %s", filename))
	}
}

func gitSave(conf *config, filename string, commit bool) {
	gitMust(
		conf, fmt.Sprintf("Couldn't add %s, does it exist?", filename),
		"add", filename,
	)
	if commit {
		gitCommit(conf, fmt.Sprintf("edit %s", filename))
	}
}

func gitCommit(conf *config, message string) {
	gitMust(
		conf, fmt.Sprintf("Could not commit! Check %s for problems.", conf.Dir),
		"commit", "-m", message,
	)
}
