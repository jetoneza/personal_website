package utils

import (
	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"
)

func MdToHtml(content []byte) []byte {
	unsafeHtml := markdown.ToHTML(content, nil, nil)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafeHtml)

	return html
}
