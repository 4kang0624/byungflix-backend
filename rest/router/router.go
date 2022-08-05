package router

import (
	"byungflix-backend/rest/upload"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/api/upload/video", upload.UploadVideo).Methods("POST")
	router.HandleFunc("/api/create/series", upload.MakeSeries).Methods("POST")
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
