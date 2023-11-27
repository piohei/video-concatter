package ffmpeg

import (
	"reflect"
	"testing"
)

func TestSecondsToTimeString(t *testing.T) {
	type test struct {
		seconds int
		want    TimeString
	}

	tests := []test{
		{seconds: 10, want: TimeString("00:00:10")},
		{seconds: 2 * 60, want: TimeString("00:02:00")},
		{seconds: 3 * 60 * 60, want: TimeString("03:00:00")},
		{seconds: 99 * 60 * 60, want: TimeString("99:00:00")},
		{seconds: 99*60*60 + 59*60, want: TimeString("99:59:00")},
		{seconds: 99*60*60 + 59*60 + 59, want: TimeString("99:59:59")},
	}

	for _, tc := range tests {
		got := secondsToTimeString(tc.seconds)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
