// +build linux

package browse

import (
	"os/exec"
)

type LBrowser struct {
}

func (l LBrowser) Browse(url string) error {
	c := exec.Command("xdg-open", url)
	
	if err := c.Run(); err != nil {
		return err
	}
	
	return nil
}
