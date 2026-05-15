package gin

import (
	"net/http"

	"github.com/Jeielsantosdev/autospec/internal/generators"
	"github.com/Jeielsantosdev/autospec/internal/openapi"

	"github.com/Jeielsantosdev/autospec/internal/runtime"
	"github.com/gin-gonic/gin"
)

func RegisterDocsRoutes(engine *gin.Engine, spec *openapi.Spec) {
	engine.GET("/openapi.json", func(ctx *gin.Context) {
		payload, err := spec.JSON()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Data(http.StatusOK, "application/json", payload)
	})
	engine.GET("/docs", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(generators.SwaggerUIPage("/openapi.json")))
	})
}

func RegisterRuntimeSnapshot(engine *gin.Engine, spec *openapi.Spec) {
	snapshot := runtime.Inspect()
	for _, route := range snapshot.Routes {
		spec.AddOperation(route.Path, route.Method, openapi.Operation{
			Summary:     route.Handler,
			Description: "Route discovered at runtime",
			Tags:        route.Middlewares,
		})
	}
	_ = engine
}
