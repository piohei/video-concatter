package configuration

import (
	"reflect"
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	type test struct {
		config *Configuration
		errors []string
	}

	tests := []test{
		{
			config: &Configuration{
				Input: &Input{
					Clips:        nil,
					OutputFormat: "",
				},
			},
			errors: []string{
				"no clips passed",
				"output format must match '^\\d+:\\d+$'",
			},
		},
		{
			config: &Configuration{
				Input: &Input{
					Clips:        nil,
					OutputFormat: "16:9",
				},
			},
			errors: []string{
				"no clips passed",
			},
		},
		{
			config: &Configuration{
				Input: &Input{
					Clips: []InputClip{
						{
							Url:   "",
							Start: 0,
							End:   1,
						},
					},
					OutputFormat: "16:9",
				},
			},
			errors: []string{
				"clip[0]: invalid url: parse \"\": empty url",
			},
		},
		{
			config: &Configuration{
				Input: &Input{
					Clips: []InputClip{
						{
							Url:   "example.com",
							Start: 0,
							End:   1,
						},
					},
					OutputFormat: "16:9",
				},
			},
			errors: []string{
				"clip[0]: invalid url: parse \"example.com\": invalid URI for request",
			},
		},
		{
			config: &Configuration{
				Input: &Input{
					Clips: []InputClip{
						{
							Url:   "http://example.com",
							Start: 0,
							End:   0,
						},
					},
					OutputFormat: "16:9",
				},
			},
			errors: []string{
				"clip[0]: end must be greater than start",
			},
		},
		{
			config: &Configuration{
				Input: &Input{
					Clips: []InputClip{
						{
							Url:   "http://example.com",
							Start: -2,
							End:   -1,
						},
					},
					OutputFormat: "16:9",
				},
			},
			errors: []string{
				"clip[0]: start must be at least 0",
				"clip[0]: end must be at least 0",
			},
		},
		{
			config: &Configuration{
				Input: &Input{
					Clips: []InputClip{
						{
							Url:   "http://example.com",
							Start: 0,
							End:   1,
						},
					},
					OutputFormat: "16:9",
				},
			},
			errors: []string{},
		},
	}

	v := newValidator()
	for _, tc := range tests {
		gotErrors := v.Validate(tc.config)
		errorsAsString := toString(gotErrors)
		if !reflect.DeepEqual(tc.errors, errorsAsString) {
			t.Fatalf("expected: [%v], got: [%v]", strings.Join(tc.errors, ","), strings.Join(errorsAsString, ","))
		}
	}
}

func toString(errors []error) []string {
	res := []string{}
	for _, err := range errors {
		res = append(res, err.Error())
	}
	return res
}
