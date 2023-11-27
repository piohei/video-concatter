package ffmpeg

type FFmpeg struct {
	commandExecutor      *commandExecutor
	commandArgsFormatter *commandArgsFormatter
}

type ClippedInput struct {
	Input      string
	Start, End int
}

func NewFFmpeg(ffmpegBin string) *FFmpeg {
	return &FFmpeg{
		commandExecutor:      &commandExecutor{binPath: ffmpegBin},
		commandArgsFormatter: &commandArgsFormatter{},
	}
}

func (f *FFmpeg) ClipAndJoinVideo(inputs []ClippedInput, output string, aspect string) error {
	args := f.commandArgsFormatter.ClipAndJoin(inputs, output, aspect)
	return f.commandExecutor.execute(args)
}
