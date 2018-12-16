package sinit

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"./templates"
)

type packageInfo struct {
	Name          string
	Author        string
	NodeVersion   string
	ChaiVersion   string
	MochaVersion  string
	ESLintVersion string
	NYCVersion    string
}

/*
InitNode Initializes a Node.JS project, including testing framework and linter
*/
func InitNode(absPath string, metaData MetaData) {
	fmt.Println("Initiating Node.JS project")

	// Create the new directory
	fmt.Println("Making Directories")
	os.Mkdir(path.Join(absPath, "src"), 0777)
	os.Mkdir(path.Join(absPath, "test"), 0777)
	os.Mkdir(path.Join(absPath, ".circleci"), 0777)

	// Get node version
	nodev, err := runCmd("node", "-v")
	if err != nil {
		log.Fatal(err)
	}

	// Write .nvmrc
	ioutil.WriteFile(path.Join(absPath, ".nvmrc"), []byte(nodev), 0666)

	// node version comes out like v10.9.0, so we need to trim the leading 'v'
	nodev = nodev[1:]

	fmt.Println("Getting Dependency Versions")
	chaiv, err := runCmd("npm", "show", "chai", "version")
	if err != nil {
		log.Fatal(err)
	}

	mochav, err := runCmd("npm", "show", "mocha", "version")
	if err != nil {
		log.Fatal(err)
	}

	eslintv, err := runCmd("npm", "show", "eslint", "version")
	if err != nil {
		log.Fatal(err)
	}

	nycv, err := runCmd("npm", "show", "nyc", "version")
	if err != nil {
		log.Fatal(err)
	}

	pInfo := packageInfo{
		Name:          metaData.Name,
		Author:        metaData.Author,
		NodeVersion:   nodev,
		ChaiVersion:   chaiv,
		MochaVersion:  mochav,
		ESLintVersion: eslintv,
		NYCVersion:    nycv,
	}

	// Initiate the project
	err = createFileFromTemplate(templates.PackageJSON(), path.Join(absPath, "package.json"), pInfo)
	if err != nil {
		log.Fatal(err)
	}

	newFile, err := os.Create(path.Join(absPath, "src", "index.js"))
	if err != nil {
		log.Fatal(err)
	}
	newFile.Close()

	fmt.Println("Installing dependencies")
	_, err = runCmdFromDir(absPath, "npm", "install")
	if err != nil {
		log.Fatal(err)
	}

	templateDir, err := filepath.Abs("templates")
	if err != nil {
		log.Fatal(err)
	}

	err = copy(path.Join(templateDir, ".eslintrc.js"), path.Join(absPath, ".eslintrc.js"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Setting up CircleCI")
	err = createFileFromTemplate(templates.CircleNode(), path.Join(absPath, ".circleci", "config.yml"), pInfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Fetching .gitignore")
	resp, err := http.Get("https://raw.githubusercontent.com/github/gitignore/master/Node.gitignore")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(path.Join(absPath, ".gitignore"), body, 0666)

	err = createFileFromTemplate("compose-node.yml", path.Join(absPath, "docker-compose.yml"), pInfo)
	if err != nil {
		log.Fatal(err)
	}
}
