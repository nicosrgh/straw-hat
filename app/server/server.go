package server

import (
	"fmt"
	"os"
	"os/signal"

	transformer "github.com/nicosrgh/straw-hat/app/transformer/dimension"
	"github.com/nicosrgh/straw-hat/config"
	"github.com/nicosrgh/straw-hat/config/repository"
	"github.com/robfig/cron"
)

// Init ...
func Init() {
	c := cron.New()

	store := repository.Init()
	c.AddFunc(fmt.Sprintf("%s %s", config.C.ScheduleEvery, config.C.ScheduleTime), func() {
		// run the schedule

		// dimension
		transformer.GenderDimension(store)
		transformer.TitleDimension(store)
	})

	c.Start()

	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
