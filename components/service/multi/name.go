package multi

import (
	"fmt"
	"reflect"

	"github.com/pipeproxy/pipe/components/service"
)

func getName(i int, svc service.Service) string {
	val := reflect.ValueOf(svc)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() == reflect.Struct {
		v := val.FieldByName("Name")
		if v.IsValid() && v.Kind() == reflect.String {
			return v.String()
		}
	}
	return fmt.Sprintf("service-%d", i)
}
