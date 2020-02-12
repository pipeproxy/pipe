package build

import (
	"reflect"
	"strings"
	"text/template"

	"github.com/wzshiming/namecase"
)

var exceptionTypeName = map[string]string{
	"io.ReadCloser":  "input.Input",
	"io.WriteCloser": "output.Output",
}

func getTypeName(typName string) string {
	if t, ok := exceptionTypeName[typName]; ok {
		typName = t
	}

	n := strings.SplitN(typName, ".", 2)
	if len(n) == 1 {
		return n[0]
	}
	if strings.ToLower(n[0]) == strings.ToLower(n[1]) {
		return n[1]
	}

	return getName(n[0]) + getName(n[1])
}

func getKindName(typName string) string {
	return getName(typName)
}

func getName(name string) string {
	return namecase.ToUpperHump(name)
}

var tempType = template.Must(template.New("_").
	Parse(`

type {{.Type}} interface {
	is{{.Type}}()
}

type Raw{{.Type}} []byte

func (Raw{{.Type}}) is{{.Type}}() {}
func (Raw{{.Type}}) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m Raw{{.Type}}) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *Raw{{.Type}}) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("Raw{{.Type}}: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}
`))

func tempKindGenType(prefix string, typ reflect.Type) string {
	return GenType(prefix, typ, getTypeName)
}

var tempKind = template.Must(template.New("_").
	Funcs(template.FuncMap{"genType": tempKindGenType}).
	Parse(`


// {{.Name}}{{.Ref.Name}} {{.Out.PkgPath}}.{{.Out.Name}}@{{.Kind}}
{{genType .Name .Ref}}

func ({{.Name}}{{.Ref.Name}}) is{{.Type}}() {}
func ({{.Name}}{{.Ref.Name}}) isPipeComponent() {}

func (m {{.Name}}{{.Ref.Name}}) MarshalJSON() ([]byte, error) {
	const kind = "{{.Out.PkgPath}}.{{.Out.Name}}@{{.Kind}}"
	type t {{.Name}}{{.Ref.Name}}
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte("{\"@Kind\":\"" + kind + "\"}")
		} else {
			data = append([]byte("{\"@Kind\":\""+kind+"\","), data[1:]...)
		}
	}
	return data, nil
}
`))
