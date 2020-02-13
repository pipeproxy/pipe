package template

import (
	"io"
	"net/http"
	"text/template/parse"
)

type Format interface {
	Format(w io.Writer, r *http.Request) error
	FormatString(r *http.Request) string
}

func NewFormat(text string) (Format, error) {
	t := &Template{}
	err := t.init(text)
	if err != nil {
		return nil, err
	}
	if t.template.Root.NodeType == parse.NodeText ||
		(t.template.Root.NodeType == parse.NodeList &&
			len(t.template.Root.Nodes) == 1 &&
			t.template.Root.Nodes[0].Type() == parse.NodeText) {
		return Text(text), nil
	}
	return t, nil
}
