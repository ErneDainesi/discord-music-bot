package yt

import (
	"os/exec"
	"strings"
)

func Download(message string) error {
    url := strings.Split(message, " ")
    runCmd := exec.Command("yt-dlp", url[1])
    return runCmd.Run()
}
