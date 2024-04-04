package router

import (
	"byungflix-backend/rest/content"
	"byungflix-backend/rest/upload"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/api/upload/video", upload.UploadVideo).Methods("POST")
	router.HandleFunc("/api/upload/subtitle", upload.UploadSubtitle).Methods("POST")
	router.HandleFunc("/api/create/series", upload.CreateSeries).Methods("POST")
	router.HandleFunc("/api/search/series", content.SearchSeries).Methods("POST")
	router.HandleFunc("/api/update/series", content.UpdateSeries).Methods("POST")
	router.HandleFunc("/api/search/video_by_series", content.SearchVideoInSeries).Methods("POST")
	router.HandleFunc("/contents/{content_title}/{episode_count}/{video_name}", content.StreamVideoAndSubtitle).Methods("GET")
	router.HandleFunc("/download/{content_title}/{episode_count}/{video_name}", content.DownloadVideo).Methods("GET")
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
