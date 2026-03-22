package gemini

import (
	"google.golang.org/genai"
)

// ExtractSources returns the list of source URLs from Google Search grounding metadata.
// Returns an empty slice (not nil) if no grounding metadata is present.
// Per D-05, D-06: caller prints "Sources: none" when this returns empty.
func ExtractSources(resp *genai.GenerateContentResponse) []string {
	var urls []string
	for _, cand := range resp.Candidates {
		if cand.GroundingMetadata == nil {
			continue
		}
		for _, chunk := range cand.GroundingMetadata.GroundingChunks {
			if chunk.Web != nil && chunk.Web.URI != "" {
				urls = append(urls, chunk.Web.URI)
			}
		}
	}
	if urls == nil {
		return []string{}
	}
	return urls
}
