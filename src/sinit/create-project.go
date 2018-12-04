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
	Email  string
}

func runCmd(cmd string, args ...string) (output string, err error) {
	command := exec.Command(cmd, args...)
	var out bytes.Buffer
	command.Stdout = &out
	err = command.Run()
	output = out.String()
	return
}

func createFileFromTemplate(templateFile string, ouputFilename string, data interface{}) error {
	templateDir, err := filepath.Abs("templates")
	if err != nil {
		return err
	}

	tmpl, err := template.New(templateFile).ParseFiles(path.Join(templateDir, templateFile))
	if err != nil {
		return err
	}

	newFile, err := os.Create(ouputFilename)
	if err != nil {
		return err
	}

	err = tmpl.Execute(newFile, data)
	if err != nil {
		return err
	}

	return nil
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
