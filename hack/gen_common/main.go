package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"

	_ "github.com/wzshiming/pipe/init"

	"github.com/wzshiming/pipe/components/common/types"
	"github.com/wzshiming/pipe/internal/logger"
)

//go:generate ../../bin/go-bindata --nomemcopy --pkg main -o template.go ./template.go.tpl

func main() {
	for _, g := range types.Global {
		t := reflect.TypeOf(g).Elem()
		g := newGen(t)
		err := g.Write()
		if err != nil {
			logger.Error(err)
		}
	}
}

func newGen(typ reflect.Type) *gen {
	return &gen{typ: typ}
}

type gen struct {
	typ     reflect.Type
	imports map[string]string
}

func (g *gen) data() interface{} {

	imports := map[string]struct{}{}
	dep := []string{
		"github.com/wzshiming/pipe/components/common/register",
		"github.com/wzshiming/pipe/internal/logger",
		"github.com/wzshiming/pipe/internal/ctxcache",
		"context",
	}
	for _, d := range dep {
		imports[d] = struct{}{}
	}
	methods := []interface{}{}
	pkgPath := g.typ.PkgPath()
	imports[pkgPath] = struct{}{}
	j := g.typ.NumMethod()
	for i := 0; i != j; i++ {
		m := g.typ.Method(i)

		args := []interface{}{}
		inNum := m.Type.NumIn()

		for i := 0; i != inNum; i++ {
			typ := m.Type.In(i)
			imports[typ.PkgPath()] = struct{}{}
			name := strings.ToLower(typ.Name())
			if name == "" {
				name = string([]byte{'a' + byte(i)})
			}
			t := typ.String()
			m := map[string]string{
				"Name": name,
				"Type": t,
			}
			args = append(args, m)
		}
		results := []interface{}{}
		outNum := m.Type.NumOut()
		for i := 0; i != outNum; i++ {
			typ := m.Type.Out(i)
			imports[typ.PkgPath()] = struct{}{}
			name := strings.ToLower(typ.Name())
			t := typ.String()
			m := map[string]string{
				"Name": name,
				"Type": t,
			}
			if t == "error" {
				imports["fmt"] = struct{}{}
				m["Value"] = `fmt.Errorf("error none")`
			}
			results = append(results, m)
		}
		methods = append(methods, map[string]interface{}{
			"FuncName": m.Name,
			"Args":     args,
			"Results":  results,
		})
	}

	imp := []map[string]string{}
	for i := range imports {
		if i == "" {
			continue
		}
		im := map[string]string{
			"PkgPath": i,
			"Alias":   "",
		}
		//if strings.Count(i, "/") > 1 {
		//	im["Alias"] = getImportName(i)
		//}
		imp = append(imp, im)
	}
	return map[string]interface{}{
		"Type":    g.typ.Name(),
		"PkgName": strings.ToLower(g.typ.Name()),
		"Pkg":     getImportName(g.typ.PkgPath()),
		"Imports": imp,
		"Methods": methods,
	}
}

func getImportName(p string) string {
	p = filepath.Base(p)
	p = strings.SplitN(p, ".", 2)[0]
	return p
}

func (g *gen) Write() error {
	buf := bytes.NewBuffer(nil)
	data := g.data()
	err := temp.Execute(buf, data)
	if err != nil {
		return err
	}
	src := buf.Bytes()
	newSrc, err := format.Source(src)
	if err == nil {
		src = newSrc
	} else {
		logger.Error(err)
	}

	file := filepath.Join(g.typ.PkgPath(), strings.ToLower(g.typ.Name()), "init.go")
	os.MkdirAll(filepath.Dir(file), 0755)
	return ioutil.WriteFile(file, src, 0644)
}

var (
	temp = template.Must(template.New("_").Parse(string(MustAsset("template.go.tpl"))))
)
