package main

import (
	"flag"
	"log"

	"github.com/piohei/video-concatter/internal/configuration"
	"github.com/piohei/video-concatter/internal/ffmpeg"
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
	var outputPath string
	var ffmpegBin string

	flag.StringVar(&inputPath, "input", "input.json", "path to file with all configuration")
	flag.StringVar(&inputPath, "output", "result.mp4", "path where to store output video")
	flag.StringVar(&inputPath, "ffmpegBin", "ffmpeg", "path to FFmpeg binary used for processing")
	flag.Parse()

	if inputPath == "" {
		log.Fatalf("flag 'input' not passed or empty.")
	}

	config, err := configuration.Load(inputPath)
	if err != nil {
		log.Fatalf("error loading configuration: %s", err)
	}

	ff := ffmpeg.NewFFmpeg(ffmpegBin)
	inputs := toFFmpegClippedInput(config.Input.Clips)
	err = ff.ClipAndJoinVideo(inputs, outputPath, config.Input.OutputFormat)
	if err != nil {
		log.Fatalf("error while processing videos: %s", err)
	}
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
