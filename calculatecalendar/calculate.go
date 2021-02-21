package calculatecalendar

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// IndexHandler is base functions that handles root post request
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s", reqBody)
}
