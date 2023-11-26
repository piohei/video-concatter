package ffmpeg

import "fmt"

type TimeString string

type commandFormatter struct {
	binPath string
}

func (cf *commandFormatter) ClipVideo(input, output string, start, end int) string {
	return fmt.Sprintf(
		"%s -ss %s -to %s -i %s -c copy %s",
		cf.binPath,
		secondsToTimeString(start),
		secondsToTimeString(end),
		input,
		output,
	)
}

func secondsToTimeString(tsec int) TimeString {
	seconds := tsec % 60
	minutes := tsec / 60 % 60
	hours := tsec / 3600
	return TimeString(fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds))
}
