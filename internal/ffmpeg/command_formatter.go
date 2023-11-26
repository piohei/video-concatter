package ffmpeg

import "fmt"

type TimeString string

type commandArgsFormatter struct {
}

func (cf *commandArgsFormatter) clipVideo(input, output string, start, end int) []string {
	return []string{
		"-ss", string(secondsToTimeString(start)),
		"-to", string(secondsToTimeString(end)),
		"-i", input,
		"-c", "copy",
		"-y", output,
	}
}

func secondsToTimeString(tsec int) TimeString {
	seconds := tsec % 60
	minutes := tsec / 60 % 60
	hours := tsec / 3600
	return TimeString(fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds))
}
