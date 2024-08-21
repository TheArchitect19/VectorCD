package config

import (
	"os/exec"
)

func ReloadNginx() error {
	cmd := exec.Command("sudo", "nginx", "-s", "reload")
	return cmd.Run()
}
