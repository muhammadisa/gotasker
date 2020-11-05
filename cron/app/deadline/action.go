package deadline

import (
	"github.com/gocraft/dbr/v2"
	"github.com/muhammadisa/go-cron-service/cron/models"
	"log"
	"time"
)

// DoAction which related to deadline cron
func DoAction(sess *dbr.Session) {
	defer log.Println("execute completed")
	formatLayout := "2006-01-02"

	var err error
	var daysRemaining []int
	var deadlinesFiltered []models.Deadline

	//Date(2020, 11, 05, 0, 0, 0, 0, time.Local)
	//Add(time.Hour * 24 * 1 * time.Duration(-deadline.DaysRemaining))
	//currentTime := time.
	//	Now().Add(time.Hour * 24 * 1 * time.Duration(3)).Format(formatLayout)
	//log.Println(currentTime)

	// Retrieve all deadlines
	_, err = sess.Select("days_remaining").
		From("deadlines").
		Load(&daysRemaining)
	if err != nil {
		log.Fatal(err)
	}

	// Scanning all deadlines
	for indexDays, day := range daysRemaining {
		// Add current time with days remaining
		ctAddByDaysRemaining := time.
			Date(2020, 11, 20, 0, 0, 0, 0, time.Local).
			AddDate(0, 0, day).
			Format(formatLayout)
		if indexDays != len(daysRemaining)-1 {
			// retrieve deadline wants tobe warned
			_, err = sess.Select("*").
				From("deadlines").
				Where("deadline = ? AND warned = ?", ctAddByDaysRemaining, false).
				Load(&deadlinesFiltered)
			if err != nil {
				log.Println(err)
			}
		}else{
			break
		}
	}

	for index, deadline := range deadlinesFiltered {
		log.Println("Publish Message", index, deadline)
	}
}
