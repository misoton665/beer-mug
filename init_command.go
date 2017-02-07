package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli"
	"time"
	"encoding/json"
	"errors"
)

/* Command:
 * init - create directory and file for app.
 * init look - show app directory location and its file content.
 */
func initCommand() cli.Command {
	return cli.Command{
		Name: "init",
		Action: runInit,
		Subcommands: []cli.Command{
			initLookCommand(),
			initForceCommand(),
		},
	}
}

func initLookCommand() cli.Command {
	return cli.Command{
		Name: "look",
		Action: runInitLook,
	}
}

func initForceCommand() cli.Command {
	return cli.Command{
		Name: "force",
		Action: runInitForce,
	}
}

// init look
func runInitLook(_ *cli.Context) error {
	appHome, err := getAppHome()
	if err != nil {
		return err
	}

	var fileExist bool
	if fileExist = isExistApp(appHome); !fileExist {
		return errors.New("The application has not initilazed.")
	}

	app, err_ := getApp(appHome)
	if err_ != nil {
		return err_
	}

	var appBytes []byte
	appBytes, err = json.Marshal(app)
	if err != nil {
		return err
	}

	fmt.Printf("home: %s\napp: %s\n", appHome, string(appBytes))
	return nil
}

func runInitForce(_ *cli.Context) error {
	appHome, err := getAppHome()
	if err != nil {
		return err
	}

	if err = os.Mkdir(appHome, os.FileMode(0755)); err != nil {
		fmt.Println(err)
	}

	app := App{
		Timer: []Timer{},
	}

	if err := setApp(appHome, app); err != nil {
		return err
	}

	fmt.Println("Force initialized.")

	return nil
}

// init
func runInit(_ *cli.Context) error {
	appHome, err := getAppHome()
	if err != nil {
		return err
	}

	if fileExist := isExistApp(appHome); fileExist {
		return errors.New("The application has already initilazed.")
	}

	if err = os.Mkdir(appHome, os.FileMode(0755)); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("mkdir to %s for app home.\n", appHome)

	app := App{
		Timer: []Timer{
			{Start: time.Now().Unix()},
		},
	}

	if err := setApp(appHome, app); err != nil {
		return err
	}

	return nil
}