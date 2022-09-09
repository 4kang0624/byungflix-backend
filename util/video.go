package util

import (
	"fmt"
	"os/exec"
	"strings"
)

func EncodeMkvToHls(video string) string {
	resultVideoTitle := strings.Split(video, ".")[0]
	cmd := exec.Command("cmd", "/c", "ffmpeg",
		"-i", video,
		"-b:v", "1M",
		"-g", "60",
		"-hls_time", "10",
		"-hls_list_size", "0",
		"-hls_segment_size", "500000",
		resultVideoTitle+".m3u8")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	return resultVideoTitle + ".m3u8"
}
