package yt

import (
	"os/exec"
	"strings"
)

func Download(message string) error {
    url := strings.Split(message, " ")
    args := []string{
        "-x",
        "--audio-format", "m4a",
        url[1],
    }
    runCmd := exec.Command("yt-dlp", args...)
    return runCmd.Run()
}
