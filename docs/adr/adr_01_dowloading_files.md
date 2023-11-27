# Downloading files

## Status

accepted

## Context

FFmpeg can treat URL as an input for transcoding. The file can also be downloaded by other solution and used as input then.

## Decision

File will be downloaded by FFmpeg. Assumption is that problem with download occurrs very occasionally with links from
example and downloading only required part of file to clip it reduces amount of data required
to download drastically. Note: FFmpeg stream file from link and process it instantly. When required part of file is
already processed download process is aborted.

## Consequences

* FFmpeg is not specialized in downloading files. It may have problems on network errors.
* Using other tools may help with retries or faster download.
* Other tools may implement different protocols than FFmpeg can handle.