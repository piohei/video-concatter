# Input data

## Status

accepted

## Context

Example input data is shown [here](https://scoreplay.notion.site/Video-backend-project-4a793802e1cb403f9ce8464fd21959c4).
Based on it some assumptions were made (that simplifies the PoC):
* only single clip can be extracted from single file passed as URL (instead of list of start/end tuples)
* video format is treated only as video aspect ratio

## Decision

Application will accept single path to file in json format with all configuration being set there.
Video format will be treated as aspect ratio that can accept only string in format `(d+):(d+)`. It will be
treated as Display Aspect Ratio and set for output video. Video format will not include various options to
re-encode videos or change their formats, etc.

## Consequences

* Changing just DAR for output video will be faster but will produce not so good quality
* All assumptions will limit the scope of PoC and simplify it