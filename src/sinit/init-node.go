package sinit

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
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
	createFileFromTemplate("package-json.gotxt", path.Join(absPath, "package.json"), pInfo)

	_, err = runCmdFromDir(absPath, "npm", "install")
	if err != nil {
		log.Fatal(err)
	}
}
