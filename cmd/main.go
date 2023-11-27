package main

import (
	"flag"
	"fmt"
	"github.com/piohei/video-concatter/internal/configuration"
	"github.com/piohei/video-concatter/internal/ffmpeg"
	"log"
)

type Input struct {
	Clips        []InputClip `json:"clips"`
	OutputFormat string      `json:"output_format"`
}

type InputClip struct {
	Url   string `json:"url"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

func main() {
	var inputPath string

	flag.StringVar(&inputPath, "input", "input.json", "path to file with all configuration.")
	flag.Parse()

	if inputPath == "" {
		log.Fatalf("flag 'input' not passed or empty.")
	}

	config, err := configuration.Load(inputPath)
	if err != nil {
		log.Fatalf("error loading configuration: %s", err)
	}

	ff := ffmpeg.NewFFmpeg()
	inputs := toFFmpegClippedInput(config.Input.Clips)
	ff.ClipAndJoinVideo(inputs, outputFilePath(0), config.Input.OutputFormat)
}

func outputFilePath(index int) string {
	return fmt.Sprintf("/tmp/o_%d.mp4", index)
}

func toFFmpegClippedInput(inputClips []configuration.InputClip) []ffmpeg.ClippedInput {
	var res []ffmpeg.ClippedInput
	for _, c := range inputClips {
		res = append(res, ffmpeg.ClippedInput{
			Input: c.Url,
			Start: c.Start,
			End:   c.End,
		})
	}
	return res
}
