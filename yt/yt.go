package yt

import (
	"os/exec"
	"strings"
)

func Download(message string) error {
    url := strings.Split(message, " ")
    runCmd := exec.Command("yt-dlp", url[1])
    err := runCmd.Run()
    if err != nil {
        return err
    }
    return nil
}
