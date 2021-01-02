// +build linux

package browse

import (
	"os/exec"
)

type DefaultBrowse struct {
}

func (d DefaultBrowse) Browse(url string) error {
	c := exec.Command("xdg-open", url)
	
	if err := c.Run(); err != nil {
		return err
	}
	
	return nil
}
