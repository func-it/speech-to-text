package speechToText

import (
	"bytes"
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/func-it/go/logi"
	"github.com/func-it/go/types"
	"github.com/func-it/speech-to-text/pkg/audioConverter"
	"github.com/func-it/speech-to-text/pkg/format"
	speechtotext "github.com/func-it/speech-to-text/pkg/speechToText"
)

type SpeechToTextService struct {
	speechToText *speechtotext.SpeechToText
}

func NewSpeechToTextService(openaiToken string) (*SpeechToTextService, error) {
	err := audioConverter.IsFFMpegAvailable()
	if err != nil {
		return nil, err
	}

	speechToText := speechtotext.NewSpeechToText(openaiToken)

	return &SpeechToTextService{
		speechToText: speechToText,
	}, nil
}

func (a *SpeechToTextService) Ping(_ context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return empty, nil
}

func (a *SpeechToTextService) SpeechToText(ctx context.Context, req *types.SpeechToTextRequest) (*types.SpeechToTextResponse, error) {
	// XXX TODO : AddCandidates validation in proto file
	rawAudio := req.Data
	sourceExtension := strings.ToLower(req.GetSourceExtension())

	if len(rawAudio) == 0 {
		return nil, logi.ErrorNReturn(status.Errorf(codes.FailedPrecondition, "raw audio is empty..."))
	}

	if sourceExtension != format.FfmpegInputFormat {
		return nil, logi.ErrorNReturn(status.Errorf(codes.FailedPrecondition, "source extension is not %s", format.FfmpegInputFormat))
	}

	rawAudioReader := bytes.NewReader(rawAudio)

	rawAudioConvertedReader, err := audioConverter.ConvertAudio(ctx, rawAudioReader, sourceExtension, format.FfmpegOutputFormat)
	if err != nil {
		return nil, logi.ErrorfNWrapNReturn(err, "cannot convert audio : %v", err)
	}

	text, err := a.speechToText.GetTextFromMp3AudioMessage(ctx, rawAudioConvertedReader)
	if err != nil {
		return nil, logi.ErrorfNWrapNReturn(err, "cannot get text from audio : %v", err)
	}

	return &types.SpeechToTextResponse{
		Text: text,
	}, nil
}
