package autospec

import (
	"errors"

	"github.com/Jeielsantosdev/autospec/internal/adapters/ginadapter"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
)

type Autospec struct {
	Spec             *openapi.Spec
	plugins          []Plugin
	preferredAdapter string
}

type Meta struct {
	Input  any
	Output any
}

type RouteMeta struct {
	Method string
	Path   string
}

type RequestSample struct {
	Method       string
	Path         string
	RequestBody  []byte
	ResponseBody []byte
	Status       int
}

func New(opts ...Option) *Autospec {
	a := &Autospec{}
	for _, o := range opts {
		o(a)
	}
	if a.Spec == nil {
		a.Spec = openapi.NewSpec("autospec", "dev", "autospec generated spec")
		a.Spec.AddServer("http://localhost:8080", "local")
	}
	return a
}

func (a *Autospec) AttachTo(app any) error {
	if ginadapter.Detect(app) {
		return ginadapter.Attach(app, a.Spec)
	}
	return errors.New("no adapter found for given app instance")
}

func (a *Autospec) Handle(app any, method, path string, handler any, meta Meta) error {
	if ginadapter.Detect(app) {
		return ginadapter.RegisterRouteWithMeta(app, method, path, handler, a.Spec, meta)
	}
	return errors.New("handle: unsupported adapter")
}

func (a *Autospec) Run(args []string) error {
	if len(args) == 0 {
		println("autospec dev")
		return nil
	}
	switch args[0] {
	case "version", "-v", "--version":
		println(a.Spec.Info.Version)
		return nil
	default:
		return errors.New("unknown command: " + args[0])
	}

}
