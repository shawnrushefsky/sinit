package sinit

import (
	"fmt"
	"log"
	"os"
	"path"
)

type metaData struct {
	Name   string
	Author string
	Email  string
}

/*
CreateProject creates the directory for a project, and sets up universal boilerplate
*/
func CreateProject(name string) {
	// Get git user name and email
	author, err := runCmd("git", "config", "user.name")
	if err != nil {
		log.Fatal(err)
	}
	email, err := runCmd("git", "config", "user.email")
	if err != nil {
		log.Fatal(err)
	}

	meta := metaData{name, author, email}

	// Get an absolute path for the new project
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	absPath := path.Join(dir, name)

	// Create the new directory
	fmt.Println("Making Directory")
	os.Mkdir(absPath, 0777)

	// Initialize a repo in the directory
	result, err := runCmd("git", "init", absPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// Create a readme
	fmt.Println("Creating README")
	err = createFileFromTemplate("readme.gotxt", path.Join(absPath, "README.md"), meta)
	if err != nil {
		log.Fatal(err)
	}

	// Create a changelog
	fmt.Println("Creating CHANGELOG")
	err = createFileFromTemplate("changelog.gotxt", path.Join(absPath, "CHANGELOG.md"), meta)
	if err != nil {
		log.Fatal(err)
	}

	// Create a codeowners
	fmt.Println("Creating CODEOWNERS")
	err = createFileFromTemplate("codeowners.gotxt", path.Join(absPath, "CODEOWNERS"), meta)
	if err != nil {
		log.Fatal(err)
	}
}
