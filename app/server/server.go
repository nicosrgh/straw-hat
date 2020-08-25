package server

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/nicosrgh/straw-hat/app/transformer"
	"github.com/nicosrgh/straw-hat/config"
	"github.com/robfig/cron"
)

// Init ...
func Init() {
	c := cron.New()

	// store := repository.Init()
	c.AddFunc(fmt.Sprintf("%s %s", config.C.ScheduleEvery, config.C.ScheduleTime), func() {
		// run the schedule

		// dimension employee
		transformer.TitleDimension()
		transformer.GenderDimension()
		transformer.LocationDimension()
		transformer.DepartmentDimension()
		transformer.StatusDimension()

		// dimension transaction
		transformer.ProductDimension()
		transformer.PartnerDimension()
		transformer.ClientDimension()

		// source
		// transformer.EmployeeSource()

		// fact
		// transformer.FactEmployeeLocation()

		// datamart
		// transformer.EmployeeTitleDatamart()
		// transformer.EmployeeTotalDatamart()

	})

	c.Start()

	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
