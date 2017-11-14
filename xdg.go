package xdg

import (
	"os/exec"
)

func Open(url string) error {
	cmd := exec.Command("xdg-open", url)
	err := cmd.Run()
	return err
}
