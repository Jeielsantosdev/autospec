# autospec - entrega completa do MVP

Este arquivo reúne tudo o que iremos criar na primeira entrega completa do `autospec`.

Objetivo desta versão:

- base CLI funcional
- configuração central
- geração inicial de OpenAPI
- integração inicial com Gin
- exposição de Swagger UI
- estrutura pronta para expansão

## Ordem sugerida de criação

1. `go.mod`
2. `cmd/autospec/main.go`
3. `internal/version/version.go`
4. `internal/config/config.go`
5. `internal/openapi/spec.go`
6. `internal/openapi/generator.go`
7. `internal/openapi/json.go`
8. `internal/runtime/inspect.go`
9. `internal/adapters/gin/adapter.go`
10. `internal/adapters/gin/routes.go`
11. `internal/generators/swaggerui.go`
12. `internal/commands/dev.go`
13. `internal/commands/watch.go`
14. `internal/autospec/app.go`

## Arquivo: go.mod

```go
module github.com/Jeielsantosdev/autospec

go 1.26.3

require github.com/gin-gonic/gin v1.10.0
```

## Arquivo: cmd/autospec/main.go

```go
package main

import (
	"fmt"
	"os"

	"github.com/Jeielsantosdev/autospec/internal/autospec"
)

func main() {
	app := autospec.New()

	if err := app.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
```

## Arquivo: internal/version/version.go

```go
package version

const value = "dev"

func String() string {
	return value
}
```

## Arquivo: internal/config/config.go

```go
package config

type Config struct {
	Name        string
	Version     string
	Title       string
	Description string
	ServerURL   string
	DocsPath    string
}

func Default() Config {
	return Config{
		Name:        "autospec",
		Version:     "dev",
		Title:       "autospec",
		Description: "Automatic OpenAPI generation for Go",
		ServerURL:   "http://localhost:8080",
		DocsPath:    "/docs",
	}
}
```

## Arquivo: internal/openapi/spec.go

```go
package openapi

type Spec struct {
	OpenAPI    string              `json:"openapi"`
	Info       Info                `json:"info"`
	Servers    []Server            `json:"servers,omitempty"`
	Paths      map[string]PathItem  `json:"paths"`
	Components Components          `json:"components,omitempty"`
}

type Info struct {
	Title       string `json:"title"`
	Version     string `json:"version"`
	Description string `json:"description,omitempty"`
}

type Server struct {
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
}

type Components struct {
	Schemas         map[string]Schema         `json:"schemas,omitempty"`
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes,omitempty"`
}

type SecurityScheme struct {
	Type         string `json:"type"`
	Scheme       string `json:"scheme,omitempty"`
	BearerFormat string `json:"bearerFormat,omitempty"`
	Description  string `json:"description,omitempty"`
}

type Schema struct {
	Type        string            `json:"type,omitempty"`
	Format      string            `json:"format,omitempty"`
	Description string            `json:"description,omitempty"`
	Properties  map[string]Schema `json:"properties,omitempty"`
	Items       *Schema           `json:"items,omitempty"`
	Required    []string          `json:"required,omitempty"`
	Example     any               `json:"example,omitempty"`
	Ref         string            `json:"$ref,omitempty"`
}

type PathItem struct {
	Get     *Operation `json:"get,omitempty"`
	Post    *Operation `json:"post,omitempty"`
	Put     *Operation `json:"put,omitempty"`
	Patch   *Operation `json:"patch,omitempty"`
	Delete  *Operation `json:"delete,omitempty"`
	Options *Operation `json:"options,omitempty"`
	Head    *Operation `json:"head,omitempty"`
	Trace   *Operation `json:"trace,omitempty"`
}

type Operation struct {
	Summary     string                `json:"summary,omitempty"`
	Description string                `json:"description,omitempty"`
	Tags        []string              `json:"tags,omitempty"`
	Security    []map[string][]string `json:"security,omitempty"`
	Parameters  []Parameter           `json:"parameters,omitempty"`
	RequestBody *RequestBody          `json:"requestBody,omitempty"`
	Responses   map[string]Response   `json:"responses,omitempty"`
}

type Parameter struct {
	Name        string `json:"name"`
	In          string `json:"in"`
	Required    bool   `json:"required,omitempty"`
	Description string `json:"description,omitempty"`
	Schema      Schema `json:"schema"`
}

type RequestBody struct {
	Required bool                 `json:"required,omitempty"`
	Content  map[string]MediaType `json:"content"`
}

type Response struct {
	Description string                `json:"description"`
	Content     map[string]MediaType  `json:"content,omitempty"`
}

type MediaType struct {
	Schema Schema `json:"schema"`
}
```

## Arquivo: internal/openapi/generator.go

```go
package openapi

func NewSpec(title string, version string, description string) *Spec {
	return &Spec{
		OpenAPI: "3.1.0",
		Info: Info{
			Title:       title,
			Version:     version,
			Description: description,
		},
		Paths: make(map[string]PathItem),
		Components: Components{
			Schemas:         make(map[string]Schema),
			SecuritySchemes: make(map[string]SecurityScheme),
		},
	}
}

func (s *Spec) AddServer(url string, description string) {
	s.Servers = append(s.Servers, Server{
		URL:         url,
		Description: description,
	})
}

func (s *Spec) AddSchema(name string, schema Schema) {
	if s.Components.Schemas == nil {
		s.Components.Schemas = make(map[string]Schema)
	}

	s.Components.Schemas[name] = schema
}

func (s *Spec) AddSecurityScheme(name string, scheme SecurityScheme) {
	if s.Components.SecuritySchemes == nil {
		s.Components.SecuritySchemes = make(map[string]SecurityScheme)
	}

	s.Components.SecuritySchemes[name] = scheme
}

func (s *Spec) AddOperation(path string, method string, operation Operation) {
	item := s.Paths[path]

	switch method {
	case "GET":
		item.Get = &operation
	case "POST":
		item.Post = &operation
	case "PUT":
		item.Put = &operation
	case "PATCH":
		item.Patch = &operation
	case "DELETE":
		item.Delete = &operation
	case "OPTIONS":
		item.Options = &operation
	case "HEAD":
		item.Head = &operation
	case "TRACE":
		item.Trace = &operation
	}

	s.Paths[path] = item
}

func (s *Spec) AddJSONResponse(path string, method string, status string, description string, schema Schema) {
	item := s.Paths[path]
	operation := operationForMethod(&item, method)
	if operation == nil {
		operation = &Operation{Responses: make(map[string]Response)}
	}

	if operation.Responses == nil {
		operation.Responses = make(map[string]Response)
	}

	operation.Responses[status] = Response{
		Description: description,
		Content: map[string]MediaType{
			"application/json": {
				Schema: schema,
			},
		},
	}

	switch method {
	case "GET":
		item.Get = operation
	case "POST":
		item.Post = operation
	case "PUT":
		item.Put = operation
	case "PATCH":
		item.Patch = operation
	case "DELETE":
		item.Delete = operation
	case "OPTIONS":
		item.Options = operation
	case "HEAD":
		item.Head = operation
	case "TRACE":
		item.Trace = operation
	}

	s.Paths[path] = item
}

func operationForMethod(item *PathItem, method string) *Operation {
	switch method {
	case "GET":
		return item.Get
	case "POST":
		return item.Post
	case "PUT":
		return item.Put
	case "PATCH":
		return item.Patch
	case "DELETE":
		return item.Delete
	case "OPTIONS":
		return item.Options
	case "HEAD":
		return item.Head
	case "TRACE":
		return item.Trace
	default:
		return nil
	}
}
```

## Arquivo: internal/openapi/json.go

```go
package openapi

import "encoding/json"

func (s *Spec) JSON() ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}

func (s *Spec) YAMLPlaceholder() string {
	return "YAML generation will be added in a later step"
}
```

## Arquivo: internal/runtime/inspect.go

```go
package runtime

type Route struct {
	Method      string
	Path        string
	Handler     string
	Middlewares []string
}

type Snapshot struct {
	Routes []Route
	Auth   bool
}

func Inspect() Snapshot {
	return Snapshot{
		Routes: make([]Route, 0),
	}
}

func DetectAuth(middlewares []string) bool {
	for _, middleware := range middlewares {
		if middleware == "AuthMiddleware" || middleware == "JWTMiddleware" {
			return true
		}
	}

	return false
}
```

## Arquivo: internal/adapters/gin/adapter.go

```go
package ginadapter

import (
	"github.com/gin-gonic/gin"

	"github.com/Jeielsantosdev/autospec/internal/openapi"
	"github.com/Jeielsantosdev/autospec/internal/runtime"
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
```

## Arquivo: internal/adapters/gin/routes.go

```go
package ginadapter

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Jeielsantosdev/autospec/internal/openapi"
	"github.com/Jeielsantosdev/autospec/internal/generators"
	"github.com/Jeielsantosdev/autospec/internal/runtime"
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
```

## Arquivo: internal/generators/swaggerui.go

```go
package generators

import "fmt"

func SwaggerUIPage(specURL string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>autospec docs</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.onload = () => {
      window.ui = SwaggerUIBundle({
        url: %q,
        dom_id: '#swagger-ui',
      })
    }
  </script>
</body>
</html>`, specURL)
}
```

## Arquivo: internal/commands/dev.go

```go
package commands

import "fmt"

func Dev() error {
	fmt.Println("autospec dev ainda nao implementado")
	return nil
}
```

## Arquivo: internal/commands/watch.go

```go
package commands

import "fmt"

func Watch() error {
	fmt.Println("autospec watch ainda nao implementado")
	return nil
}
```

## Arquivo: internal/autospec/app.go

```go
package autospec

import (
	"errors"
	"fmt"

	"github.com/Jeielsantosdev/autospec/internal/commands"
	"github.com/Jeielsantosdev/autospec/internal/config"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
	"github.com/Jeielsantosdev/autospec/internal/version"
)

type App struct {
	config config.Config
}

func New() *App {
	return &App{
		config: config.Default(),
	}
}

func (a *App) Spec() *openapi.Spec {
	spec := openapi.NewSpec(a.config.Title, a.config.Version, a.config.Description)
	spec.AddServer(a.config.ServerURL, "Local development server")
	spec.AddSecurityScheme("bearerAuth", openapi.SecurityScheme{
		Type:         "http",
		Scheme:       "bearer",
		BearerFormat: "JWT",
		Description:  "Bearer token authentication",
	})
	return spec
}

func (a *App) Run(args []string) error {
	if len(args) == 0 {
		fmt.Printf("%s %s\n", a.config.Name, version.String())
		fmt.Println("Use: autospec [version|dev|watch]")
		return nil
	}

	switch args[0] {
	case "dev":
		return commands.Dev()
	case "watch":
		return commands.Watch()
	case "version", "--version", "-v":
		fmt.Println(version.String())
		return nil
	default:
		return errors.New("comando desconhecido: " + args[0])
	}
}
```

## Exemplo de uso no Gin

```go
package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Jeielsantosdev/autospec/internal/autospec"
	ginadapter "github.com/Jeielsantosdev/autospec/internal/adapters/gin"
)

func main() {
	engine := gin.Default()
	app := autospec.New()
	spec := app.Spec()

	ginadapter.RegisterDocsRoutes(engine, spec)
	ginadapter.New(engine, spec).Attach()

	engine.Run(":8080")
}
```

## O que esta entrega cobre

- CLI base com `version`, `dev` e `watch`
- configuração centralizada
- estrutura OpenAPI 3.1
- geração de JSON da spec
- suporte inicial a security scheme Bearer JWT
- integração inicial com Gin
- página de Swagger UI pronta para servir
- base para runtime inspection e expansão futura

## Próximos arquivos que eu criaria depois desta etapa

- `internal/openapi/schema_builder.go`
- `internal/openapi/reflection.go`
- `internal/runtime/gin_snapshot.go`
- `internal/adapters/fiber/adapter.go`
- `internal/adapters/echo/adapter.go`
- `internal/cli/root.go`
- `internal/cli/dev.go`

## Resumo

Esta é a entrega completa inicial que eu faria agora para o `autospec`: uma base copiável, organizada por arquivo, com o núcleo do SDK, documentação automática inicial e integração pronta com Gin.