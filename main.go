package main

import (
	"bytes"
	"flag"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	seachingCli := flag.String("searching", "", "The path to search for Git reposoties within.")
	flag.Parse()

	seaching := *seachingCli

	if *seachingCli == "" {
		currentWorkingDirectory, err := os.Getwd()

		if err != nil {
			log.Fatal(err)
		}

		seaching = currentWorkingDirectory
	}

	// Getting user home directory so we can find replace ~ as Go does not handle it.
	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	home := usr.HomeDir

	if strings.HasPrefix(seaching, "~/") {
		seaching = filepath.Join(home, seaching[2:])
	}

	repositories := getRepositories(seaching)

	for _, repository := range repositories {
		// Checking if their are uncommited local changes.
		cmd := exec.Command("git", "diff-index", "HEAD")
		// Execute the command inside specfic Git repository.
		cmd.Dir = repository
		// Setup reading the command's output.
		stdout := new(bytes.Buffer)
		stderr := new(bytes.Buffer)
		cmd.Stdout = stdout
		cmd.Stderr = stderr

		err := cmd.Run()

		if err != nil {
			log.Error(err.Error())
		}

		if stdout.String() != "" {
			log.Infof("Has %s has uncommited local changes.", repository)
		}
	}
}

func getRepositories(seaching string) []string {
	var repositories []string

	entries, err := os.ReadDir(seaching)

	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {

			if entry.Name() == ".git" {
				repositories = append(repositories, seaching)
			} else {
				path := seaching + "/" + entry.Name()
				repositories = append(repositories, getRepositories(path)...)
			}
		}
	}

	return repositories
}
