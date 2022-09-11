# byungflix-backend
## introduction
HLS Video Streaming Backend Server

This server adds video(.mkv) and WebVTT subtitle files to the server via HTTP request.
Uploaded videos are convert to m3u8 extension.
Also, can stream this m3u8 videos in web.
This HLS streaming is for front-end server.

## Tech Stack
- Go
  - gorilla/mux
- ffmpeg
- mongoDB
