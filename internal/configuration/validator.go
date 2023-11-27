package configuration

import (
	"fmt"
	"net/url"
	"regexp"
)

var validOutputFormat = regexp.MustCompile(`^\d+:\d+$`)

type Validator struct{}

func newValidator() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(c *Configuration) []error {
	var errors []error
	if clipsErrors := v.validateClips(c.Input.Clips); clipsErrors != nil {
		errors = append(errors, clipsErrors...)
	}
	if clipsErrors := v.validateOutputFormat(c.Input.OutputFormat); clipsErrors != nil {
		errors = append(errors, clipsErrors...)
	}
	return errors
}

func (v *Validator) validateOutputFormat(outputFormat string) []error {
	var errors []error
	if !validOutputFormat.MatchString(outputFormat) {
		errors = append(errors, fmt.Errorf(fmt.Sprintf("output format must match '%s'", validOutputFormat.String())))
	}
	return errors
}

func (v *Validator) validateClips(clips []InputClip) []error {
	var errors []error
	if len(clips) == 0 {
		errors = append(errors, fmt.Errorf("no clips passed"))
	}
	for i, c := range clips {
		if clipErrors := v.validateClip(i, c); clipErrors != nil {
			errors = append(errors, clipErrors...)
		}
	}
	return errors
}

func (v *Validator) validateClip(index int, clip InputClip) []error {
	var errors []error
	if clip.Start < 0 {
		errors = append(errors, clipValidationError(index, "start must be at least 0"))
	}
	if clip.End < 0 {
		errors = append(errors, clipValidationError(index, "end must be at least 0"))
	}
	if clip.End <= clip.Start {
		errors = append(errors, clipValidationError(index, "end must be greater than start"))
	}
	if _, err := url.ParseRequestURI(clip.Url); err != nil {
		errors = append(errors, clipValidationError(index, fmt.Sprintf("invalid url: %s", err)))
	}

	return errors
}

func clipValidationError(index int, msg string) error {
	return fmt.Errorf(fmt.Sprintf("clip[%d]: %s", index, msg))
}
