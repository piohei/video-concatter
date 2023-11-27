package ffmpeg

type FFmpeg struct {
	commandExecutor      *commandExecutor
	commandArgsFormatter *commandArgsFormatter
}

type ClippedInput struct {
	Input      string
	Start, End int
}

func NewFFmpeg() *FFmpeg {
	return &FFmpeg{
		//commandExecutor:      &commandExecutor{binPath: "ffmpeg"},
		commandExecutor:      &commandExecutor{binPath: "/home/piotr/Downloads/ffmpeg-5.1.1-amd64-static/ffmpeg"},
		commandArgsFormatter: &commandArgsFormatter{},
	}
}

func (f *FFmpeg) ClipAndJoinVideo(inputs []ClippedInput, output string, aspect string) error {
	args := f.commandArgsFormatter.ClipAndJoin(inputs, output, aspect)
	return f.commandExecutor.execute(args)
}
