package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	curentStatus := AppStatus{
		Status:      "Availablesss",
		Environment: app.config.env,
		Version:     version,
	}

	js, err := json.MarshalIndent(curentStatus, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)

}
