package config

import (
	"os/exec"
)

func ReloadNginx() error {
	cmd := exec.Command("nginx", "-s", "reload")
	return cmd.Run()
}
