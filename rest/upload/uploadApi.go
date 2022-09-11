package upload

import (
	"byungflix-backend/database"
	"byungflix-backend/util"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type uploadVideoResponse struct {
	Status             string `json:"status"`
	OriginalVideoTitle string `json:"original_video_size"`
	OriginalVideoSize  int    `json:"video_size"`
	VideoPath          string `json:"video_path"`
}

type uploadSubtitleResponse struct {
	Status                string `json:"status"`
	OriginalSubtitleTitle string `json:"original_subtitle_size"`
	OriginalSubtitleSize  int    `json:"subtitle_size"`
	Language              string `json:"language"`
	SubtitlePath          string `json:"subtitle_path"`
}

func MakeSeries(rw http.ResponseWriter, r *http.Request) {
	databaseInput := database.Series{
		Title:       r.FormValue("title"),
		TitleKor:    r.FormValue("title_kor"),
		Description: r.FormValue("description"),
	}
	os.Mkdir("contents/"+r.FormValue("title"), 0755)
	database.CreateSeries(databaseInput)
}

func UploadVideo(rw http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 * 1024 * 1024 * 1024) // Upload video size limit 10GB
	rw.Header().Set("Content-Type", "application/json")
	response := uploadVideoResponse{
		Status:             "fail",
		OriginalVideoTitle: "",
		OriginalVideoSize:  0,
		VideoPath:          "",
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

	tempVideo, _ := ioutil.TempFile(path, r.FormValue("content_title")+"_"+r.FormValue("episode_count")+"_*.mkv")
	defer tempVideo.Close()
	tempVidName := strings.Replace(tempVideo.Name(), "\\", "/", -1)

	fileBytes, _ := ioutil.ReadAll(video)
	tempVideo.Write(fileBytes)

	videoPathHls := util.EncodeMkvToHls(tempVidName)

	response = uploadVideoResponse{
		Status:             "success",
		OriginalVideoTitle: handlerVid.Filename,
		OriginalVideoSize:  int(handlerVid.Size),
		VideoPath:          tempVidName,
	}

	responseJSON, _ := json.Marshal(response)
	rw.Write(responseJSON)

	videoCnt, _ := strconv.Atoi(r.FormValue("episode_count"))
	databaseInput := database.Video{
		Title:        r.FormValue("content_title") + " - " + r.FormValue("episode_count"),
		ContentTitle: r.FormValue("content_title"),
		EpisodeCount: videoCnt,
		ReleaseDate:  r.FormValue("release_date"),
		UploadDate:   r.FormValue("upload_date"),
		VideoPath:    tempVidName,
		VideoPathHls: videoPathHls,
		SubtitlePath: map[string]string{},
	}
	database.CreateVideo(databaseInput)
}

func UploadSubtitle(rw http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(500 * 1024 * 1024) // Upload subtitle size limit 500MB
	rw.Header().Set("Content-Type", "application/json")
	response := uploadSubtitleResponse{
		Status:                "fail",
		OriginalSubtitleTitle: "",
		OriginalSubtitleSize:  0,
		Language:              "",
		SubtitlePath:          "",
	}

	path := "contents/" + r.FormValue("content_title") + "/" + r.FormValue("episode_count")

	subtitle, handlerSub, _ := r.FormFile("subtitle")
	defer subtitle.Close()

	tempSub, _ := ioutil.TempFile(path, r.FormValue("content_title")+"_"+r.FormValue("episode_count")+"_*.vtt")
	defer tempSub.Close()
	tempSubName := strings.Replace(tempSub.Name(), "\\", "/", -1)

	fileBytes, _ := ioutil.ReadAll(subtitle)
	tempSub.Write(fileBytes)

	response = uploadSubtitleResponse{
		Status:                "success",
		OriginalSubtitleTitle: handlerSub.Filename,
		OriginalSubtitleSize:  int(handlerSub.Size),
		Language:              r.FormValue("language"),
		SubtitlePath:          tempSubName,
	}

	responseJSON, _ := json.Marshal(response)
	rw.Write(responseJSON)

	database.UpdateSubtitle(r.FormValue("language"), tempSubName, r.FormValue("content_title")+" - "+r.FormValue("episode_count"))
}
