package calculatecalendar

import (
	"net/http"

	"github.com/recokioshi/orangeWorkSchedule-back/model"
)

// IndexHandler is base functions that handles root post request
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	valid, _ := model.GetCalculationInput(w, r)
	if valid {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
