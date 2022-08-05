package database

import (
	"fmt"
)

type Series struct {
	Title       string `json:"title"`
	TitleKor    string `json:"title_kor"`
	Description string `json:"description"`
}

type Video struct {
	Title         string `json:"title"`
	ContentTitle  string `json:"content_title"`
	EpisodeCount  int    `json:"episode_count"`
	Description   string `json:"description"`
	BroadcastDate string `json:"broadcast_date"`
	UploadDate    string `json:"upload_time"`
	VideoPath     string `json:"video_path"`
	SubtitlePath  string `json:"subtitle_path"`
}

func CreateSeries(series Series) {
	fmt.Println(series)
}