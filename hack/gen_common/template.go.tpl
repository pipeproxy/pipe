// DO NOT EDIT! Code generated.
package reference

import (
	{{range .Imports}}
    {{- .Alias}} "{{.PkgPath}}"
	{{end}}
)

func init() {
	register.Register("ref", New{{.Type}}RefWithConfig)
	register.Register("def", New{{.Type}}DefWithConfig)
	register.Register("none", New{{.Type}}None)
}

{{$Type := .Type}}
{{$Pkg := .Pkg}}

type Config struct {
	Name string
	Def  {{.Pkg}}.{{.Type}} `json:",omitempty"`
}

func New{{.Type}}RefWithConfig(conf *Config) ({{.Pkg}}.{{.Type}}, error) {
	o := &{{.Type}}{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func New{{.Type}}DefWithConfig(conf *Config) ({{.Pkg}}.{{.Type}}, error) {
	{{.Type}}Store[conf.Name] = conf.Def
	return conf.Def, nil
}

var {{.Type}}Store = map[string]{{.Pkg}}.{{.Type}}{}

func {{.Type}}Find(name string, defaults {{.Pkg}}.{{.Type}}) {{.Pkg}}.{{.Type}} {
	o, ok := {{.Type}}Store[name]
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return {{.Type}}None{}
}

type {{.Type}}None struct{}

func New{{.Type}}None() {{.Pkg}}.{{.Type}} {
	return {{.Type}}None{}
}

{{range .Methods}}
	func ({{$Type}}None) {{.FuncName}}(
    {{- range .Args -}}
		_ {{.Type}},
    {{- end -}}
	) (
    {{- range .Results -}}
		_ {{.Type}},
    {{- end -}}
	)  {
	logger.Warn("this is none of {{$Pkg}}.{{$Type}}")
	return
	}
{{end}}

type {{.Type}} struct {
	Name string
	Def  {{.Pkg}}.{{.Type}}
}

{{range .Methods}}
	func (o *{{$Type}}) {{.FuncName}}(
    {{- range .Args -}}
        {{- .Name }} {{.Type}},
    {{- end -}}
	) (
    {{- range .Results -}}
        {{.Type}},
    {{- end -}}
	) {
    {{if .Results}}return{{end}} {{$Type}}Find(o.Name, o.Def).{{.FuncName}}(
    {{- range .Args -}}
        {{- .Name }},
    {{- end -}}
	)
	}
{{end}}