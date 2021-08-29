package aviator

import (
	"fmt"
	"os/exec"
)

const (
	preScript string = `sudo apt upgrade \
sudo apt update`
)

type Aviator struct {
}

func (a *Aviator) Run() error {
	preCmd := exec.Command(preScript)
	if err := preCmd.Run(); err != nil {
		return fmt.Errorf("unable to run initial setup: %v", err)
	}

	return nil
}
