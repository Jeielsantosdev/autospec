package autospec

import "reflect"

type Plugin interface {
	Name() string
	OnRoute(route RouteMeta)
	OnSample(sample RequestSample)
	SchemaFromType(t reflect.Type) (map[string]any, error)
}
