package processors

import (
	"fmt"
	"golang.org/x/exp/slices"
	"io/ioutil"
	"kirby.lensreader.com/utility"
	xmlUtil "kirby.lensreader.com/xml"
	"strings"
)

func gistParagraph(content *string) string {
	indices := utility.IndeciesInString(*content, ". ", "! ", "? ")
	endDotIndex := len(*content) - len(`</p>`) - 1
	if !slices.Contains(indices, endDotIndex) {
		indices = append(indices, endDotIndex)
	}
	if len(indices) <= 2 {
		return *content
	}
	firstSentEnd, LastSentStart := indices[0], indices[len(indices)-2]+1
	summary := "<span style=\"color:#EEEEEE;display:inline\">" + (*content)[firstSentEnd+1:LastSentStart] + "</span>"
	return (*content)[:firstSentEnd+1] + summary + (*content)[LastSentStart:]
}

func getXMLGist(xml xmlUtil.XML) xmlUtil.XML {
	stack := utility.NewStack[xmlUtil.XMLElement]()
	stack.Push(&xml)
	for stack.IsNotEmpty() {
		c, _ := stack.Pop()
		if strings.ToLower(c.GetTag()) == "p" {
			text := c.GetContent()
			newText := gistParagraph(&text)
			xmlUtil.Patch(&c, &newText)
		} else {
			for _, elm := range c.GetChildren() {
				stack.Push(elm)
			}
		}
	}
	return xml
}

func GistEpub(path string, destination string) {
	epub := "temp/" + utility.MD5(path) + "/"
	err := utility.Unzip(path, epub)
	if err != nil {
		fmt.Println(err)
	}
	meta := utility.GetFileContent(epub + "/META-INF/container.xml")
	if strings.Contains(meta, "application/oebps-package+xml") {
		files, _ := ioutil.ReadDir(epub + "/OEBPS")
		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".htm") {
				fullPath := epub + "OEBPS/" + file.Name()
				xml := xmlUtil.ParseXML(utility.GetFileContent(fullPath))
				gist := getXMLGist(xml)
				utility.SaveToFile(gist.GetContent(), fullPath)
			}
		}
	}
	utility.Zip(epub, destination)
}
