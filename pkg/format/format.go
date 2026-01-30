package format

const (
	// FfmpegOutputFormat is the supported audio extension for ffmpeg conversion
	FfmpegOutputFormat = "mp3"
)

// SupportedInputFormats lists audio formats that ffmpeg can decode via pipe:0
var SupportedInputFormats = map[string]bool{
	"oga":  true, // WhatsApp voice notes (Opus/OGG)
	"ogg":  true, // OGG Vorbis/Opus
	"mp4":  true, // Instagram/FB audio (AAC)
	"m4a":  true, // MPEG-4 Audio
	"aac":  true, // Advanced Audio Coding
	"webm": true, // WebM (Opus/Vorbis)
	"wav":  true, // Uncompressed
	"mp3":  true, // MPEG Layer 3
}
