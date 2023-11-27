# Build

To build project just run `make build`.

# Run

To run application with input file from docs ([click](./examples/input.json)) just run `make run`. If you want to overwrite parameters after building just run the binary `./bin/video-concatter` and pass your parameters.

## Parameters

* `-input` - path to file with all configuration (ex. `-input ./docs/examples/input.json`)
* `-output` - path where to store output video (ex. `-output ./result.mp4`)
* `-ffmpegBin` - path to FFmpeg binary used for processing (ex. `-ffmpegBin /usr/local/bin/ffmpeg`)

# Some decisions records

1. [ADR-01 Downloading files](./adr/adr_01_dowloading_files.md)
1. [ADR-02 Input data](./adr/adr_02_input_data.md)
1. [ADR-03 Log format](./adr/adr_03_log_format.md)

# Challenges

[Click](./challenges.md)

# Improvements

1. Parsing FFmpeg output to better show progress and reduce noise.
2. Capture exit codes of FFmpeg to better inform about errors.
3. Decide if containers/isolating should be used or not (file from URL is downloaded - my cause security issues).
4. Allow for more detailed description of output format.
5. Videos are not encoded and there might be problems when videos are in different format/codec.
6. Allow for more audio/video streams when concatenating.
7. Introduce better logging (uber-go/zap lib for example).
8. Use Cobra/Viper if application will grow.
9. Use structured logging for easier automation and log parsing.
10. Add more tests fot better test coverage.
11. Add integration e2e test to check everything works.