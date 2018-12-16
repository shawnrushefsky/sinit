package sinit

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
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

	templateDir, err := filepath.Abs("templates")
	if err != nil {
		log.Fatal(err)
	}

	err = copy(path.Join(templateDir, "circle-hugo.yml"), path.Join(absPath, ".circleci", "config.yml"))
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
			err = appendTemplateToFile("circle-deploy-s3.yml", path.Join(absPath, ".circleci", "config.yml"), deployInfo)
			if err != nil {
				log.Fatal(err)
			}
		case "gcs":
			deployInfo := DeployInfo{
				Flags:       "-r",
				PersistPath: "hugo/public",
			}
			err = appendTemplateToFile("circle-deploy-gcs.yml", path.Join(absPath, ".circleci", "config.yml"), deployInfo)
			if err != nil {
				log.Fatal(err)
			}
		}

	}

}
