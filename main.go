package main

import (
	"flag"
	"fmt"
	"strings"

	"./src/sinit"
)

func main() {
	var (
		projName string
		stack    string
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

	flag.Parse()

	stack = strings.ToLower(stack)

	fmt.Println("Name: " + projName)
	fmt.Println("Stack: " + stack)

	absPath, metaData := sinit.CreateProject(projName)

	if stack == "node" {
		sinit.InitNode(absPath, metaData)
	}

	fmt.Printf("\nProject Created. Just type cd %v", projName)
}
