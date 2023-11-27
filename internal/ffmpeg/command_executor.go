package ffmpeg

import (
	"bufio"
	"bytes"
	"log"
	"os/exec"
)

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

	go func() {
		so := bufio.NewScanner(stdout)
		// We split also by carriage return because that's how FFmpeg shows progress same line
		so.Split(scanLinesOrCarriageReturn)
		for so.Scan() {
			log.Println("[FFmpeg][stdout] " + so.Text())
		}
	}()

	go func() {
		se := bufio.NewScanner(stderr)
		// We split also by carriage return because that's how FFmpeg shows progress same line
		se.Split(scanLinesOrCarriageReturn)
		for se.Scan() {
			log.Println("[FFmpeg][stderr] " + se.Text())
		}
	}()

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

// scanLinesOrCarriageReturn is a modified function from bufio package to allow also split on carriage return.
func scanLinesOrCarriageReturn(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil
	}
	if i := bytes.IndexByte(data, '\r'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
