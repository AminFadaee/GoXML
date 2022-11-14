package main

import (
	"fmt"
	"goxml.com/xml"
)

func main() {
	xmlSample := `<bridge>
					<question id="1">What is your name?</question>
					<question id="2">What is your quest?</question>
					<question id="3">What is your favorite colour?</question>
				 </bridge>`
	x := xml.ParseXML(xmlSample)
	for _, element := range x.GetChildren() {
		if element.GetTag() == "question" {
			fmt.Println(element.GetChildren()[0].GetContent())
		}
	}
	/* Prints:
	What is your name?
	What is your quest?
	What is your favorite colour?
	*/
}
