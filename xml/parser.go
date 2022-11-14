package xml

import (
	"regexp"
	"strings"

	"goxml.com/utility"
)

type rawElement struct {
	document *string
	start    int
	end      int
	children []XMLElement
}

func (relm *rawElement) getContent() string {
	return (*relm.document)[relm.start:relm.end]
}

func (relm *rawElement) addChild(elm XMLElement) {
	relm.children = append(relm.children, elm)
}

func createRawElement(document *string, start int, end int) rawElement {
	relm := rawElement{document: document, start: start, end: end}
	relm.children = make([]XMLElement, 0, 5)
	return relm
}

func ParseXML(text string) XML {
	stack := utility.NewStack[*rawElement]()
	r, _ := regexp.Compile(`(?:<(\??\w*)\s*((?:[\w:-]*?=".*?"\s*)*)\s*/?\w*\??>)+?`)
	indices := r.FindAllStringIndex(text, -1)
	xml := XML{}
	rawXML := createRawElement(&text, 0, 0)
	stack.Push(&rawXML)
	for i, index := range indices {
		tag := text[index[0]:index[1]]
		if i > 0 && index[0] != indices[i-1][1] { // text element
			element := textElement{text: text[indices[i-1][1]:index[0]]}
			father, err := stack.Pop()
			if err == nil {
				father.addChild(&element)
				stack.Push(father)
			}
		}
		switch {
		case strings.HasSuffix(tag, "?>"): // prolog
			xml.Prolog = tag
		case strings.HasSuffix(tag, "/>"): // non-container tag
			tagType := r.FindStringSubmatch(tag)[1]
			attributes := strings.Trim(r.FindStringSubmatch(tag)[2], " ")
			element := nonContainerElement{attributes: attributes, tag: tagType}
			father, err := stack.Pop()
			if err == nil {
				father.addChild(&element)
				stack.Push(father)
			}
		case strings.HasPrefix(tag, "</"): // closing tag
			openrawElement, err := stack.Pop()
			if err == nil {
				tagType := r.FindStringSubmatch(openrawElement.getContent())[1]
				attributes := strings.Trim(r.FindStringSubmatch(openrawElement.getContent())[2], " ")
				element := containerElement{attributes: attributes, tag: tagType}
				element.chlidren = openrawElement.children
				for _, child := range element.chlidren {
					switch v := child.(type) {
					case *textElement:
						v.father = &element
					case *nonContainerElement:
						v.father = &element
					case *containerElement:
						v.father = &element
					}
				}
				father, err := stack.Pop()
				if err == nil {
					father.addChild(&element)
					stack.Push(father)
				}
			}
		default: // opening tag
			relm := createRawElement(&text, index[0], index[1])
			stack.Push(&relm)
		}
	}
	firstInsertedRawXML, err := stack.Pop()
	if err == nil {
		xml.chlidren = firstInsertedRawXML.children
	}
	return xml
}
