package server

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/nicosrgh/straw-hat/config"
	"github.com/robfig/cron"
)

// Init ...
func Init() {
	c := cron.New()

	// store := repository.Init()
	c.AddFunc(fmt.Sprintf("%s %s", config.C.ScheduleEvery, config.C.ScheduleTime), func() {
		// run the schedule
	})

	c.Start()

	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
