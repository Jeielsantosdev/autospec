package gin

import (
	"github.com/Jeielsantosdev/autospec/internal/openapi"
	"github.com/Jeielsantosdev/autospec/internal/runtime"
	"github.com/gin-gonic/gin"
)

type Adapter struct {
	Engine *gin.Engine
	Spec   *openapi.Spec
}

func New(engine *gin.Engine, spec *openapi.Spec) *Adapter {
	return &Adapter{
		Engine: engine,
		Spec:   spec,
	}
}

func (a *Adapter) Attach() {
	if a.Engine == nil || a.Spec == nil {
		return
	}

	snapshot := runtime.Inspect()
	for _, route := range snapshot.Routes {
		a.Spec.AddOperation(route.Path, route.Method, openapi.Operation{
			Summary: route.Handler,
			Tags:    route.Middlewares,
		})
	}
}
