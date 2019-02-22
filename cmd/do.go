package cmd

import "os/exec"

func do(c command) error {
	if c.extraFlag != "" {
		_, err := exec.Command(c.Command, c.Flag, c.Path, c.extraFlag, c.outPath).Output()
		if err != nil {
			return err
		}
	} else {
		_, err := exec.Command(c.Command, c.Flag, c.Path).Output()
		if err != nil {
			return err
		}
	}
	return nil
}

type command struct {
	Command   string
	Flag      string
	Path      string
	extraFlag string
	outPath   string
}
