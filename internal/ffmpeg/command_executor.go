package ffmpeg

import (
	"bufio"
	"fmt"
	"os/exec"
)

type command struct {
}

type commandExecutor struct {
	binPath string
}

func (c *commandExecutor) execute(commandArgs []string) error {
	cmd := exec.Command(c.binPath, commandArgs...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}

	s := bufio.NewScanner(stdout)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		fmt.Println(s.Text())
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
