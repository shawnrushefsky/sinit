package main

import (
	"flag"
	"fmt"

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

	fmt.Println("Name: " + projName)
	fmt.Println("Stack: " + stack)

	sinit.CreateProject(projName)
}
