package calculatecalendar

import (
	"encoding/json"
	"net/http"

	"github.com/recokioshi/orangeWorkSchedule-back/model"
)

func calculate(calculationInput model.CalculationInput) model.CalculationInput {
	return calculationInput
}

// IndexHandler is base functions that handles root post request
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	valid, calculationInput := model.GetCalculationInput(w, r)
	calculationOutput := calculate(calculationInput)
	if valid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(calculationOutput)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
