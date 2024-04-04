package content

import (
	"byungflix-backend/database"
	"byungflix-backend/database/components"
	"encoding/json"
	"net/http"
)

type UpdateSeriesResponse struct {
	Status string          `json:"status"`
	Result database.Series `json:"result"`
}

func UpdateSeries(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	response := UpdateSeriesResponse{
		Status: "success",
		Result: database.Series{},
	}

	oldTitle := r.FormValue("old_title")
	title := r.FormValue("title")
	titlekor := r.FormValue("titlekor")
	description := r.FormValue("description")

	series := database.Series{
		Title:       title,
		TitleKor:    titlekor,
		Description: description,
	}

	components.UpdadeSeries(oldTitle, series)

	responseJSON, _ := json.Marshal(response)
	rw.Write(responseJSON)
}
