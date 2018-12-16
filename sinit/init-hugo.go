package sinit

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"./templates"
)

/*
InitHugo initializes a hugo static site
*/
func InitHugo(absPath string, theme string, themeRepo string, deploy string, metaData MetaData) {
	fmt.Println("Initializing Hugo Static Site.")

	_, err := runCmd("hugo", "new", "site", absPath, "--force")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Installing theme %v from %v\n", theme, themeRepo)
	_, err = runCmdFromDir(absPath, "git", "submodule", "add", themeRepo, "themes/"+theme)
	if err != nil {
		log.Fatal(err)
	}

	// Add theme to config.toml
	f, err := os.OpenFile(path.Join(absPath, "config.toml"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteString("theme = \"" + theme + "\"")
	if err != nil {
		log.Fatal(err)
	}

	f.Close()

	fmt.Println("Setting up CI Pipeline")
	os.Mkdir(path.Join(absPath, ".circleci"), 0777)

	err = ioutil.WriteFile(path.Join(absPath, ".circleci", "config.yml"), []byte(templates.CircleHugo()), 0666)
	if err != nil {
		log.Fatal(err)
	}

	if len(deploy) > 0 {
		switch deploy {
		case "s3":
			deployInfo := DeployInfo{
				Flags:       "--recursive",
				PersistPath: "hugo/public",
			}
			err = appendTemplateToFile(templates.CircleDeployS3(), path.Join(absPath, ".circleci", "config.yml"), deployInfo)
			if err != nil {
				log.Fatal(err)
			}
		case "gcs":
			deployInfo := DeployInfo{
				Flags:       "-r",
				PersistPath: "hugo/public",
			}
			err = appendTemplateToFile(templates.CircleDeployGCS(), path.Join(absPath, ".circleci", "config.yml"), deployInfo)
			if err != nil {
				log.Fatal(err)
			}
		case "firebase":
			deployInfo := DeployInfo{
				PersistPath: "hugo",
			}
			err = appendTemplateToFile(templates.CircleDeployFirebase(), path.Join(absPath, ".circleci", "config.yml"), deployInfo)
			if err != nil {
				log.Fatal(err)
			}

			err = runInteractiveCmdFromDir(absPath, "firebase", "init")
			if err != nil {
				log.Fatal(err)
			}
		}

	}

}
