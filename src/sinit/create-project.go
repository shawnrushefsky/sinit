package sinit

import (
	"fmt"
	"log"
	"os"
	"path"
)

/*
MetaData contains information about the project and its author
*/
type MetaData struct {
	Name   string
	Author string
	Email  string
}

/*
CreateProject creates the directory for a project, and sets up universal boilerplate
*/
func CreateProject(name string) (absPath string, meta MetaData) {
	// Get git user name and email
	author, err := runCmd("git", "config", "user.name")
	if err != nil {
		log.Fatal(err)
	}
	email, err := runCmd("git", "config", "user.email")
	if err != nil {
		log.Fatal(err)
	}

	meta = MetaData{name, author, email}

	// Get an absolute path for the new project
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	absPath = path.Join(dir, name)

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
	err = createFileFromTemplate("readme.md", path.Join(absPath, "README.md"), meta)
	if err != nil {
		log.Fatal(err)
	}

	// Create a changelog
	fmt.Println("Creating CHANGELOG")
	err = createFileFromTemplate("changelog.md", path.Join(absPath, "CHANGELOG.md"), meta)
	if err != nil {
		log.Fatal(err)
	}

	// Create a codeowners
	fmt.Println("Creating CODEOWNERS")
	err = createFileFromTemplate("codeowners", path.Join(absPath, "CODEOWNERS"), meta)
	if err != nil {
		log.Fatal(err)
	}

	return
}
