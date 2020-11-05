package flush

import (
	"fmt"
	"github.com/gocraft/dbr/v2"
)

// DoAction which related to flush cron
func DoAction(sess *dbr.Session) {
	fmt.Println("Flush")
}
