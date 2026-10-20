package transcript

import (
	"fmt"
	"time"

	"github.com/horiagug/youtube-transcript-api-go/pkg/yt_transcript"
)

type Transcription struct {
	Text     string
	Start    string
	Duration time.Duration
}

func FetchTranscription(videoID string, languages []string) {

	client := yt_transcript.NewClient()

	transcript, err := client.GetFormattedTranscripts(videoID, languages, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(transcript)

}
