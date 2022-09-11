package content

import (
	"github.com/gorilla/mux"
	"net/http"
)

func StreamVideoAndSubtitle(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contentTitle := vars["content_title"]
	episodeCount := vars["episode_count"]
	videoName := vars["video_name"]
	filePath := "contents/" + contentTitle + "/" + episodeCount + "/" + videoName
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(rw, r, filePath)
}
