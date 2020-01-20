package template

import (
	"io"
	"net/http"
	"unsafe"
)

type Text string

func (t Text) Format(w io.Writer, r *http.Request) error {
	_, err := w.Write(*(*[]byte)(unsafe.Pointer(&t)))
	return err
}

func (t Text) FormatString(r *http.Request) string {
	return string(t)
}
