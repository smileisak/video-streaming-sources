# FFMPEG

> Source: <http://slhck.info/ffmpeg-encoding-course/>
> Source: <https://github.com/leandromoreira/ffmpeg-libav-tutorial#intro>

```bash
$ ffmpeg \
-y \
-c:a aac -c:v libopenh264 \
-i bunny_1080p_60fps.mp4 \
-c:v libvpx-vp9 -c:a libvorbis \
bunny_1080p_60fps_vp9.webm
```

## Convert the video container

```bash
ffmpeg -i bunny_1080p_60fps.mp4  -c copy bunny_1080p_60fps.mkv
```

## Splits the video

```bash
ffmpeg -ss <start> -i <input> -t <duration> -c copy <output>
ffmpeg -ss <start> -i <input> -to <end> -c copy <output>
```

```bash
ffmpeg -ss 00:01:50 -i bunny_1080p_60fps.mkv  -t 10.5 -c copy splited_bunny_1080p_60fps.mkv
```

## CRF Guide (Constant Rate Factor in x264, x265 and libvpx)

The Constant Rate Factor (CRF) is the default quality (and rate control) setting for the x264 and x265 encoders, and it’s also available for libvpx. With x264 and x265, you can set the values between **0 and 51**, where **lower values would result in better quality**, at the expense of higher file sizes. Higher values mean more compression, but at some point you will notice the quality degradation.

**For x264, sane values are between 18 and 28. The default is 23, so you can use this as a starting point.**

```bash
$ ffmpeg -i bunny_1080p_60fps.mp4 -c:v libx264 -crf 51 lower_crf_51_bunny_1080p_60fps.mp4

ffmpeg -i bunny_1080p_60fps.mp4 -c:v libx264 -crf 0 higher_crf_0_bunny_1080p_60fps.mp4
```

If you’re unsure about what CRF to use, begin with the default and change it according to your subjective impression of the output. Is the quality good enough? No? Then set a lower CRF. Is the file size too high? Choose a higher CRF. **A change of ±6 should result in about half/double the file size, although your results might vary.**

## Presets

```bash
ffmpeg -i <input> -c:v libx264 -crf 23 -preset ultrafast -an output.mkv
ffmpeg -i <input> -c:v libx264 -crf 23 -preset medium -an output.mkv
ffmpeg -i <input> -c:v libx264 -crf 23 -preset veryslow -an output.mkv
```

## Getting media information with ffprobe

```bash
ffprobe veryslow_bunny_1080p_60fps.mkv -show_streams
ffprobe veryslow_bunny_1080p_60fps.mkv -select_streams v -show_format
```

JSON Format:

```bash
ffprobe bunny_1080p_60fps.mkv -select_streams v -show_packets -of json | jq
```

## Common video operations

```bash
ffmpeg \
-i bunny_1080p_60fps.mp4 \
-c:v libx265 \
bunny_1080p_60fps_h265.mp4
```
