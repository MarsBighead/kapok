package util

import (
	"os"
	"path"
	"path/filepath"
)

//Application  path info or some configuration logic
// --configure file in directory of
//          1. Same to execute command line path: config.yaml
//          2. Not find with #1, check Aplication.Home/etc/application/default.yaml
//          3. Aplication.Home/etc/application
//          4. Aplication.Home/etc/application
///Executable binary package:
//          conifgure #1: Application repository/cmd/application/main.go
//          conifgure #2.#3,#4: Application home/bin/application
// --output file in directory
//          conifgure #1,#2: current path
//          conifgure #3,#4: Aplication.Home/var/application
type Application struct {
	Home      string
	Pwd       string
	Config    string
	IsDefault bool
}

//Traversing the application directory
func Traversing(name, config string) (*Application, error) {
	app := new(Application)
	current, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}

	app.Pwd = current

	ss := path.Base(ex)
	if ss == "main" || ss == "main.exe" {
		app.Home = current
	} else {
		app.Home = filepath.Dir(ex)
	}

	if config != "" && isFileExist(config) {
		app.Config = config
	} else {
		app.Config = current + "/config.yaml"
	}

	return app, nil

}

func isDirExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
func isFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
