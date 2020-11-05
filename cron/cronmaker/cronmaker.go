package cronmaker

import (
	"github.com/gocraft/dbr/v2"
	"github.com/jasonlvhit/gocron"
	"github.com/muhammadisa/go-cron-service/cron/app/deadline"
	"github.com/muhammadisa/go-cron-service/cron/app/flush"
	"log"
)

type cronJob struct {
	Session *dbr.Session
}

type ICronJob interface {
	StartCronJobs()
}

func InitCrons(session *dbr.Session) ICronJob {
	return &cronJob{
		Session: session,
	}
}

func (cj cronJob) StartCronJobs() {
	log.Println("Merging crons and start them")
	cj.initDeadlineCron()
	//cj.initFlushCron()
	<-gocron.Start()
}

func (cj *cronJob) initDeadlineCron() {
	_ = gocron.Every(1).
		Seconds().
		Do(func() {
			deadline.DoAction(cj.Session)
		})
}

func (cj *cronJob) initFlushCron() {
	_ = gocron.Every(1).
		Seconds().
		Do(func() {
			flush.DoAction(cj.Session)
		})
}
