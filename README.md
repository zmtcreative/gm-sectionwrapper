# Goldmark Section Wrapper Extension

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ZMT-Creative/gm-sectionwrapper)
![GitHub License](https://img.shields.io/github/license/ZMT-Creative/gm-sectionwrapper)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/ZMT-Creative/gm-sectionwrapper)
![GitHub Tag](https://img.shields.io/github/v/tag/ZMT-Creative/gm-sectionwrapper?include_prereleases&sort=semver)

This is a Goldmark extension that automatically wraps headings and their content in HTML `<section>` elements with proper nesting. This extension transforms your Markdown document structure into semantic HTML sections, making it easier to style and navigate.

## Installation

```bash
go get github.com/ZMT-Creative/gm-sectionwrapper
```

## Basic Usage

```go
package main

import (
    "bytes"
    "fmt"

    "github.com/yuin/goldmark"
    sectionwrapper "github.com/ZMT-Creative/gm-sectionwrapper"
)

func main() {
    md := goldmark.New(
        goldmark.WithExtensions(
            sectionwrapper.NewSectionWrapper(),
        ),
    )

    source := `# Main Section
This is some content.

## Subsection
More content here.`

    var buf bytes.Buffer
    if err := md.Convert([]byte(source), &buf); err != nil {
        panic(err)
    }
    fmt.Print(buf.String())
}
```

**Output:**

```html
<section class="section-h1"><h1>Main Section</h1>
<p>This is some content.</p>
<section class="section-h2"><h2>Subsection</h2>
<p>More content here.</p>
</section></section>
```

### Alternative Initialization Method

If you just want the default behavior without any options (*i.e., just use the `section-h?` class
in the HTML output*) you can also initialize the extension using the following method:

```go
    md := goldmark.New(
        goldmark.WithExtensions(sectionwrapper.SectionWrapper),
    )
```

## Configuration Options

The extension provides several configuration options through functional options:

### WithSectionClass(enabled bool)

Controls whether to add `section-h{level}` classes to section elements.

- **Default:** `true`
- **Example:** `section-h1`, `section-h2`, etc.

```go
// Disable section-h{level} classes
md := goldmark.New(
    goldmark.WithExtensions(
        sectionwrapper.NewSectionWrapper(
            sectionwrapper.WithSectionClass(false),
        ),
    ),
)
```

### WithHeadingClass(enabled bool)

Controls whether to add `h{level}` classes to section elements.

- **Default:** `false`
- **Example Class Entry:** `h1`, `h2`, etc.

```go
// Enable heading-level classes
md := goldmark.New(
    goldmark.WithExtensions(
        sectionwrapper.NewSectionWrapper(
            sectionwrapper.WithHeadingClass(true),
        ),
    ),
)
```

### WithCustomClassPrefix(prefix string)

Adds a custom prefix followed by the heading level to section elements.

- **Default:** `""` (empty)
- **Example Class Entry:** With prefix `"custom-"` => `custom-h1`, `custom-h2`, etc.

```go
// Add custom prefix
md := goldmark.New(
    goldmark.WithExtensions(
        sectionwrapper.NewSectionWrapper(
            sectionwrapper.WithCustomClassPrefix("my-"),
        ),
    ),
)
```

### WithCustomClass(class string)

Adds a fixed custom class to all section elements regardless of heading level.

- **Default:** `""` (empty)
- **Example Class Entry:** `"content-section"`

```go
// Add custom class to all sections
md := goldmark.New(
    goldmark.WithExtensions(
        sectionwrapper.NewSectionWrapper(
            sectionwrapper.WithCustomClass("content-section"),
        ),
    ),
)
```

### Combining Options

You can combine multiple options:

```go
md := goldmark.New(
    goldmark.WithExtensions(
        sectionwrapper.NewSectionWrapper(
            sectionwrapper.WithSectionClass(true),
            sectionwrapper.WithHeadingClass(true),
            sectionwrapper.WithCustomClassPrefix("article-"),
            sectionwrapper.WithCustomClass("content"),
        ),
    ),
)
```

This would produce sections with classes like: `"section-h1 h1 article-h1 content"`

## Behavior

### Nesting

Sections are properly nested based on heading hierarchy:

```markdown
# Level 1
Content for level 1
## Level 2
Content for level 2
### Level 3
Content for level 3
## Another Level 2
More content
```

Produces properly nested `<section>` elements where Level 3 is inside Level 2, which is inside Level 1.

### Content Handling

- All content between headings is included in the appropriate section
- Content before the first heading remains outside any section
- Empty headings create empty sections
- Supports all Markdown content types (paragraphs, lists, code blocks, blockquotes, etc.)

### Heading Levels

- Supports all heading levels (H1 through H6)
- Handles skipped heading levels gracefully
- Maintains proper nesting regardless of heading level jumps

## Examples

### Default Configuration

```markdown
# Main Title
Introduction text.
## Section One
Section content.
```

```html
<section class="section-h1"><h1>Main Title</h1>
<p>Introduction text.</p>
<section class="section-h2"><h2>Section One</h2>
<p>Section content.</p>
</section></section>
```

### With All Options Enabled

```go
sectionwrapper.NewSectionWrapper(
    sectionwrapper.WithSectionClass(true),
    sectionwrapper.WithHeadingClass(true),
    sectionwrapper.WithCustomClassPrefix("doc-"),
    sectionwrapper.WithCustomClass("section"),
)
```

```html
<section class="section-h1 h1 doc-h1 section"><h1>Main Title</h1>
<p>Introduction text.</p>
<section class="section-h2 h2 doc-h2 section"><h2>Section One</h2>
<p>Section content.</p>
</section></section>
```

## Compatibility Warning

Unit tests have been created with this extension, and it currently passes all the tests. This
includes tests with the built-in Goldmark extensions (`extension.GFM`, `extension.DefinitionList`,
and `extension.Footnote`). However, this does not guarantee the extension will work with any other
extensions.

> [!Important]
>
> This extension transforms the AST structure by wrapping headings in section nodes, which could
> potentially conflict with other extensions that also modify heading behavior or document
> structure.

If you encounter issues when using this extension alongside others, please test with the extensions
individually to identify conflicts.

## License

This project is licensed under the MIT License. See the [LICENSE.md](LICENSE.md) file for details.
