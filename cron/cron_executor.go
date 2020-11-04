package cron

import (
	"github.com/muhammadisa/go-cron-service/cron/cronmaker"
)

// Run start run cron scheduler
func Run() {
	cronmaker.
		InitCrons().
		StartCronJobs()
}
