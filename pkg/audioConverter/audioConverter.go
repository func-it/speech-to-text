package audioConverter

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"

	"github.com/func-it/speech-to-text/pkg/format"
)

type ffmpegCommandArgs struct {
	OutputCodec  string
	Bitrate      string // in kbps (ex: 192k)
	OutputFormat string
}

func newDefaultffmpegCommandArgs() ffmpegCommandArgs {
	return ffmpegCommandArgs{
		OutputCodec:  format.FfmpegOutputFormat,
		Bitrate:      "192k",
		OutputFormat: format.FfmpegOutputFormat,
	}
}

func IsFFMpegAvailable() error {
	err := exec.Command("ffmpeg", "-version").Run()
	if err != nil {
		return fmt.Errorf("error running ffmpeg command (probably not available on the system): %w", err)
	}

	return nil
}

// XXX TODO : Use context to stop the ffmpeg command if it takes too long to execute
func ConvertAudio(ctx context.Context, rawAudio io.Reader, sourceExtension string, targetExtension string) (io.Reader, error) {
	if sourceExtension != format.FfmpegInputFormat {
		return nil, fmt.Errorf("only %s extension is supported, got %s", format.FfmpegInputFormat, sourceExtension)
	}

	if targetExtension != format.FfmpegOutputFormat {
		return nil, fmt.Errorf("only %s extension is supported, got %s", format.FfmpegOutputFormat, targetExtension)
	}

	cmgArgs := newDefaultffmpegCommandArgs()

	// Create the ffmpeg command
	cmd := exec.Command(
		"ffmpeg",
		"-i", "pipe:0", // Input from the pipe
		"-acodec", cmgArgs.OutputCodec,
		"-ab", cmgArgs.Bitrate,
		"-f", cmgArgs.OutputFormat,
		"pipe:1", // Output to the pipe
	)

	var outputBuffer bytes.Buffer
	cmd.Stdin = rawAudio
	cmd.Stdout = &outputBuffer
	// cmd.Stdout = outputPipeWriter

	/*
		// Create an in-memory buffer to store the output
		var outputBuffer bytes.Buffer
		go func() {
			defer outputPipeReader.Close()
			// copy the data written to the PipeReader via the cmd to stdout
			if _, err := io.Copy(&outputBuffer, outputPipeReader); err != nil {
				log.Fatal(err)
			}
		}()
	*/

	// Start the ffmpeg command
	err := cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("error starting ffmpeg command: %w", err)
	}

	// Wait for the ffmpeg command to complete
	err = cmd.Wait()
	if err != nil {
		return nil, fmt.Errorf("error waiting for ffmpeg command: %w", err)
	}

	return bytes.NewReader(outputBuffer.Bytes()), nil
}
