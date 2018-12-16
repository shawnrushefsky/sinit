package main

import (
	"flag"
	"fmt"
	"strings"

	"./src/sinit"
)

func main() {
	var (
		projName  string
		stack     string
		theme     string
		themeRepo string
		deploy    string
	)

	const (
		projNameDefault = "new-project"
		projNameUsage   = "The name of your new project"
	)
	flag.StringVar(&projName, "name", projNameDefault, projNameUsage)
	flag.StringVar(&projName, "n", projNameDefault, projNameUsage+" (shorthand)")

	const (
		stackDefault = "node"
		stackUsage   = "The project stack you want to create. e.g. node"
	)
	flag.StringVar(&stack, "stack", stackDefault, stackUsage)
	flag.StringVar(&stack, "s", stackDefault, stackUsage+" (shorthand)")

	const (
		themeDefault = "mediumish"
		themeUsage   = "The name of the theme you want to use (for static sites, etc.)"
	)
	flag.StringVar(&theme, "theme", themeDefault, themeUsage)
	flag.StringVar(&theme, "t", themeDefault, themeUsage+" (shorthand)")

	const (
		themeRepoDefault = "https://github.com/lgaida/mediumish-gohugo-theme.git"
		themeRepoUsage   = "The URL for the repo of the theme you want to use (for hugo)"
	)
	flag.StringVar(&themeRepo, "theme-repo", themeRepoDefault, themeRepoUsage)
	flag.StringVar(&themeRepo, "u", themeRepoDefault, themeRepoUsage+" (shorthand)")

	const (
		deployDefault = ""
		deployUsage   = "The deployment target, e.g. s3"
	)
	flag.StringVar(&deploy, "deploy", deployDefault, deployUsage)
	flag.StringVar(&deploy, "d", deployDefault, deployUsage+" (shorthand)")

	flag.Parse()

	stack = strings.ToLower(stack)
	theme = strings.ToLower(theme)
	deploy = strings.ToLower(deploy)

	fmt.Println("Name: " + projName)
	fmt.Println("Stack: " + stack)

	absPath, metaData := sinit.CreateProject(projName)

	switch stack {
	case "node":
		sinit.InitNode(absPath, metaData)
	case "hugo":
		sinit.InitHugo(absPath, theme, themeRepo, deploy, metaData)
	}

	fmt.Printf("\nProject Created. Just type `cd %v`", projName)
}
