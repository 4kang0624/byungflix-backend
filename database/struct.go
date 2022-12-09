package database

type Series struct {
	Title       string `json:"title"`
	TitleKor    string `json:"title_kor"`
	Description string `json:"description"`
}

type Video struct {
	Title        string            `json:"title"`
	SeriesTitle  string            `json:"content_title"`
	EpisodeCount int               `json:"episode_count"`
	ReleaseDate  string            `json:"release_date"`
	UploadDate   string            `json:"upload_time"`
	VideoPath    string            `json:"video_path"`
	VideoPathHls string            `json:"video_path_hls"`
	SubtitlePath map[string]string `json:"subtitle_path"`
}
