package utils

import (
	"fmt"
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	mdhtml "github.com/gomarkdown/markdown/html"
	"github.com/microcosm-cc/bluemonday"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

var (
	htmlFormatter  *html.Formatter
	highlightStyle *chroma.Style
)

func init() {
	htmlFormatter = html.New(html.WithClasses(true), html.TabWidth(2))
	if htmlFormatter == nil {
		panic("couldn't create html formatter")
	}

	styleName := "catppuccin-macchiato"

	highlightStyle = styles.Get(styleName)
	if highlightStyle == nil {
		panic(fmt.Sprintf("didn't find style '%s'", styleName))
	}
}

func MdToHtml(content []byte) []byte {
	renderer := createRenderer()

	unsafeHtml := markdown.ToHTML(content, nil, renderer)

	p := bluemonday.UGCPolicy()
	p.AllowStyling()

	html := p.SanitizeBytes(unsafeHtml)

	return html
}

func createRenderer() *mdhtml.Renderer {
	opts := mdhtml.RendererOptions{
		Flags:          mdhtml.CommonFlags,
		RenderNodeHook: renderHook,
	}

	return mdhtml.NewRenderer(opts)
}

func renderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if code, ok := node.(*ast.CodeBlock); ok {
		renderCode(w, code, entering)
		return ast.GoToNext, true
	}

	return ast.GoToNext, false
}

func renderCode(w io.Writer, codeBlock *ast.CodeBlock, entering bool) {
	defaultLang := ""

	lang := string(codeBlock.Info)

	htmlHighlight(w, string(codeBlock.Literal), lang, defaultLang)
}

func htmlHighlight(w io.Writer, source, lang, defaultLang string) error {
	if lang == "" {
		lang = defaultLang
	}

	l := lexers.Get(lang)
	if l == nil {
		l = lexers.Analyse(source)
	}

	if l == nil {
		l = lexers.Fallback
	}

	l = chroma.Coalesce(l)

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}

	return htmlFormatter.Format(w, highlightStyle, it)
}
