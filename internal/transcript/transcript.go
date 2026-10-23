package transcript

import (
	"time"

	"github.com/horiagug/youtube-transcript-api-go/pkg/yt_transcript"
)

type Transcription struct {
	Text     string
	Start    string
	Duration time.Duration
}

func FetchTranscription(videoID string, languages []string) (string, error) {

	client := yt_transcript.NewClient()

	transcript, err := client.GetFormattedTranscripts(videoID, languages, true)
	if err != nil {
		return "", nil
	}

	return transcript, nil
}
