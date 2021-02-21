package main

import (
	"github.com/recokioshi/orangeWorkSchedule-back/calculatecalendar"
	"github.com/recokioshi/orangeWorkSchedule-back/router"
)

func main() {
	router.RootRouter(calculatecalendar.IndexHandler)
}
