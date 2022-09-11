package util

import (
	"fmt"
	"os/exec"
	"strings"
)

func EncodeMkvToHls(video string) string {
	resultVideoTitle := strings.Split(video, ".")[0]
	cmd := exec.Command("ffmpeg",
		"-i", video,
		"-profile:v", "baseline",
		"-level", "3.0",
		"-start_number", "0",
		"-hls_time", "10",
		"-hls_list_size", "0",
		"-f", "hls",
		resultVideoTitle+".m3u8")
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}

	return resultVideoTitle + ".m3u8"
}
