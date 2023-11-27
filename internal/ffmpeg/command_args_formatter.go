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
		args = append(args, cf.input(i.Input)...)
	}
	args = append(args, cf.trimAndConcat(inputs)...)
	args = append(args, cf.aspect(aspect)...)
	args = append(args, cf.output(output)...)
	return args
}

func (cf *commandArgsFormatter) input(input string) []string {
	return []string{
		"-i", input,
	}
}

func (cf *commandArgsFormatter) trimAndConcat(inputs []ClippedInput) []string {
	var filter []string
	for i, c := range inputs {
		filter = append(filter, cf.trimFilters(i, c.Start, c.End)...)
	}
	filter = append(filter, cf.concatFilter(len(inputs)))
	return []string{
		"-filter_complex", strings.Join(filter, ";"),
		"-map", "[v]",
		"-map", "[a]",
	}
}

func (cf *commandArgsFormatter) trimFilters(index, start, end int) []string {
	return []string{
		fmt.Sprintf("[%d:v]trim=start=%d:end=%d,setpts=PTS-STARTPTS[v%d]", index, start, end, index),
		fmt.Sprintf("[%d:a]atrim=start=%d:end=%d,asetpts=PTS-STARTPTS[a%d]", index, start, end, index),
	}
}

func (cf *commandArgsFormatter) concatFilter(numOfClips int) string {
	var concatFilter []string
	for i := 0; i < numOfClips; i++ {
		concatFilter = append(concatFilter, fmt.Sprintf("[v%d][a%d]", i, i))
	}
	concatFilter = append(concatFilter, fmt.Sprintf("concat=n=%d:v=1:a=1[v][a]", numOfClips))
	return strings.Join(concatFilter, "")
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
