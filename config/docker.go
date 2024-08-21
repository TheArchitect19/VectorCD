package config

import (
	"fmt"
	"os/exec"
)

func RunDockerContainer(imageName string, port int) error {
	cmd := exec.Command("docker", "run", "-d", "-p", fmt.Sprintf("%d:8000", port), "--name", fmt.Sprintf("container_%d", port), imageName)
	return cmd.Run()
}
