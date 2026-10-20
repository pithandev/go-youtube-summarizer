package main

import (
	"fmt"
	"os"

	"github.com/pithandev/go-youtube-summarizer/internal/transcript"
)

func main() {
	fmt.Println(os.Args)
	videoID := os.Args[1]
	languages := make([]string, 1)
	languages[0] = "en"
	for i, arg := range os.Args[1:] {
		fmt.Printf("Arg %d: %s\n", i+1, arg)
	}

	transcript.FetchTranscription(videoID, languages)

}
