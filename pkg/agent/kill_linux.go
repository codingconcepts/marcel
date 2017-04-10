// +build linux

package agent

import (
	"fmt"
	"os/exec"

	"github.com/Sirupsen/logrus"
)

func (a *Agent) killMachine() (err error) {
	a.Logger.Info("killing machine")

	cmd := exec.Command("reboot", "-f")
	return cmd.Run()
}

func (a *Agent) killProcess(name string) (err error) {
	a.Logger.WithFields(logrus.Fields{
		"process": name,
	}).Info("killing process")

	cmd := exec.Command("kill", "-KILL", fmt.Sprintf("'pgrep %s'", name))
	return cmd.Run()
}
