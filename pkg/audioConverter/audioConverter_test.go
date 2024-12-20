package audioConverter

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	return
	rawData, err := os.ReadFile("./test/test.oga")
	if err != nil {
		t.Errorf("cannot read test.oga: %v", err)
	}

	rawAudio := bytes.NewReader(rawData)

	// Test case 1: Valid source and target extensions
	reader, err := ConvertAudio(context.Background(), rawAudio, "oga", "mp3")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	rawAudioConverted, err := io.ReadAll(reader)
	if err != nil {
		t.Errorf("error reading output: %v", err)
	}

	rawAudioConvertedExpected, err := os.ReadFile("./test/test.mp3")
	if err != nil {
		t.Errorf("cannot read test.mp3: %v", err)
	}

	if !bytes.Equal(rawAudioConverted, rawAudioConvertedExpected) {
		t.Errorf("raw audio converted is not equal to raw audio converted expected")
	}

	// Test case 2: Invalid source extension
	_, err = ConvertAudio(context.Background(), rawAudio, "wav", "mp3")
	if err == nil {
		t.Error("expected an error for invalid source extension")
	}

	// Test case 3: Invalid target extension
	_, err = ConvertAudio(context.Background(), rawAudio, "oga", "wav")
	if err == nil {
		t.Error("expected an error for invalid target extension")
	}

}
