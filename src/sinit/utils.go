package sinit

import (
	"bytes"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"text/template"
)

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
