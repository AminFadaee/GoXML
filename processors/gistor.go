package processors

import (
	"fmt"
	"golang.org/x/exp/slices"
	"kirby.lensreader.com/utility"
	xmlUtil "kirby.lensreader.com/xml"
	"strings"
)

func gistParagraph(content *string) string {
	indices := utility.IndeciesInString(*content, ".", "!", "?")
	endDotIndex := len(*content) - len(`</p>`) - 1
	if !slices.Contains(indices, endDotIndex) {
		indices = append(indices, endDotIndex)
	}
	if len(indices) <= 2 {
		return *content
	}
	firstSentEnd, LastSentStart := indices[0], indices[len(indices)-2]+1
	summary := fmt.Sprintf("<<%d>>", len((*content)[firstSentEnd+1:LastSentStart]))
	return (*content)[:firstSentEnd+1] + summary + (*content)[LastSentStart:]
}

func getGist(xml xmlUtil.XML) {
	stack := utility.NewStack[xmlUtil.XMLElement]()
	stack.Push(&xml)
	for stack.IsNotEmpty() {
		c, _ := stack.Pop()
		if strings.ToLower(c.GetTag()) == "p" {
			text := c.GetContent()
			newText := gistParagraph(&text)
			xmlUtil.Patch(c, &newText)
		}
		for _, elm := range c.GetChildren() {
			stack.Push(elm)
		}
	}
}
