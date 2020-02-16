package template

import (
	"io"
	"net/http"
)

type Text string

func (t Text) Format(w io.Writer, r *http.Request) error {
	_, err := io.WriteString(w, string(t))
	return err
}

func (t Text) FormatString(r *http.Request) string {
	return string(t)
}
