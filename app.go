package main

import (
	"io/ioutil"
	"encoding/json"
	"os"
	"errors"
)

type (
	App struct {
		Timer []Timer `json:"timer"`
	}

	Timer struct {
		Start int64 `json:"start"`
		End int64 `json:"end"`
	}
)

const (
	AppFileName = "app.json"
	AppDirName = ".beer-mug"
)

func getAppHome() (string, error) {
	home := os.Getenv("HOME")

	if len(home) <= 0 {
		return "", errors.New("$HOME is empty")
	}

	last_char := home[len(home) - 1:]
	if last_char == "/" {
		home = home[: len(home) - 2]
	}
	return home + "/" + AppDirName, nil
}

func isExistApp(appHome string) bool {
	_, err := os.Stat(appHome + "/" + AppFileName)
	return err == nil
}

func setApp(appHome string, app App) error {
	b, err := json.Marshal(app)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(appHome + "/" + AppFileName, b, os.FileMode(0755))
	return err
}

func getApp(appHome string) (App, error) {
	appFile, err := ioutil.ReadFile(appHome + "/" + AppFileName)
	if err != nil {
		return App{}, err
	}

	app := new(App)

	if err := json.Unmarshal(appFile, app); err != nil {
		return App{}, err
	}

	return *app, nil
}