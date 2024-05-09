package content

import (
	"byungflix-backend/database/components"
	"encoding/json"
	"net/http"
	"strconv"
)

type removeResponse struct {
	Status string `json:"status"`
}

func RemoveSeries(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response := removeResponse{
		Status: "fail",
	}
	seriesName := r.FormValue("series")

	if components.GetVideoListBySeriesTitle(seriesName) != nil {
		response.Status = "fail"
		responseJSON, _ := json.Marshal(response)
		rw.Write(responseJSON)
		return
	}

	components.RemoveSeries(seriesName)

	response.Status = "success"
	responseJSON, _ := json.Marshal(response)
	rw.Write(responseJSON)
	return
}

func RemoveVideo(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response := removeResponse{
		Status: "fail",
	}
	seriesName := r.FormValue("series")
	episodeCount, _ := strconv.Atoi(r.FormValue("episode"))
	components.RemoveVideo(seriesName, episodeCount)

	response.Status = "success"
	responseJSON, _ := json.Marshal(response)
	rw.Write(responseJSON)
	return
}
