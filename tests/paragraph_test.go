package tests

import (
	"fmt"
	"testing"

	"kirby.lensreader.com/elements"
)

type Case struct {
	input    string
	expected string
}

func (testcase *Case) errorLog(base, actual string) string {
	return fmt.Sprintf("%s\n\nCase:\n=========\n%s\nExpected:\n=========\n%s\nActual:\n=========\n%s", base, testcase.input, testcase.expected, actual)
}

func TestParagraphGist(t *testing.T) {
	samples := [...]Case{
		{`<p></p>`, `<p></p>`},
		{`<p>First Sentence.</p>`, `<p>First Sentence.</p>`},
		{`<p>First Sentence. Second Sentence.</p>`, `<p>First Sentence. Second Sentence.</p>`},
		{`<p>First Sentence. Second Sentence. Third Sentence.</p>`, `<p>First Sentence.<<17>> Third Sentence.</p>`},
		{`<p>First Sentence. Second Sentence. Third Sentence. Fourth Sentence.</p>`, `<p>First Sentence.<<33>> Fourth Sentence.</p>`},
		{`<p>First Sentence! Second Sentence? Third Sentence. Fourth Sentence!</p>`, `<p>First Sentence!<<33>> Fourth Sentence!</p>`},
		{`<p>First Sentence. Second Sentence. Third Sentence. Missed Dot</p>`, `<p>First Sentence.<<33>> Missed Dot</p>`},
	}
	for _, sample := range samples {
		paragraph := elements.Paragraph{Content: sample.input}
		actual := paragraph.Gist().Content
		if actual != sample.expected {
			t.Error(sample.errorLog("Gist is not as expected!", actual))
		}
	}
}
