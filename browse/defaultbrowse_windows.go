// +build windows

package browse

import "os/exec"

type DefaultBrowse struct {}

func (d DefaultBrowse) Browse(url string) error {
	c := exec.Command("start", url)

	if err := c.Run(); err != nil {
		return err
	}

	return nil
}
