package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/piohei/video-concatter/internal/downloader"
	"github.com/piohei/video-concatter/internal/ffmpeg"
	"io"
	"log"
	"os"
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

	input, err := loadInput(inputPath)
	if err != nil {
		log.Fatalf("error loading input data: %s", err)
	}

	d := downloader.NewDownloader(downloader.Retries(3))
	ff := ffmpeg.NewFFmpeg()
	for i, c := range input.Clips {
		in := fmt.Sprintf("/tmp/i_%d.mp4", i)
		out := fmt.Sprintf("/tmp/o_%d.mp4", i)
		d.Download(c.Url, in)
		ff.ClipVideo(in, out, c.Start, c.End)
	}
}

func loadInput(path string) (*Input, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("error closing input file: %s", err)
		}
	}(file)

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %s", err)
	}

	input := &Input{}
	err = json.Unmarshal(bytes, input)
	if err != nil {
		return nil, fmt.Errorf("error decoding file: %s", err)
	}

	return input, nil
}
