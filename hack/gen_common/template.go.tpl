// DO NOT EDIT! Code generated.
package {{ .PkgName }}

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

func New{{ .Type }}RefWithConfig(ctx context.Context, conf *Config) {{ .Pkg }}.{{ .Type }} {
	o := &{{ .Type }}{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func New{{ .Type }}DefWithConfig(ctx context.Context, conf *Config) {{ .Pkg }}.{{ .Type }} {
	return {{ .Type }}Put(ctx, conf.Name, conf.Def)
}

func {{ .Type }}Put(ctx context.Context, name string, def {{ .Pkg }}.{{ .Type }}) {{ .Pkg }}.{{ .Type }} {
	if def == nil {
		return {{ .Type }}None
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return {{ .Type }}None
	}
	store, _ := m.LoadOrStore("{{ .Pkg }}.{{ .Type }}", map[string]{{ .Pkg }}.{{ .Type }}{})
	store.(map[string]{{ .Pkg }}.{{ .Type }})[name] = def
	return def
}

func {{ .Type }}Get(ctx context.Context, name string, defaults {{ .Pkg }}.{{ .Type }}) {{ .Pkg }}.{{ .Type }} {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("{{ .Pkg }}.{{ .Type }}")
		if ok {
			o, ok := store.(map[string]{{ .Pkg }}.{{ .Type }})[name]
			if ok {
				return o
			}   
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.Warnf("{{ $Pkg }}.{{ $Type }} %q is not defined", name)
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
	Ctx  context.Context
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
	{{ if .Results }}return{{ end }} {{ $Type }}Get(o.Ctx, o.Name, o.Def).{{ .FuncName }}(
	{{- range .Args -}}
		{{- .Name }},
	{{- end -}}
	)
}
{{ end }}