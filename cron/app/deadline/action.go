package deadline

import (
	"github.com/gocraft/dbr/v2"
	"github.com/muhammadisa/go-cron-service/cron/models"
	"log"
	"time"
)

func currentTimeAddByString(days int) string {
	return time.Date(
		2020, 11, 5,
		0, 0, 0, 0, time.Local).
		AddDate(0, 0, days).
		Format("2006-01-02")
}

func dates(sess *dbr.Session) []string {
	var days []int
	var dates []string
	_, err := sess.Select("days_remaining").
		From("deadlines").
		Load(&days)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	for _, day := range days {
		dates = append(dates, currentTimeAddByString(day))
	}
	return dates
}

func deadlines(sess *dbr.Session, date []string) []models.Deadline {
	var deadlines []models.Deadline
	_, err := sess.Select("*").
		From("deadlines").
		Where("deadline IN ? AND warned = ?", date, false).
		Load(&deadlines)
	if err != nil {
		log.Println(err)
	}
	return deadlines
}

// DoAction which related to deadline cron
func DoAction(sess *dbr.Session) {
	defer log.Println("execute completed")
	dates := dates(sess)
	deadlines := deadlines(sess, dates)
	for index, deadline := range deadlines {
		log.Println("Publish Message", index, deadline)
	}
}
