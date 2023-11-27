package configuration

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Configuration struct {
	Input *Input
}

type Input struct {
	Clips        []InputClip `json:"clips"`
	OutputFormat string      `json:"output_format"`
}

type InputClip struct {
	Url   string `json:"url"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

func Load(inputPath string) (*Configuration, error) {
	input, err := loadInput(inputPath)
	if err != nil {
		return nil, err
	}
	c := &Configuration{
		Input: input,
	}

	v := newValidator()
	errors := v.Validate(c)
	if errors != nil {
		for _, err := range errors {
			log.Printf("error: %s", err)
		}
		return nil, fmt.Errorf("errors while reading configuration")
	}

	return c, nil
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
