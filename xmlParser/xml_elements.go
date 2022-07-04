package xmlParser

import "fmt"

type XMLElement interface {
	GetContent() string
	GetFather() XMLElement
	GetChildren() []XMLElement
	GetTag() string
	GetAttributes() string
}

type textElement struct {
	text   string
	father XMLElement
	XMLElement
}

func (txt *textElement) GetContent() string {
	return txt.text
}
func (txt *textElement) GetFather() XMLElement {
	return txt.father
}
func (txt *textElement) GetChildren() []XMLElement {
	return make([]XMLElement, 0)
}
func (txt *textElement) GetTag() string {
	return ""
}
func (txt *textElement) GetAttributes() string {
	return ""
}

type nonContainerElement struct {
	father     XMLElement
	tag        string
	attributes string
}

func (nelm *nonContainerElement) GetContent() string {
	if nelm.attributes == "" {
		return fmt.Sprintf("<%s />", nelm.tag)
	}
	return fmt.Sprintf("<%s %s />", nelm.tag, nelm.attributes)
}
func (nelm *nonContainerElement) GetFather() XMLElement {
	return nelm.father
}
func (nelm *nonContainerElement) GetChildren() []XMLElement {
	return make([]XMLElement, 0)
}
func (nelm *nonContainerElement) GetTag() string {
	return nelm.tag
}
func (nelm *nonContainerElement) GetAttributes() string {
	return nelm.attributes
}

type containerElement struct {
	father     XMLElement
	chlidren   []XMLElement
	tag        string
	attributes string
}

func (elm *containerElement) getInnerContent() string {
	content := ""
	for _, child := range elm.GetChildren() {
		content += child.GetContent()
	}
	return content
}
func (elm *containerElement) GetContent() string {
	if elm.attributes == "" {
		return fmt.Sprintf("<%[1]s>%s</%[1]s>", elm.tag, elm.getInnerContent())
	}
	return fmt.Sprintf("<%[1]s %s>%s</%[1]s>", elm.tag, elm.attributes, elm.getInnerContent())
}
func (elm *containerElement) GetFather() XMLElement {
	return elm.father
}
func (elm *containerElement) GetChildren() []XMLElement {
	return elm.chlidren
}
func (elm *containerElement) GetTag() string {
	return elm.tag
}
func (elm *containerElement) GetAttributes() string {
	return elm.attributes
}

type XML struct {
	Prolog   string
	chlidren []XMLElement
}

func (xml *XML) GetContent() string {
	if len(xml.chlidren) == 1 {
		return xml.chlidren[0].GetContent()
	}
	content := ""
	for _, child := range xml.GetChildren() {
		content += child.GetContent()
	}
	return content
}
func (xml *XML) GetFather() XMLElement {
	var nullfather XMLElement
	return nullfather
}
func (xml *XML) GetChildren() []XMLElement {
	if len(xml.chlidren) == 1 {
		return xml.chlidren[0].GetChildren()
	}
	return xml.chlidren
}
func (xml *XML) GetTag() string {
	if len(xml.chlidren) == 1 {
		return xml.chlidren[0].GetTag()
	}
	return ""
}
func (xml *XML) GetAttributes() string {
	if len(xml.chlidren) == 1 {
		return xml.chlidren[0].GetAttributes()
	}
	return ""
}
