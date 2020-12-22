package register

import (
	"reflect"
	"strings"

	"github.com/wzshiming/funcfg/types"
	"github.com/wzshiming/funcfg/types/extra"
	"github.com/wzshiming/logger"
)

func Register(kind string, fun interface{}) error {
	name, err := getKindName(kind, fun)
	if err != nil {
		logger.Log.Error(err, "GetKindName", "kind", kind)
		return err
	}

	err = types.Default.Register(name, fun)
	if err != nil {
		logger.Log.Error(err, "Register", "kind", kind, "name", name)
		return err
	}
	return nil
}

func getKindName(kind string, fun interface{}) (string, error) {
	typ, err := types.CheckFunc(reflect.ValueOf(fun))
	if err != nil {
		return "", err
	}
	name := typ.Name()
	pkg := typ.PkgPath()
	if pkg != "" {
		p := strings.Split(pkg, "components/")
		if len(p) != 1 {
			pkg = p[len(p)-1]
		}
		name = strings.Join([]string{pkg, name}, ".")
	}
	return strings.Join([]string{kind, name}, "@"), nil
}

func RegisterWithBuildFunc(name string, f, i interface{}) error {
	fun, err := extra.BuildFunc(f, i)
	if err != nil {
		logger.Log.Error(err, "RegisterWithBuildFunc", "name", name)
		return err
	}
	return Register(name, fun.Interface())
}
