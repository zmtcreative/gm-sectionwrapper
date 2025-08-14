package sectionwrapper

import (
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

type TestCase struct {
	desc string
	md   string
	html string
}

var mdStandard = goldmark.New(
	goldmark.WithExtensions(
		&SectionWrapper{},
	),
)

var mdWithHeadingClass = goldmark.New(
	goldmark.WithExtensions(
		&SectionWrapper{
            useHeadingClass: true,
        },
	),
)

func TestSectionWrapperStandard(t *testing.T) {
	for i, c := range casesStandard {
		testutil.DoTestCase(mdStandard, testutil.MarkdownTestCase{
			No:          i,
			Description: c.desc,
			Markdown:    c.md,
			Expected:    c.html,
		}, t)
	}
}

func TestSectionWrapperWithHeadingClass(t *testing.T) {
	for i, c := range casesWithHeadingClass {
		testutil.DoTestCase(mdWithHeadingClass, testutil.MarkdownTestCase{
			No:          i,
			Description: c.desc,
			Markdown:    c.md,
			Expected:    c.html,
		}, t)
	}
}

var casesWithHeadingClass = [...]TestCase{
	{
		desc: "Basic section wrapper with a paragraph",
		md:   `# Section 1
This is a paragraph in section 1.`,
		html: `<section class="section-h1 h1"><h1>Section 1</h1>
<p>This is a paragraph in section 1.</p>
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
		html: `<section class="section-h1 h1"><h1>Title</h1>
<section class="section-h2 h2"><h2>Section 1</h2>
<p>This is a paragraph in section 1.</p>
</section><section class="section-h2 h2"><h2>Section 2</h2>
<p>This is a paragraph in section 2.</p>
<section class="section-h3 h3"><h3>Section 2.1</h3>
<p>This is a paragraph in section 2.1.</p>
</section></section><section class="section-h2 h2"><h2>Section 3</h2>
<p>This is a paragraph in section 3.</p>
</section ></section>`,
	},
}

var casesStandard = [...]TestCase{
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
    {
        desc: "Empty document",
        md:   ``,
        html: ``,
    },
    {
        desc: "Content without headings",
        md:   `This is just a paragraph.

And another paragraph.`,
        html: `<p>This is just a paragraph.</p>
<p>And another paragraph.</p>`,
    },
    {
        desc: "Heading without content",
        md:   `# Empty Section`,
        html: `<section class="section-h1"><h1>Empty Section</h1>
</section>`,
    },
    {
        desc: "Multiple empty headings",
        md:   `# Section 1
## Section 2
### Section 3`,
        html: `<section class="section-h1"><h1>Section 1</h1>
<section class="section-h2"><h2>Section 2</h2>
<section class="section-h3"><h3>Section 3</h3>
</section></section></section>`,
    },
    {
        desc: "Content before first heading",
        md:   `This is content before any heading.

# First Heading

This is content after the heading.`,
        html: `<p>This is content before any heading.</p>
<section class="section-h1"><h1>First Heading</h1>
<p>This is content after the heading.</p>
</section>`,
    },
    {
        desc: "Skip heading levels",
        md:   `# Level 1

### Level 3 (skipped 2)

This is in level 3.

## Level 2

This is in level 2.`,
        html: `<section class="section-h1"><h1>Level 1</h1>
<section class="section-h3"><h3>Level 3 (skipped 2)</h3>
<p>This is in level 3.</p>
</section><section class="section-h2"><h2>Level 2</h2>
<p>This is in level 2.</p>
</section></section>`,
    },
    {
        desc: "All heading levels (1-6)",
        md:   `# H1
## H2
### H3
#### H4
##### H5
###### H6
Content at deepest level.`,
        html: `<section class="section-h1"><h1>H1</h1>
<section class="section-h2"><h2>H2</h2>
<section class="section-h3"><h3>H3</h3>
<section class="section-h4"><h4>H4</h4>
<section class="section-h5"><h5>H5</h5>
<section class="section-h6"><h6>H6</h6>
<p>Content at deepest level.</p>
</section></section></section></section></section></section>`,
    },
    {
        desc: "Complex nesting with various content types",
        md:   `# Main Section

Some introduction text.

## Subsection A

- List item 1
- List item 2

### Deep subsection

> This is a blockquote

## Subsection B

` + "```" + `
code block
` + "```" + `

# Another Main Section

Final paragraph.`,
        html: `<section class="section-h1"><h1>Main Section</h1>
<p>Some introduction text.</p>
<section class="section-h2"><h2>Subsection A</h2>
<ul>
<li>List item 1</li>
<li>List item 2</li>
</ul>
<section class="section-h3"><h3>Deep subsection</h3>
<blockquote>
<p>This is a blockquote</p>
</blockquote>
</section></section><section class="section-h2"><h2>Subsection B</h2>
<pre><code>code block
</code></pre>
</section></section><section class="section-h1"><h1>Another Main Section</h1>
<p>Final paragraph.</p>
</section>`,
    },
    {
        desc: "Multiple consecutive headings at different levels",
        md:   `## First
### Second
#### Third
### Fourth
## Fifth`,
        html: `<section class="section-h2"><h2>First</h2>
<section class="section-h3"><h3>Second</h3>
<section class="section-h4"><h4>Third</h4>
</section></section><section class="section-h3"><h3>Fourth</h3>
</section></section><section class="section-h2"><h2>Fifth</h2>
</section>`,
    },
    {
        desc: "Starting with lower-level heading",
        md:   `### Starting with H3

Content here.

## Then H2

More content.

# Finally H1

Final content.`,
        html: `<section class="section-h3"><h3>Starting with H3</h3>
<p>Content here.</p>
</section><section class="section-h2"><h2>Then H2</h2>
<p>More content.</p>
</section><section class="section-h1"><h1>Finally H1</h1>
<p>Final content.</p>
</section>`,
    },}

