package main

import (
	"time"
	"errors"
)

func timeDiff(u_start int64, u_end int64) time.Duration {
	start := time.Unix(u_start, 0)
	end := time.Unix(u_end, 0)
	return end.Sub(start)
}

func startTimer(appHome string, time_ time.Time) error {
	app, err := getApp(appHome)
	if err != nil {
		return err
	}

	timers := app.Timer
	if len(timers) > 0 {
		lastTimer := timers[len(timers) - 1:][0]
		if lastTimer.End == 0 {
			return errors.New("The timer is already started.")
		}
	}

	newTimer := Timer{
		Start: time_.Unix(),
	}
	timers = append(timers, newTimer)
	app.Timer = timers

	setApp(appHome, app)
	return nil
}

func endTimer(appHome string, time_ time.Time) (time.Duration, error) {
	app, err := getApp(appHome)
	if err != nil {
		return 0, err
	}

	timers := app.Timer

	if len(timers) <= 0 {
		return 0, errors.New("A timer has not started.")
	}

	lastTimer := &timers[len(timers) - 1:][0]
	if lastTimer.End != 0 {
		return 0, errors.New("The timer is already ended.")
	}

	lastTimer.End = time_.Unix()

	app.Timer = timers

	setApp(appHome, app)
	return timeDiff(lastTimer.Start, lastTimer.End), nil
}