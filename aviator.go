package aviator

import (
	"fmt"
	"os"
	"os/exec"
)

// Define bash scripts that will be used to setup various aspects of the deployment.
const (
	preScript string = `sudo apt upgrade && \
sudo apt update`

	nginxScript string = `sudo apt install nginx && \
sudo ufw allow 'Nginx Full' && \
sudo ufw allow 'OpenSSH' && \
sudo ufw enable && \
sudo systemctl start nginx
`
)

type Aviator struct {
}

func (a *Aviator) Run() error {

	if err := runBashScript(preScript); err != nil {
		return fmt.Errorf("unable to run initial setup: %v", err)
	}

	if err := runBashScript(nginxScript); err != nil {
		return fmt.Errorf("unable to run initial setup: %v", err)
	}

	return nil
}

// runBashScript executes the given script (which should be a series of commands in the form
// `cmdA && cmdB && ...`). This function blocks until script execution is complete, returning
// an error if the script fails to successfully complete.
func runBashScript(script string) error {
	cmd := exec.Command("bash", "-c", script)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
