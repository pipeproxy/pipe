package build

import (
	"reflect"
	"text/template"

	"github.com/wzshiming/namecase"
	"github.com/wzshiming/pipe/configure/alias"
)

func getTypeName(t reflect.Type) string {
	return getName(alias.GetType(t))
}

func getKindName(typName string) string {
	return getName(typName)
}

func getName(name string) string {
	return namecase.ToUpperHumpInitialisms(name)
}

var tempType = template.Must(template.New("_").
	Parse(`

type {{.Type}} interface {
	is{{.Type}}()
	PipeComponent
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

type Name{{.Type}} struct {
	Name string
	{{.Type}}
}

func (Name{{.Type}}) is{{.Type}}() {}
func (Name{{.Type}}) isPipeComponent() {}

func (n Name{{.Type}}) MarshalJSON() ([]byte, error) {
	data, err := n.{{.Type}}.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Name\":%q}", n.Name))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Name\":%q,", n.Name)), data[1:]...)
		}
	}

	return data, nil
}

type Ref{{.Type}} string

func (Ref{{.Type}}) is{{.Type}}() {}
func (Ref{{.Type}}) isPipeComponent() {}

func (m Ref{{.Type}}) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
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

// MarshalJSON returns m as the JSON encoding of m.
func (m {{.Name}}{{.Ref.Name}}) MarshalJSON() ([]byte, error) {
	const kind = "{{.Out.PkgPath}}.{{.Out.Name}}@{{.Kind}}"
	type t {{.Name}}{{.Ref.Name}}
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}
`))

var tempConfig = `

type PipeComponent interface {
	isPipeComponent()
	json.Marshaler
}

type RawPipeComponent []byte

func (RawPipeComponent) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawPipeComponent) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawPipeComponent) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawPipeComponent: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NamePipeComponent struct {
	Name string
	PipeComponent
}

func (NamePipeComponent) isPipeComponent() {}

// MarshalJSON returns n as the JSON encoding of n.
func (n NamePipeComponent) MarshalJSON() ([]byte, error) {
	data, err := n.PipeComponent.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Name\":%q}", n.Name))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Name\":%q,", n.Name)), data[1:]...)
		}
	}

	return data, nil
}

type RefPipeComponent string

func (RefPipeComponent) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of r.
func (r RefPipeComponent) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", r)), nil
}

`
