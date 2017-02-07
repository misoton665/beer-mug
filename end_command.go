package main

import (
	"fmt"
	"github.com/urfave/cli"
	"time"
)

func endCommand() cli.Command {
	return cli.Command{
		Name:   "end",
		Action: runEnd,
	}
}

func beerLine() string {
	return fourBeers() + " " + fourBeers() + " " + fourBeers() + " " + fourBeers()
}

func fourBeers() string {
	return twoBeers() + " " + twoBeers()
}

func twoBeers() string {
	return "\U0001F37A \U0001F37A"
}

func runEnd(*cli.Context) error {
	appHome, err := getAppHome()
	if err != nil {
		return err
	}

	workingTime, err_ := endTimer(appHome, time.Now())
	if err_ != nil {
		return err_
	}

	allsec := int(workingTime.Seconds())

	hours := allsec / 60 / 60
	minutes := allsec/60 - hours*60
	seconds := allsec - (minutes*60 - hours*60*60)

	fmt.Printf(beerLine()+"\n"+
		beerLine()+"\n"+
		twoBeers()+"                         "+twoBeers()+"\n"+
		twoBeers()+"  YOU WORKED %03d:%02d:%02d!! "+twoBeers()+"\n"+
		twoBeers()+"    CONGRATULATIONS!!!   "+twoBeers()+"\n"+
		twoBeers()+"                         "+twoBeers()+"\n"+
		beerLine()+"\n"+
		beerLine()+"\n",
		hours, minutes, seconds)

	return nil
}
