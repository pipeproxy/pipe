package register

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/wzshiming/funcfg/types"
	"github.com/wzshiming/funcfg/types/extra"
	"github.com/wzshiming/logger"
)

func Register(kind string, fun interface{}) error {
	kind, err := GetKindName(kind, fun)
	if err != nil {
		return err
	}

	return types.Default.Register(kind, fun)
}

func GetKindName(kind string, fun interface{}) (string, error) {
	typ, err := types.CheckFunc(reflect.ValueOf(fun))
	if err != nil {
		logger.Log.Error(err, "CheckFunc", "kind", kind)
		return "", err
	}
	pkg := typ.PkgPath()
	if pkg == "" {
		return fmt.Sprintf("%s@%s", kind, typ.Name()), nil
	}
	p := strings.Split(pkg, "components/")
	if len(p) != 1 {
		pkg = p[len(p)-1]
	}
	return fmt.Sprintf("%s@%s.%s", kind, pkg, typ.Name()), nil
}

func RegisterWithBuildFunc(name string, f, i interface{}) error {
	fun, err := extra.BuildFunc(f, i)
	if err != nil {
		return err
	}
	return Register(name, fun.Interface())
}
