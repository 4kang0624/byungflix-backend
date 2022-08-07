package upload

import (
	"byungflix-backend/database"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type uploadVideoResponse struct {
	Status                string `json:"status"`
	OriginalVideoTitle    string `json:"original_video_size"`
	OriginalSubtitleTitle string `json:"original_subtitle_size"`
	OriginalVideoSize     int    `json:"video_size"`
	OriginalSubtitleSize  int    `json:"subtitle_size"`
	VideoPath             string `json:"video_path"`
	SubtitlePath          string `json:"subtitle_path"`
}

func MakeSeries(rw http.ResponseWriter, r *http.Request) {
	series := database.Series{
		Title:       r.FormValue("title"),
		TitleKor:    r.FormValue("title_kor"),
		Description: r.FormValue("description"),
	}
	os.Mkdir("contents/"+r.FormValue("title"), 0755)
	database.CreateSeries(series)
}

func UploadVideo(rw http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 * 1024 * 1024 * 1024) // Upload video size limit 10GB
	rw.Header().Set("Content-Type", "application/json")
	response := uploadVideoResponse{
		Status:                "fail",
		OriginalSubtitleTitle: "",
		OriginalVideoTitle:    "",
		OriginalSubtitleSize:  0,
		OriginalVideoSize:     0,
		VideoPath:             "",
		SubtitlePath:          "",
	}

	path := "contents/" + r.FormValue("content_title") + "/" + r.FormValue("episode_count")
	err := os.Mkdir(path, 0755)
	if err != nil {
		responseJSON, _ := json.Marshal(response)
		rw.Write(responseJSON)
		return
	}

	video, handlerVid, _ := r.FormFile("video")
	defer video.Close()

	subtitle, handlerSub, _ := r.FormFile("subtitle")
	defer subtitle.Close()

	tempVideo, _ := ioutil.TempFile(path, r.FormValue("content_title")+"_"+r.FormValue("episode_count")+"_*.mkv")
	defer tempVideo.Close()

	tempSub, _ := ioutil.TempFile(path, r.FormValue("content_title")+"_"+r.FormValue("episode_count")+"_*.ass")
	defer tempSub.Close()

	fileBytes, _ := ioutil.ReadAll(video)
	tempVideo.Write(fileBytes)

	fileBytes, _ = ioutil.ReadAll(subtitle)
	tempSub.Write(fileBytes)

	response = uploadVideoResponse{
		Status:                "success",
		OriginalVideoTitle:    handlerVid.Filename,
		OriginalSubtitleTitle: handlerSub.Filename,
		OriginalVideoSize:     int(handlerVid.Size),
		OriginalSubtitleSize:  int(handlerSub.Size),
		VideoPath:             tempVideo.Name(),
		SubtitlePath:          tempSub.Name(),
	}

	responseJSON, _ := json.Marshal(response)
	rw.Write(responseJSON)
}
