package chunker

import "strings"

type Chunk struct {
	ID   int
	Text string
}

func Split(text string, wordsPerChunk int) []Chunk {
	words := strings.Fields(text)

	var chunks []Chunk

	for start := 0; start < len(words); start += wordsPerChunk {
		end := start + wordsPerChunk

		if end > len(words) {
			end = len(words)
		}

		chunkText := strings.Join(words[start:end], " ")

		chunkID := 1

		chunk := Chunk{
			ID:   chunkID,
			Text: chunkText,
		}

		chunks = append(chunks, chunk)
		chunkID++
	}

	return chunks
}
