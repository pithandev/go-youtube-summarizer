package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pithandev/go-youtube-summarizer/internal/chunker"
	"github.com/pithandev/go-youtube-summarizer/internal/transcript"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		log.Fatal("Error: Necessario o link ou o id de um video.")
	}
	videoID := os.Args[1]
	languages := make([]string, 1)
	languages[0] = "en"

	text, err := transcript.FetchTranscription(videoID, languages)
	if err != nil {
		log.Fatal(err)
	}

	chunked := chunker.Split(text, 5)
	for i, c := range chunked {
		fmt.Printf("%d, %v\n", i, c)
	}

}
