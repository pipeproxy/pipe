package build

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

func GenType(prefix string, typ reflect.Type, getTypeName func(reflect.Type) string) string {
	buf := bytes.NewBuffer(nil)
	g := genType{
		prefix:      prefix,
		out:         buf,
		getTypeName: getTypeName,
		todos:       []reflect.Type{},
		nameOnce:    map[string]struct{}{},
	}
	g.gen(typ)
	return buf.String()
}

type genType struct {
	prefix      string
	out         io.Writer
	getTypeName func(reflect.Type) string
	todos       []reflect.Type
	nameOnce    map[string]struct{}
}

func (g *genType) gen(typ reflect.Type) {
	g.toType(typ)
	todo := []reflect.Type{}
	for len(g.todos) != 0 {
		todo, g.todos = g.todos, todo
		for _, t := range todo {
			g.toType(t)
		}
		todo = todo[:0]
	}
}

func (g *genType) name(s string) string {
	return g.prefix + s
}

func (g *genType) toType(typ reflect.Type) {
	name := typ.Name()
	name = g.name(name)
	if _, ok := g.nameOnce[name]; ok {
		return
	}
	g.nameOnce[name] = struct{}{}

	if typ.Kind() == reflect.Interface {
		return
	}

	fmt.Fprintf(g.out, `type `)
	fmt.Fprintf(g.out, name)
	fmt.Fprintf(g.out, " ")
	g.to(typ, false)
}

func (g *genType) to(typ reflect.Type, define bool) {
	switch typ.Kind() {
	case reflect.Struct:
		if define {
			g.toOther(typ)
			g.todos = append(g.todos, typ)
		} else {
			g.toStruct(typ)
		}
	case reflect.Slice:
		g.toSlice(typ)
	case reflect.Ptr:
		g.to(typ.Elem(), define)
	case reflect.Interface:
		g.toInterface(typ)
	default:
		g.toOther(typ)
	}
}

func (g *genType) toStruct(typ reflect.Type) {
	fmt.Fprint(g.out, `struct {
`)
	num := typ.NumField()
	for i := 0; i != num; i++ {
		f := typ.Field(i)
		if f.Anonymous {
			continue
		}
		fmt.Fprint(g.out, "\t")
		fmt.Fprint(g.out, f.Name)

		fmt.Fprint(g.out, " ")
		g.to(f.Type, true)

		fmt.Fprint(g.out, "\n")
	}

	fmt.Fprint(g.out, `}
`)
}

func (g *genType) toSlice(typ reflect.Type) {
	fmt.Fprint(g.out, "[]")
	g.to(typ.Elem(), true)
}

func (g *genType) toOther(typ reflect.Type) {
	name := typ.Name()
	pkg := typ.PkgPath()
	if pkg != "" {
		g.todos = append(g.todos, typ)
		name = g.name(name)
	}
	fmt.Fprint(g.out, name)
}

func (g *genType) toInterface(typ reflect.Type) {
	typName := g.getTypeName(typ)
	fmt.Fprint(g.out, typName)
}
