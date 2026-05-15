package ginadapter

import (
	"net/http"

	"github.com/Jeielsantosdev/autospec/internal/generators"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
	"github.com/gin-gonic/gin"
)

func Detect(app any) bool {
	_, ok := app.(*gin.Engine)
	return ok
}

func Attach(app any, spec *openapi.Spec) error {
	engine, ok := app.(*gin.Engine)
	if !ok {
		return nil
	}

	engine.GET("/openapi.json", func(c *gin.Context) {
		payload, err := spec.JSON()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Data(http.StatusOK, "application/json", payload)
	})

	engine.GET("/docs", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(generators.SwaggerUIPage("/openapi.json")))
	})

	engine.Use(CaptureMiddleware(spec))

	return nil
}

func RegisterRouteWithMeta(app any, method, path string, handler any, spec *openapi.Spec, meta any) error {
	engine := app.(*gin.Engine)
	// aqui só exemplo: registrar o handler na rota
	switch method {
	case "GET":
		engine.GET(path, handler.(gin.HandlerFunc))
	case "POST":
		engine.POST(path, handler.(gin.HandlerFunc))
	default:
		engine.Handle(method, path, handler.(gin.HandlerFunc))
	}
	// registrar placeholder em spec
	spec.AddOperation(path, method, openapi.Operation{Summary: "(declared)", Responses: map[string]openapi.Response{"200": {Description: "OK"}}})
	return nil
}
