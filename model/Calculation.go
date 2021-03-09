package model

import (
	"errors"
	"log"
	"net/http"
)

// Shift is representation of working hours of particular worker
type Shift struct {
	StartTime int
	EndTime   int
}

// Day is an object that holds all data related to one day. It defines if day is free,
// what workers should or can't be assigned to, and holds shifts for each worker.
type Day struct {
	Whitelist []int
	Blacklist []int
	Active    bool
	Shifts    map[int]Shift
}

// CalculationInput is on input object got from http request.
// It contains workers count and list of Days
type CalculationInput struct {
	Workers  int
	Calendar []Day
}

func GetCalculationInput(w http.ResponseWriter, r *http.Request) (bool, CalculationInput) {
	var ci CalculationInput

	err := decodeJSONBody(w, r, &ci)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return false, CalculationInput{}
	}

	return true, ci

}
