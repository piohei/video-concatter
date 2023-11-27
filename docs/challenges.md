# Challenges

During PoC preparation I encountered seg fault from time to time depending on used set of FFmpeg paramters.

## File downloading

First idea was to download files before processing and allow to pass multiple clips per single file, retrying download, etc. It turned out in example video files are very big and only part of the start is being extracted. Because of that download was moved to FFmpeg so it can download only the required part of the video instead of whole.

## Trimming and concatenating

### First version

First version used a `-ss` and `-to` parameter to clip input video (placed before `-i` option). When used together with concat filter it resulted in FFmpeg sef faults.

### Second version

Second version uses trim filter to trim inputs also and works well with newest FFmpeg version.

## Time limit

I could only spend about 5-6 hours for that app which was mostly spent on figuring out proper set of parameters for FFmpeg to work (due to seg faults).