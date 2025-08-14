package sectionwrapper

import (
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

var markdown = goldmark.New(
	goldmark.WithExtensions(
		&SectionWrapper{},
	),
)

type TestCase struct {
	desc string
	md   string
	html string
}

var cases = [...]TestCase{
	{
		desc: "Basic section wrapper with a paragraph",
		md:   `# Section 1
This is a paragraph in section 1.`,
		html: `<section class="section-h1"><h1>Section 1</h1>
<p>This is a paragraph in section 1.</p>
</section>`,
	},
	{
		desc: "Basic section wrapper with a two paragraphs",
		md:   `# Section 1
This is a paragraph in section 1.

This is a second paragraph in section 1.`,
		html: `<section class="section-h1"><h1>Section 1</h1>
<p>This is a paragraph in section 1.</p>
<p>This is a second paragraph in section 1.</p>
</section>`,
	},
	{
		desc: "Wrap multiple sections -- same level",
		md:   `## Section 1

This is a paragraph in section 1.

## Section 2

This is a paragraph in section 2.`,
		html: `<section class="section-h2"><h2>Section 1</h2>
<p>This is a paragraph in section 1.</p>
</section><section class="section-h2"><h2>Section 2</h2>
<p>This is a paragraph in section 2.</p>
</section>`,
	},
	{
		desc: "Wrap multiple sections -- nested levels",
		md:   `# Title

## Section 1

This is a paragraph in section 1.

## Section 2

This is a paragraph in section 2.

### Section 2.1

This is a paragraph in section 2.1.

## Section 3

This is a paragraph in section 3.
`,
		html: `<section class="section-h1"><h1>Title</h1>
<section class="section-h2"><h2>Section 1</h2>
<p>This is a paragraph in section 1.</p>
</section><section class="section-h2"><h2>Section 2</h2>
<p>This is a paragraph in section 2.</p>
<section class="section-h3"><h3>Section 2.1</h3>
<p>This is a paragraph in section 2.1.</p>
</section></section><section class="section-h2"><h2>Section 3</h2>
<p>This is a paragraph in section 3.</p>
</section></section>`,
	},
}

func TestSectionWrapper(t *testing.T) {
	for i, c := range cases {
		testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
			No:          i,
			Description: c.desc,
			Markdown:    c.md,
			Expected:    c.html,
		}, t)
	}
}
