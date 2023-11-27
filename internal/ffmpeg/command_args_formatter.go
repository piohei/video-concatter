package ffmpeg

import (
	"fmt"
	"strings"
)

type TimeString string

type commandArgsFormatter struct {
}

func (cf *commandArgsFormatter) ClipAndJoin(inputs []ClippedInput, output string, aspect string) []string {
	var args []string
	for _, i := range inputs {
		args = append(args, cf.clippedInput(i.Input, i.Start, i.End)...)
	}
	args = append(args, cf.concatFilter(len(inputs))...)
	args = append(args, cf.aspect(aspect)...)
	args = append(args, cf.output(output)...)
	return args
}

func (cf *commandArgsFormatter) clippedInput(input string, start, end int) []string {
	return []string{
		"-ss", string(secondsToTimeString(start)),
		"-to", string(secondsToTimeString(end)),
		"-i", input,
	}
}

func (cf *commandArgsFormatter) concatFilter(numOfClips int) []string {
	var inputClips []string
	for i := 0; i < numOfClips; i++ {
		inputClips = append(inputClips, fmt.Sprintf("[%d:v] [%d:a]", i, i))
	}
	return []string{
		"-filter_complex", fmt.Sprintf("%s concat=n=%d:v=1:a=1 [v] [a]", strings.Join(inputClips, " "), numOfClips),
		"-map", "[v]",
		"-map", "[a]",
	}
}

func (cf *commandArgsFormatter) aspect(aspect string) []string {
	return []string{
		"-aspect", aspect,
	}
}

func (cf *commandArgsFormatter) output(output string) []string {
	return []string{
		"-y", output,
	}
}

func secondsToTimeString(tsec int) TimeString {
	seconds := tsec % 60
	minutes := tsec / 60 % 60
	hours := tsec / 3600
	return TimeString(fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds))
}
