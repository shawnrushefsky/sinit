package sinit

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"text/template"
)

func runCmd(cmd string, args ...string) (output string, err error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return runCmdFromDir(dir, cmd, args...)
}

func runCmdFromDir(dir string, cmd string, args ...string) (output string, err error) {
	command := exec.Command(cmd, args...)
	command.Dir = dir
	var out bytes.Buffer
	command.Stdout = &out
	err = command.Run()
	output = out.String()
	if len(output) > 0 && output[len(output)-1] == '\n' {
		output = output[:len(output)-1]
	}
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

func copy(source, destination string) error {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(destination, input, 0666)
	if err != nil {
		return err
	}

	return nil
}
