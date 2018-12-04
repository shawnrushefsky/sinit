package sinit

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"text/template"
)

type metaData struct {
	Name   string
	Author string
}

/*
CreateProject creates the directory for a project, and sets up universal boilerplate
*/
func CreateProject(name string) {
	// Get the absolute path for the templates we're going to use
	templateDir, err := filepath.Abs("templates")
	if err != nil {
		log.Fatal(err)
	}

	meta := metaData{name, "Someone Special"}
	// Get an absolute path for the new project
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	absPath := path.Join(dir, name)

	// Create the new directory
	os.Mkdir(absPath, 0777)

	// Initialize a repo in the directory
	gitInit := exec.Command("git", "init", absPath)
	var out bytes.Buffer
	gitInit.Stdout = &out
	err = gitInit.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())

	readme, err := template.New("readme.gotxt").ParseFiles(path.Join(templateDir, "readme.gotxt"))
	if err != nil {
		log.Fatal(err)
	}

	// Create a readme
	readmeFile, err := os.Create(path.Join(absPath, "README.md"))
	if err != nil {
		log.Fatal(err)
	}

	err = readme.Execute(readmeFile, meta)
	if err != nil {
		log.Fatal(err)
	}

	// Create a changelog

	// Create a codeowners
}
