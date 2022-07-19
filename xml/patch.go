package xml

func Patch(elm *XMLElement, newText *string) {
	switch velm := (*elm).(type) {
	case *containerElement:
		parsed := ParseXML(*newText)
		velm.chlidren = parsed.chlidren
	default:
		return
	}
}
