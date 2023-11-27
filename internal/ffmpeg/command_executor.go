package ffmpeg

import (
	"bufio"
	"log"
	"os/exec"
)

type command struct {
}

type commandExecutor struct {
	binPath string
}

func (c *commandExecutor) execute(commandArgs []string) error {
	cmd := exec.Command(c.binPath, commandArgs...)
	log.Printf("Executing command: '%s'.", cmd.String())
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}

	so := bufio.NewScanner(stdout)
	for so.Scan() {
		log.Println("stdout: " + so.Text())
	}

	se := bufio.NewScanner(stderr)
	for se.Scan() {
		log.Println("stderr: " + se.Text())
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
