# FFMPEG

> Source: <http://slhck.info/ffmpeg-encoding-course/>

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