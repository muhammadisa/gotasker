package deadline

import (
	"github.com/gocraft/dbr/v2"
	"github.com/muhammadisa/go-cron-service/cron/models"
	"log"
	"time"
)

// DoAction which related to deadline cron
func DoAction(sess *dbr.Session) {
	formatLayout := "2006-01-02"

	var err error
	var deadlines []models.Deadline

	//Date(2020, 11, 05, 0, 0, 0, 0, time.Local)
	//Add(time.Hour * 24 * 1 * time.Duration(-deadline.DaysRemaining))
	//currentTime := time.
	//	Now().Add(time.Hour * 24 * 1 * time.Duration(3)).Format(formatLayout)
	//log.Println(currentTime)

	_, err = sess.Select("*").
		From("deadlines").
		Load(&deadlines)
	if err != nil {
		log.Fatal(err)
	}

	for _, deadline := range deadlines {
		ctAddByDaysRemaining := time.
			Date(2020, 11, 20, 0, 0, 0, 0, time.Local).
			AddDate(0, 0, int(deadline.DaysRemaining)).
			Format(formatLayout)

		var deadlinesFiltered []models.Deadline
		_, err = sess.Select("*").
			From("deadlines").
			Where("deadline = ? AND warned = ?", ctAddByDaysRemaining, false).
			Load(&deadlinesFiltered)
		if err != nil {
			log.Println(err)
		}

		if len(deadlinesFiltered) != 0 {
			log.Println(ctAddByDaysRemaining, deadlinesFiltered)
			deadlinesFiltered = nil
		}else{
			log.Println(ctAddByDaysRemaining)
		}
	}
}
