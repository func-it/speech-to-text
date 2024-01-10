FROM golang:1.21

# Update package list and install ffmpeg
RUN apt-get update && apt-get install -y ffmpeg

ENV APP_HOME /app

RUN mkdir -p "$APP_HOME"

COPY . "/app"

# Set the Current Working Directory inside the container
WORKDIR "$APP_HOME"

# Build the binary
RUN go mod tidy && go build -o speech_to_text 

#Exposing port 12007
EXPOSE 12007

CMD ["./speech_to_text", "service", "speechToText"]