package content

import (
	"byungflix-backend/database"
	"byungflix-backend/database/components"
	"encoding/json"
	"net/http"
)

type SearchSeriesResult struct {
	Status string            `json:"status"`
	Result []database.Series `json:"result"`
}

func SearchSeries(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	var response = SearchSeriesResult{
		Status: "fail",
		Result: nil,
	}

	keyword := r.FormValue("keyword")

	seriesList := components.GetSeriesList(keyword)

	response.Status = "success"
	response.Result = seriesList

	responseJSON, _ := json.Marshal(response)
	rw.Write(responseJSON)
	return
}
