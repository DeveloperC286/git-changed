package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	seachingCli := flag.String("searching", "", "The path to search for Git reposoties within.")
	flag.Parse()

	seaching := *seachingCli

	if *seachingCli == "" {
		currentWorkingDirectory, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		seaching = currentWorkingDirectory
	}

	repositories := getRepositories(seaching)
	fmt.Println(repositories)
}

func getRepositories(seaching string) []string {
	var repositories []string

	entries, err := os.ReadDir(seaching)

	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			path := seaching + "/" + entry.Name()

			if entry.Name() == ".git" {
				repositories = append(repositories, path)
			} else {
				repositories = append(repositories, getRepositories(path)...)
			}
		}
	}

	return repositories
}
