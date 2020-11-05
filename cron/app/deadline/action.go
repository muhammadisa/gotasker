package deadline

import (
	"github.com/gocraft/dbr/v2"
	"github.com/muhammadisa/go-cron-service/cron/models"
	"log"
	"time"
)

func currentTimeAddByString(days int) string {
	return time.Date(
		2020, 11, 16,
		0, 0, 0, 0, time.Local).
		AddDate(0, 0, days).
		Format("2006-01-02")
}

func retrieveDaysRemaining(sess *dbr.Session) []int {
	var days []int
	_, err := sess.Select("days_remaining").
		From("deadlines").
		Load(&days)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return days
}

func retrieveDeadlines(sess *dbr.Session, day int, deadlinesFiltered *[]models.Deadline) {
	_, err := sess.Select("*").
		From("deadlines").
		Where("deadline = ? AND warned = ?", currentTimeAddByString(day), false).
		Load(deadlinesFiltered)
	if err != nil {
		log.Println(err)
	}
}

func extractDeadlines(sess *dbr.Session, daysRemaining []int) []models.Deadline {
	var deadlinesFiltered []models.Deadline
	for i := 1; i <= len(daysRemaining)-1; i++ {
		retrieveDeadlines(sess, daysRemaining[i], &deadlinesFiltered)
	}
	return deadlinesFiltered
}

// DoAction which related to deadline cron
func DoAction(sess *dbr.Session) {
	defer log.Println("execute completed")

	daysRemaining := retrieveDaysRemaining(sess)
	deadlinesFiltered := extractDeadlines(sess, daysRemaining)

	for index, deadline := range deadlinesFiltered {
		log.Println("Publish Message", index, deadline)
	}
}
