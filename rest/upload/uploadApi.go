package upload

import (
	"byungflix-backend/database"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

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
	video, handlerVid, errFormFileVid := r.FormFile("video")
	if errFormFileVid != nil {
		fmt.Println(errFormFileVid)
		return
	}
	defer video.Close()

	subtitle, handlerSub, errFormFileSub := r.FormFile("subtitle")
	if errFormFileSub != nil {
		fmt.Println(errFormFileSub)
		return
	}
	defer subtitle.Close()

	path := "contents/" + r.FormValue("content_title") + "/" + r.FormValue("episode_count")
	MkdirErr := os.Mkdir(path, 0755)
	if MkdirErr != nil {
		log.Fatal(MkdirErr)
	}

	fmt.Println("File Info")
	fmt.Println("File Name: ", handlerVid.Filename)
	fmt.Println("File Size: ", handlerVid.Size)
	fmt.Println("File Type: ", handlerVid.Header)
	fmt.Println("File Name: ", handlerSub.Filename)
	fmt.Println("File Size: ", handlerSub.Size)
	fmt.Println("File Type: ", handlerSub.Header)

	tempVideo, errTempVideo := ioutil.TempFile(path, r.FormValue("content_title")+"_"+r.FormValue("episode_count")+"_*.mkv")
	if errTempVideo != nil {
		log.Fatal(errTempVideo)
	}
	defer tempVideo.Close()

	tempSub, errTempSub := ioutil.TempFile(path, r.FormValue("content_title")+"_"+r.FormValue("episode_count")+"_*.ass")
	if errTempSub != nil {
		log.Fatal(errTempSub)
	}
	defer tempSub.Close()

	fileBytes, errFileBytes := ioutil.ReadAll(video)
	if errFileBytes != nil {
		log.Fatal(errFileBytes)
	}
	tempVideo.Write(fileBytes)
	fmt.Println(tempVideo.Name())

	fileBytes, errFileBytes = ioutil.ReadAll(subtitle)
	if errFileBytes != nil {
		log.Fatal(errFileBytes)
	}
	tempSub.Write(fileBytes)
	fmt.Println("Done")
}
