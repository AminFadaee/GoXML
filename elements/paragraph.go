package elements

import (
	"fmt"

	"golang.org/x/exp/slices"
	"kirby.lensreader.com/utility"
)

type Paragraph struct {
	Content string
}

func (paragraph *Paragraph) Gist() Paragraph {
	indices := utility.IndeciesInString(paragraph.Content, ".", "!", "?")
	endDotIndex := len(paragraph.Content) - len(`</p>`) - 1
	if !slices.Contains(indices, endDotIndex) {
		indices = append(indices, endDotIndex)
	}
	fmt.Println(paragraph, indices)
	if len(indices) <= 2 {
		return *paragraph
	}
	firstSentEnd, LastSentStart := indices[0], indices[len(indices)-2]+1
	summary := fmt.Sprintf("<<%d>>", len(paragraph.Content[firstSentEnd+1:LastSentStart]))
	return Paragraph{paragraph.Content[:firstSentEnd+1] + summary + paragraph.Content[LastSentStart:]}
}
