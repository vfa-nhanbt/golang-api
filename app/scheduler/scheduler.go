package scheduler

import (
	"fmt"

	"github.com/go-co-op/gocron/v2"
)

var cronScheduler gocron.Scheduler

func CreateScheduler() error {
	s, err := gocron.NewScheduler()
	if err != nil {
		return err
	}
	cronScheduler = s
	return nil
}

func StartScheduler() {
	cronScheduler.Start()
}

func ShutdownScheduler() error {
	err := cronScheduler.Shutdown()
	return err
}

func StartDailyScheduler(interval uint, atTimes gocron.AtTimes, handler interface{}) error {
	job, err := cronScheduler.NewJob(
		gocron.DailyJob(
			interval,
			atTimes,
		),
		gocron.NewTask(handler),
	)
	if err != nil {
		return err
	}
	fmt.Println("Daily Job created:", job)
	return nil
}

func StartAppScheduler(job *SendNotificationJob) error {
	err := CreateScheduler()
	if err != nil {
		return err
	}

	/// Send daily notification to user to recommend a book for today at 9:00 AM
	err = StartDailyScheduler(1, gocron.NewAtTimes(
		gocron.NewAtTime(9, 0, 0),
	), job.SendBookNotificationToUser)
	if err != nil {
		return err
	}

	StartScheduler()
	return nil
}
