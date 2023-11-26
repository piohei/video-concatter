package ffmpeg

type FFmpeg struct {
	commandExecutor      *commandExecutor
	commandArgsFormatter *commandArgsFormatter
}

func NewFFmpeg() *FFmpeg {
	return &FFmpeg{
		commandExecutor:      &commandExecutor{binPath: "ffmpeg"},
		commandArgsFormatter: &commandArgsFormatter{},
	}
}

func (f *FFmpeg) ClipVideo(input, output string, start, end int) error {
	args := f.commandArgsFormatter.clipVideo(input, output, start, end)
	return f.commandExecutor.execute(args)
}
