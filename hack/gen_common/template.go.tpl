// DO NOT EDIT! Code generated.
package reference

import (
	{{ range .Imports }}
    {{- .Alias }} "{{ .PkgPath }}"
	{{ end }}
)

func init() {
	register.Register("ref", New{{ .Type }}RefWithConfig)
	register.Register("def", New{{ .Type }}DefWithConfig)
	register.Register("none", new{{ .Type }}None)
}

{{ $Type := .Type }}
{{ $Pkg := .Pkg }}

type Config struct {
	Name string
	Def  {{ .Pkg }}.{{ .Type }} `json:",omitempty"`
}

func New{{ .Type }}RefWithConfig(conf *Config) {{ .Pkg }}.{{ .Type }} {
	o := &{{ .Type }}{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func New{{ .Type }}DefWithConfig(conf *Config) {{ .Pkg }}.{{ .Type }} {
	return {{ .Type }}Put(conf.Name, conf.Def)
}

var (
    mut sync.RWMutex
    _{{ .Type }}Store = map[string]{{ .Pkg }}.{{ .Type }}{}
)

func {{ .Type }}Put(name string, def {{ .Pkg }}.{{ .Type }}) {{ .Pkg }}.{{ .Type }} {
    if def == nil {
        def = {{ .Type }}None
    }
    mut.Lock()
	_{{ .Type }}Store[name] = def
	mut.Unlock()
	return def
}

func {{ .Type }}Get(name string, defaults {{ .Pkg }}.{{ .Type }}) {{ .Pkg }}.{{ .Type }} {
    mut.RLock()
	o, ok := _{{ .Type }}Store[name]
	mut.RUnlock()
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return {{ .Type }}None
}

var {{ .Type }}None _{{ .Type }}None

type _{{ .Type }}None struct{}

func new{{ .Type }}None() {{ .Pkg }}.{{ .Type }} {
	return {{ .Type }}None
}

{{ range .Methods }}
func (_{{ $Type }}None) {{ .FuncName }}(
    {{- range .Args -}}
		_ {{ .Type }},
    {{- end -}}
	) (
    {{- range .Results -}}
		{{ if .Value }} {{ .Name }} {{ else }} _ {{ end }} {{ .Type }},
    {{- end -}}
	) {
	logger.Warn("this is none of {{ $Pkg }}.{{ $Type }}")
	{{ range .Results }}
    {{ if .Value }} {{ .Name }} = {{ .Value }} {{ end }}
    {{ end }}
	return
}
{{ end }}

type {{ .Type }} struct {
	Name string
	Def  {{ .Pkg }}.{{ .Type }}
}

{{ range .Methods }}
func (o *{{ $Type }}) {{ .FuncName }}(
    {{- range .Args -}}
        {{- .Name }} {{ .Type }},
    {{- end -}}
	) (
    {{- range .Results -}}
        {{ .Type }},
    {{- end -}}
	) {
    {{ if .Results }}return{{ end }} {{ $Type }}Get(o.Name, o.Def).{{ .FuncName }}(
    {{- range .Args -}}
        {{- .Name }},
    {{- end -}}
	)
}
{{ end }}