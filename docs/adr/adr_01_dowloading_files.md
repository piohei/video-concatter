# Downloading files

## Status

accepted

## Context

FFmpeg can treat URL as an input for transcoding. The file can also be downloaded by other solution and used as input then.

## Decision

File will be downloaded by other tool instead of FFmpeg.

## Consequences

* FFmpeg is not specialized in downloading files. It may have problems on network errors.
* Using other tools may help with retries or faster download hat standard.
* Other tools may implement different protocols than FFmpeg can handle.