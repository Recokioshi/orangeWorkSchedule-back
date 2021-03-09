package calculatecalendar

import (
	"fmt"
	"net/http"

	"github.com/recokioshi/orangeWorkSchedule-back/model"
)

// IndexHandler is base functions that handles root post request
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	valid, calendarInput := model.GetCalculationInput(w, r)

	if valid {
		fmt.Println("calendarInput: ", calendarInput)
	} else {
		fmt.Println("bad request")
	}
}
