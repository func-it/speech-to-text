package speechToText

import (
	"bytes"
	"context"
	"strings"

	types "github.com/func-it/speechToText/gen/service"
	"github.com/func-it/speechToText/pkg/audioConverter"
	"github.com/func-it/speechToText/pkg/logi"
	speechtotext "github.com/func-it/speechToText/pkg/speechToText"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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
	// XXX TODO : Add validation in proto file
	rawAudio := req.Data
	sourceExtension := strings.ToLower(req.GetSourceExtension())

	if len(rawAudio) == 0 {
		return nil, logi.ErrorNReturn(status.Errorf(codes.FailedPrecondition, "raw audio is empty"))
	}

	if sourceExtension != "oga" {
		return nil, logi.ErrorNReturn(status.Errorf(codes.FailedPrecondition, "source extension is not oga"))
	}

	rawAudioReader := bytes.NewReader(rawAudio)

	rawAudioConvertedReader, err := audioConverter.ConvertAudio(ctx, rawAudioReader, sourceExtension, "mp3")
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
