package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data
	js, err := json.Marshal(wrapper)
	fmt.Println(err)
	// if err != nil {
	// 	return err
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	// return nil
}
