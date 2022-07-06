package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is", id)

	movie, err := app.models.DB.Get(id)

	app.logger.Println("Movie is", movie)
	app.logger.Println("Err is", err)

	// movie := models.Movie{
	// 	ID:          id,
	// 	Title:       "Some movie",
	// 	Description: "Some description",
	// 	Year:        2021,
	// 	ReleaseDate: time.Date(2021, 01, 01, 01, 0, 0, 0, time.Local),
	// 	Runtime:     100,
	// 	Rating:      5,
	// 	MPAARating:  "PG-13",
	// 	CreatedAt:   time.Now(),
	// 	UpdatedAt:   time.Now(),
	// }

	app.writeJSON(w, http.StatusOK, movie, "movie")

}

func (app *application) errorJSON(w http.ResponseWriter, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	app.writeJSON(w, http.StatusBadRequest, theError, "error")
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.ALL()

	if err != nil {
		app.errorJSON(w, err)
		return
	}
	app.writeJSON(w, http.StatusOK, movies, "movies")
}
