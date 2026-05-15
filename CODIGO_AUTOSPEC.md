Vou atualizar este documento com código e instruções copiáveis: core, adapter Gin, middleware de captura, exemplos e estrutura.

------

# autospec — implementação pronta para copiar

Este arquivo contém uma proposta DX-first e implementações mínimas copiáveis para acelerar a integração do `autospec` em apps Go com Gin. Copie os blocos de código para os arquivos correspondentes.

Objetivo: plug-and-play, não invasivo, zero-config por padrão, com opção declarativa para precisão.

------

## Estrutura de pastas recomendada

- cmd/autospec/main.go
- internal/autospec/           # core público (API simples: New, Attach)
	- autospec.go
	- options.go
	- plugin.go
- internal/adapters/gin/
	- attach.go
	- capture_middleware.go
- internal/openapi/
	- spec.go        (já presente — use o seu)
	- generator.go
- internal/inspector/
	- inspector.go
- internal/reflect/
	- schema_builder.go
- internal/generators/
	- swaggerui.go   (já presente — use o seu)

------

## API pública — exemplos copy/paste

Zero-config (apenas inclua e anexe):

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Jeielsantosdev/autospec/internal/autospec"
)

func main() {
	r := gin.Default()
	as := autospec.New() // cria spec e sistema de plugins internamente
	as.AttachTo(r)       // autodetecta Gin e registra /docs + /openapi.json

	r.GET("/hello", func(c *gin.Context){ c.JSON(200, gin.H{"msg":"ok"}) })
	r.POST("/users", CreateUser)

	r.Run(":8080")
}
```

Declarative (opcional — para precisão):

```go
as.Handle(r, "POST", "/users", CreateUser, autospec.Meta{
	Input:  autospec.TypeOf[CreateUserRequest](),
	Output: autospec.TypeOf[User](),
})
```

Helpers (opcionais, para permitir inferência precisa dentro do handler):

```go
func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	autospec.BindJSON(c, &req) // grava tipo/amostra
	user := doCreate(req)
	autospec.RespondJSON(c, 201, user) // grava amostra de resposta
}
```

------

## Implementações copiáveis (arquivos essenciais)

Coloque os arquivos a seguir exatamente nos caminhos indicados.

### File: cmd/autospec/main.go

```go
package mainVou atualizar este documento com código e instruções copiáveis: core, adapter Gin, middleware de captura, exemplos e estrutura.

------

# autospec — implementação pronta para copiar

Este arquivo contém uma proposta DX-first e implementações mínimas copiáveis para acelerar a integração do `autospec` em apps Go com Gin. Copie os blocos de código para os arquivos correspondentes.

Objetivo: plug-and-play, não invasivo, zero-config por padrão, com opção declarativa para precisão.

------

## Estrutura de pastas recomendada

- cmd/autospec/main.go
- internal/autospec/           # core público (API simples: New, Attach)
	- autospec.go
	- options.go
	- plugin.go
- internal/adapters/gin/
	- attach.go
	- capture_middleware.go
- internal/openapi/
	- spec.go        (já presente — use o seu)
	- generator.go
- internal/inspector/
	- inspector.go
- internal/reflect/
	- schema_builder.go
- internal/generators/
	- swaggerui.go   (já presente — use o seu)

------

## API pública — exemplos copy/paste

Zero-config (apenas inclua e anexe):

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Jeielsantosdev/autospec/internal/autospec"
)

func main() {
	r := gin.Default()
	as := autospec.New() // cria spec e sistema de plugins internamente
	as.AttachTo(r)       // autodetecta Gin e registra /docs + /openapi.json

	r.GET("/hello", func(c *gin.Context){ c.JSON(200, gin.H{"msg":"ok"}) })
	r.POST("/users", CreateUser)

	r.Run(":8080")
}
```

Declarative (opcional — para precisão):

```go
as.Handle(r, "POST", "/users", CreateUser, autospec.Meta{
	Input:  autospec.TypeOf[CreateUserRequest](),
	Output: autospec.TypeOf[User](),
})
```

Helpers (opcionais, para permitir inferência precisa dentro do handler):

```go
func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	autospec.BindJSON(c, &req) // grava tipo/amostra
	user := doCreate(req)
	autospec.RespondJSON(c, 201, user) // grava amostra de resposta
}
```

------

## Implementações copiáveis (arquivos essenciais)

Coloque os arquivos a seguir exatamente nos caminhos indicados.

### File: cmd/autospec/main.go

```go
package main

import (
	"fmt"
	"os"

	"github.com/Jeielsantosdev/autospec/internal/autospec"
)

func main() {
	app := autospec.New()
	if err := app.RVou atualizar este documento com código e instruções copiáveis: core, adapter Gin, middleware de captura, exemplos e estrutura.

------

# autospec — implementação pronta para copiar

Este arquivo contém uma proposta DX-first e implementações mínimas copiáveis para acelerar a integração do `autospec` em apps Go com Gin. Copie os blocos de código para os arquivos correspondentes.

Objetivo: plug-and-play, não invasivo, zero-config por padrão, com opção declarativa para precisão.

------

## Estrutura de pastas recomendada

- cmd/autospec/main.go
- internal/autospec/           # core público (API simples: New, Attach)
	- autospec.go
	- options.go
	- plugin.go
- internal/adapters/gin/
	- attach.go
	- capture_middleware.go
- internal/openapi/
	- spec.go        (já presente — use o seu)
	- generator.go
- internal/inspector/
	- inspector.go
- internal/reflect/
	- schema_builder.go
- internal/generators/
	- swaggerui.go   (já presente — use o seu)

------

## API pública — exemplos copy/paste

Zero-config (apenas inclua e anexe):

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Jeielsantosdev/autospec/internal/autospec"
)

func main() {
	r := gin.Default()
	as := autospec.New() // cria spec e sistema de plugins internamente
	as.AttachTo(r)       // autodetecta Gin e registra /docs + /openapi.json

	r.GET("/hello", func(c *gin.Context){ c.JSON(200, gin.H{"msg":"ok"}) })
	r.POST("/users", CreateUser)

	r.Run(":8080")
}
```

Declarative (opcional — para precisão):

```go
as.Handle(r, "POST", "/users", CreateUser, autospec.Meta{
	Input:  autospec.TypeOf[CreateUserRequest](),
	Output: autospec.TypeOf[User](),
})
```

Helpers (opcionais, para permitir inferência precisa dentro do handler):

```go
func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	autospec.BindJSON(c, &req) // grava tipo/amostra
	user := doCreate(req)
	autospec.RespondJSON(c, 201, user) // grava amostra de resposta
}
```

------

## Implementações copiáveis (arquivos essenciais)

Coloque os arquivos a seguir exatamente nos caminhos indicados.

### File: cmd/autospec/main.go

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

### File: internal/autospec/options.go

```go
package autospec

type Option func(*Autospec)

func WithPlugin(p Plugin) Option {
	return func(a *Autospec) { a.plugins = append(a.plugins, p) }
}

func WithAdapterName(name string) Option {
	return func(a *Autospec) { a.preferredAdapter = name }
}
```

### File: internal/autospec/plugin.go

```go
package autospec

import "reflect"

type Plugin interface {
	Name() string
	OnRoute(route RouteMeta)
	OnSample(sample RequestSample)
	SchemaFromType(t reflect.Type) (map[string]any, error)
}
```

### File: internal/autospec/autospec.go

```go
package autospec

import (
	"errors"

	ginadapter "github.com/Jeielsantosdev/autospec/internal/adapters/gin"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
)

type Autospec struct {
	Spec            *openapi.Spec
	plugins         []Plugin
	preferredAdapter string
}

func New(opts ...Option) *Autospec {
	a := &Autospec{}
	for _, o := range opts { o(a) }
	if a.Spec == nil {
		a.Spec = openapi.NewSpec("autospec", "dev", "autospec generated spec")
		a.Spec.AddServer("http://localhost:8080", "local")
	}
	return a
}

// AttachTo tenta detectar o framework e anexar automaticamente
func (a *Autospec) AttachTo(app any) error {
	// Gin
	if ginadapter.Detect(app) {
		return ginadapter.Attach(app, a.Spec)
	}

	return errors.New("no adapter found for given app instance")
}

// Handle: opção declarativa para registrar rota + metadados
func (a *Autospec) Handle(app any, method, path string, handler any, meta Meta) error {
	// Para simplicidade, delegamos a adapters quando necessário.
	if ginadapter.Detect(app) {
		return ginadapter.RegisterRouteWithMeta(app, method, path, handler, a.Spec, meta)
	}
	return errors.New("handle: unsupported adapter")
}

// Run minimal CLI surface
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

// Meta (declarative) e tipos de amostra
type Meta struct {
	Input any
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

// Helpers para uso em handlers (pequeno wrapper)
func BindJSON(ctx any, dest any) error { // implementado por adapters via type switch
	// placeholder: adapters podem prover BindJSON mais preciso
	return nil
}

func RespondJSON(ctx any, status int, body any) error {
	return nil
}
```

### File: internal/adapters/gin/attach.go

```go
package ginadapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
	"github.com/Jeielsantosdev/autospec/internal/generators"
)

// Detect tenta identificar se `app` é um *gin.Engine
func Detect(app any) bool {
	_, ok := app.(*gin.Engine)
	return ok
}

// Attach registra rotas de docs e instala middleware de captura
func Attach(app any, spec *openapi.Spec) error {
	engine, ok := app.(*gin.Engine)
	if !ok {
		return nil
	}

	// registrar docs
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

	// instalar middleware de captura (coleta amostras para inferência)
	engine.Use(CaptureMiddleware(spec))

	return nil
}

// RegisterRouteWithMeta permite registro declarativo
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
```

### File: internal/adapters/gin/capture_middleware.go

```go
package ginadapter

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// CaptureMiddleware captura request e response e adiciona exemplo simples à spec
func CaptureMiddleware(spec *openapi.Spec) gin.HandlerFunc {
	return func(c *gin.Context) {
		// capturar request
		var reqBuf []byte
		if c.Request.Body != nil {
			b, _ := io.ReadAll(c.Request.Body)
			reqBufVou atualizar este documento com código e instruções copiáveis: core, adapter Gin, middleware de captura, exemplos e estrutura.

------

# autospec — implementação pronta para copiar

Este arquivo contém uma proposta DX-first e implementações mínimas copiáveis para acelerar a integração do `autospec` em apps Go com Gin. Copie os blocos de código para os arquivos correspondentes.

Objetivo: plug-and-play, não invasivo, zero-config por padrão, com opção declarativa para precisão.

------

## Estrutura de pastas recomendada

- cmd/autospec/main.go
- internal/autospec/           # core público (API simples: New, Attach)
	- autospec.go
	- options.go
	- plugin.go
- internal/adapters/gin/
	- attach.go
	- capture_middleware.go
- internal/openapi/
	- spec.go        (já presente — use o seu)
	- generator.go
- internal/inspector/
	- inspector.go
- internal/reflect/
	- schema_builder.go
- internal/generators/
	- swaggerui.go   (já presente — use o seu)

------

## API pública — exemplos copy/paste

Zero-config (apenas inclua e anexe):

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Jeielsantosdev/autospec/internal/autospec"
)

func main() {
	r := gin.Default()
	as := autospec.New() // cria spec e sistema de plugins internamente
	as.AttachTo(r)       // autodetecta Gin e registra /docs + /openapi.json

	r.GET("/hello", func(c *gin.Context){ c.JSON(200, gin.H{"msg":"ok"}) })
	r.POST("/users", CreateUser)

	r.Run(":8080")
}
```

Declarative (opcional — para precisão):

```go
as.Handle(r, "POST", "/users", CreateUser, autospec.Meta{
	Input:  autospec.TypeOf[CreateUserRequest](),
	Output: autospec.TypeOf[User](),
})
```

Helpers (opcionais, para permitir inferência precisa dentro do handler):

```go
func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	autospec.BindJSON(c, &req) // grava tipo/amostra
	user := doCreate(req)
	autospec.RespondJSON(c, 201, user) // grava amostra de resposta
}
```

------

## Implementações copiáveis (arquivos essenciais)

Coloque os arquivos a seguir exatamente nos caminhos indicados.

### File: cmd/autospec/main.go

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

### File: internal/autospec/options.go

```go
package autospec

type Option func(*Autospec)

func WithPlugin(p Plugin) Option {
	return func(a *Autospec) { a.plugins = append(a.plugins, p) }
}

func WithAdapterName(name string) Option {
	return func(a *Autospec) { a.preferredAdapter = name }
}
```

### File: internal/autospec/plugin.go

```go
package autospec

import "reflect"

type Plugin interface {
	Name() string
	OnRoute(route RouteMeta)
	OnSample(sample RequestSample)
	SchemaFromType(t reflect.Type) (map[string]any, error)
}
```

### File: internal/autospec/autospec.go

```go
package autospec

import (
	"errors"

	ginadapter "github.com/Jeielsantosdev/autospec/internal/adapters/gin"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
)

type Autospec struct {
	Spec            *openapi.Spec
	plugins         []Plugin
	preferredAdapter string
}

func New(opts ...Option) *Autospec {
	a := &Autospec{}
	for _, o := range opts { o(a) }
	if a.Spec == nil {
		a.Spec = openapi.NewSpec("autospec", "dev", "autospec generated spec")
		a.Spec.AddServer("http://localhost:8080", "local")
	}
	return a
}

// AttachTo tenta detectar o framework e anexar automaticamente
func (a *Autospec) AttachTo(app any) error {
	// Gin
	if ginadapter.Detect(app) {
		return ginadapter.Attach(app, a.Spec)
	}

	return errors.New("no adapter found for given app instance")
}

// Handle: opção declarativa para registrar rota + metadados
func (a *Autospec) Handle(app any, method, path string, handler any, meta Meta) error {
	// Para simplicidade, delegamos a adapters quando necessário.
	if ginadapter.Detect(app) {
		return ginadapter.RegisterRouteWithMeta(app, method, path, handler, a.Spec, meta)
	}
	return errors.New("handle: unsupported adapter")
}

// Run minimal CLI surface
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

// Meta (declarative) e tipos de amostra
type Meta struct {
	Input any
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

// Helpers para uso em handlers (pequeno wrapper)
func BindJSON(ctx any, dest any) error { // implementado por adapters via type switch
	// placeholder: adapters podem prover BindJSON mais preciso
	return nil
}

func RespondJSON(ctx any, status int, body any) error {
	return nil
}
```

### File: internal/adapters/gin/attach.go

```go
package ginadapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
	"github.com/Jeielsantosdev/autospec/internal/generators"
)

// Detect tenta identificar se `app` é um *gin.Engine
func Detect(app any) bool {
	_, ok := app.(*gin.Engine)
	return ok
}

// Attach registra rotas de docs e instala middleware de captura
func Attach(app any, spec *openapi.Spec) error {
	engine, ok := app.(*gin.Engine)
	if !ok {
		return nil
	}

	// registrar docs
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

	// instalar middleware de captura (coleta amostras para inferência)
	engine.Use(CaptureMiddleware(spec))

	return nil
}

// RegisterRouteWithMeta permite registro declarativo
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
```

### File: internal/adapters/gin/capture_middleware.go

```go
package ginadapter

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// CaptureMiddleware captura request e response e adiciona exemplo simples à spec
func CaptureMiddleware(spec *openapi.Spec) gin.HandlerFunc {
	return func(c *gin.Context) {
		// capturar request
		var reqBuf []byte
		if c.Request.Body != nil {
			b, _ := io.ReadAll(c.Request.Body)
			reqBuf = b
			c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
		}

		// capturar response
		bw := &bodyWriter{ResponseWriter: c.Writer, body: bytes.NewBuffer(nil)}
		c.Writer = bw

		c.Next()

		status := c.Writer.Status()
		path := c.FullPath()
		method := c.Request.Method

		// anexar exemplo simples na spec: usamos example sem schema real
		// para uma evolução futura, integrar reflect/schema_builder
		if path != "" {
			example := map[string]any{"request": string(reqBuf), "response": bw.body.String()}
			schema := openapi.Schema{Type: "object", Example: example}
			spec.AddJSONResponse(path, method, http.StatusText(status), "captured example", schema)
		}
	}
}
```

### File: internal/reflect/schema_builder.go (esqueleto)

```go
package reflect

// Implementação futura: reflect.Type -> JSON Schema
// Forneça utilitários para transformar structs em openapi.Schema
```

### File: internal/inspector/inspector.go (esqueleto)

```go
package inspector

// Aqui ficará a lógica de amostragem, merge de exemplos e deduplicação.
// Para a v0 inicial usamos a CaptureMiddleware que registra exemplos diretamente na spec.
```

------

## Fluxo interno resumido

- `autospec.New()` cria um `openapi.Spec` mínimo.
- `as.AttachTo(app)` detecta framework via adapters e chama `Attach` do adapter.
- Adapter registra `/openapi.json` e `/docs` e instala um `CaptureMiddleware` (coleta amostras sem alterar handlers).
- Middleware captura request/response e insere exemplos em `spec`.
- Opcional: `autospec.Handle(...)` registra rotas declarativamente e adiciona metadados precisos.

------

## Roadmap técnico curto

- v0.1: Core + Gin adapter + capture middleware + Swagger UI.
- v0.2: Schema builder reflect + sample merge + masking de campos sensíveis.
- v0.3: Plugin system, AuthDetector plugin, middleware detector.
- v0.4: Adapters: Echo, Fiber, Chi, net/http.

------

## Problemas técnicos e recomendações

- Inferência apenas via amostras é heurística; oferecer combinador declarativo (Meta) para precisão.
- Limitar tamanho das amostras por padrão; mascarar campos sensíveis.
- Fornecer modo "offline" (análise estática com go/packages) para projetos que preferem zero-runtime-overhead.

------

Se quiser, implemento agora esses arquivos no repositório do workspace (criar/atualizar arquivos), e então executo um build Go rápido para validar. Quer que eu faça isso agora? = b
			c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
		}

		// capturar response
		bw := &bodyWriter{ResponseWriter: c.Writer, body: bytes.NewBuffer(nil)}
		c.Writer = bw

		c.Next()

		status := c.Writer.Status()
		path := c.FullPath()
		method := c.Request.Method

		// anexar exemplo simples na spec: usamos example sem schema real
		// para uma evolução futura, integrar reflect/schema_builder
		if path != "" {
			example := map[string]any{"request": string(reqBuf), "response": bw.body.String()}
			schema := openapi.Schema{Type: "object", Example: example}
			spec.AddJSONResponse(path, method, http.StatusText(status), "captured example", schema)
		}
	}
}
```

### File: internal/reflect/schema_builder.go (esqueleto)

```go
package reflect

// Implementação futura: reflect.Type -> JSON Schema
// Forneça utilitários para transformar structs em openapi.Schema
```

### File: internal/inspector/inspector.go (esqueleto)

```go
package inspector

// Aqui ficará a lógica de amostragem, merge de exemplos e deduplicação.
// Para a v0 inicial usamos a CaptureMiddleware que registra exemplos diretamente na spec.
```

------

## Fluxo interno resumido

- `autospec.New()` cria um `openapi.Spec` mínimo.
- `as.AttachTo(app)` detecta framework via adapters e chama `Attach` do adapter.
- Adapter registra `/openapi.json` e `/docs` e instala um `CaptureMiddleware` (coleta amostras sem alterar handlers).
- Middleware captura request/response e insere exemplos em `spec`.
- Opcional: `autospec.Handle(...)` registra rotas declarativamente e adiciona metadados precisos.

------

## Roadmap técnico curto

- v0.1: Core + Gin adapter + capture middleware + Swagger UI.
- v0.2: Schema builder reflect + sample merge + masking de campos sensíveis.
- v0.3: Plugin system, AuthDetector plugin, middleware detector.
- v0.4: Adapters: Echo, Fiber, Chi, net/http.

------

## Problemas técnicos e recomendações

- Inferência apenas via amostras é heurística; oferecer combinador declarativo (Meta) para precisão.
- Limitar tamanho das amostras por padrão; mascarar campos sensíveis.
- Fornecer modo "offline" (análise estática com go/packages) para projetos que preferem zero-runtime-overhead.

------

Se quiser, implemento agora esses arquivos no repositório do workspace (criar/atualizar arquivos), e então executo um build Go rápido para validar. Quer que eu faça isso agora?un(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
```

### File: internal/autospec/options.go

```go
package autospec

type Option func(*Autospec)

func WithPlugin(p Plugin) Option {
	return func(a *Autospec) { a.plugins = append(a.plugins, p) }
}

func WithAdapterName(name string) Option {
	return func(a *Autospec) { a.preferredAdapter = name }
}
```

### File: internal/autospec/plugin.go

```go
package autospec

import "reflect"

type Plugin interface {
	Name() string
	OnRoute(route RouteMeta)
	OnSample(sample RequestSample)
	SchemaFromType(t reflect.Type) (map[string]any, error)
}
```

### File: internal/autospec/autospec.go

```go
package autospec

import (
	"errors"

	ginadapter "github.com/Jeielsantosdev/autospec/internal/adapters/gin"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
)

type Autospec struct {
	Spec            *openapi.Spec
	plugins         []Plugin
	preferredAdapter string
}

func New(opts ...Option) *Autospec {
	a := &Autospec{}
	for _, o := range opts { o(a) }
	if a.Spec == nil {
		a.Spec = openapi.NewSpec("autospec", "dev", "autospec generated spec")
		a.Spec.AddServer("http://localhost:8080", "local")
	}
	return a
}

// AttachTo tenta detectar o framework e anexar automaticamente
func (a *Autospec) AttachTo(app any) error {
	// Gin
	if ginadapter.Detect(app) {
		return ginadapter.Attach(app, a.Spec)
	}

	return errors.New("no adapter found for given app instance")
}

// Handle: opção declarativa para registrar rota + metadados
func (a *Autospec) Handle(app any, method, path string, handler any, meta Meta) error {
	// Para simplicidade, delegamos a adapters quando necessário.
	if ginadapter.Detect(app) {
		return ginadapter.RegisterRouteWithMeta(app, method, path, handler, a.Spec, meta)
	}
	return errors.New("handle: unsupported adapter")
}

// Run minimal CLI surface
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

// Meta (declarative) e tipos de amostra
type Meta struct {
	Input any
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

// Helpers para uso em handlers (pequeno wrapper)
func BindJSON(ctx any, dest any) error { // implementado por adapters via type switch
	// placeholder: adapters podem prover BindJSON mais preciso
	return nil
}

func RespondJSON(ctx any, status int, body any) error {
	return nil
}
```

### File: internal/adapters/gin/attach.go

```go
package ginadapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
	"github.com/Jeielsantosdev/autospec/internal/generators"
)

// Detect tenta identificar se `app` é um *gin.Engine
func Detect(app any) bool {
	_, ok := app.(*gin.Engine)
	return ok
}

// Attach registra rotas de docs e instala middleware de captura
func Attach(app any, spec *openapi.Spec) error {
	engine, ok := app.(*gin.Engine)
	if !ok {
		return nil
	}

	// registrar docs
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

	// instalar middleware de captura (coleta amostras para inferência)
	engine.Use(CaptureMiddleware(spec))

	return nil
}

// RegisterRouteWithMeta permite registro declarativo
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
```

### File: internal/adapters/gin/capture_middleware.go

```go
package ginadapter

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// CaptureMiddleware captura request e response e adiciona exemplo simples à spec
func CaptureMiddleware(spec *openapi.Spec) gin.HandlerFunc {
	return func(c *gin.Context) {
		// capturar request
		var reqBuf []byte
		if c.Request.Body != nil {
			b, _ := io.ReadAll(c.Request.Body)
			reqBuf = b
			c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
		}

		// capturar response
		bw := &bodyWriter{ResponseWriter: c.Writer, body: bytes.NewBuffer(nil)}
		c.Writer = bw

		c.Next()

		status := c.Writer.Status()
		path := c.FullPath()
		method := c.Request.Method

		// anexar exemplo simples na spec: usamos example sem schema real
		// para uma evolução futura, integrar reflect/schema_builder
		if path != "" {
			example := map[string]any{"request": string(reqBuf), "response": bw.body.String()}
			schema := openapi.Schema{Type: "object", Example: example}
			spec.AddJSONResponse(path, method, http.StatusText(status), "captured example", schema)
		}
	}
}
```

### File: internal/reflect/schema_builder.go (esqueleto)

```go
package reflect

// Implementação futura: reflect.Type -> JSON Schema
// Forneça utilitários para transformar structs em openapi.Schema
```

### File: internal/inspector/inspector.go (esqueleto)

```go
package inspector

// Aqui ficará a lógica de amostragem, merge de exemplos e deduplicação.
// Para a v0 inicial usamos a CaptureMiddleware que registra exemplos diretamente na spec.
```

------

## Fluxo interno resumido

- `autospec.New()` cria um `openapi.Spec` mínimo.
- `as.AttachTo(app)` detecta framework via adapters e chama `Attach` do adapter.
- Adapter registra `/openapi.json` e `/docs` e instala um `CaptureMiddleware` (coleta amostras sem alterar handlers).
- Middleware captura request/response e insere exemplos em `spec`.
- Opcional: `autospec.Handle(...)` registra rotas declarativamente e adiciona metadados precisos.

------

## Roadmap técnico curto

- v0.1: Core + Gin adapter + capture middleware + Swagger UI.
- v0.2: Schema builder reflect + sample merge + masking de campos sensíveis.
- v0.3: Plugin system, AuthDetector plugin, middleware detector.
- v0.4: Adapters: Echo, Fiber, Chi, net/http.

------

## Problemas técnicos e recomendações

- Inferência apenas via amostras é heurística; oferecer combinador declarativo (Meta) para precisão.
- Limitar tamanho das amostras por padrão; mascarar campos sensíveis.
- Fornecer modo "offline" (análise estática com go/packages) para projetos que preferem zero-runtime-overhead.

------

Se quiser, implemento agora esses arquivos no repositório do workspace (criar/atualizar arquivos), e então executo um build Go rápido para validar. Quer que eu faça isso agora?

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

### File: internal/autospec/options.go

```go
package autospec

type Option func(*Autospec)

func WithPlugin(p Plugin) Option {
	return func(a *Autospec) { a.plugins = append(a.plugins, p) }
}

func WithAdapterName(name string) Option {
	return func(a *Autospec) { a.preferredAdapter = name }
}
```

### File: internal/autospec/plugin.go

```go
package autospec

import "reflect"

type Plugin interface {
	Name() string
	OnRoute(route RouteMeta)
	OnSample(sample RequestSample)
	SchemaFromType(t reflect.Type) (map[string]any, error)
}
```

### File: internal/autospec/autospec.go

```go
package autospec

import (
	"errors"

	ginadapter "github.com/Jeielsantosdev/autospec/internal/adapters/gin"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
)

type Autospec struct {
	Spec            *openapi.Spec
	plugins         []Plugin
	preferredAdapter string
}

func New(opts ...Option) *Autospec {
	a := &Autospec{}
	for _, o := range opts { o(a) }
	if a.Spec == nil {
		a.Spec = openapi.NewSpec("autospec", "dev", "autospec generated spec")
		a.Spec.AddServer("http://localhost:8080", "local")
	}
	return a
}

// AttachTo tenta detectar o framework e anexar automaticamente
func (a *Autospec) AttachTo(app any) error {
	// Gin
	if ginadapter.Detect(app) {
		return ginadapter.Attach(app, a.Spec)
	}

	return errors.New("no adapter found for given app instance")
}

// Handle: opção declarativa para registrar rota + metadados
func (a *Autospec) Handle(app any, method, path string, handler any, meta Meta) error {
	// Para simplicidade, delegamos a adapters quando necessário.
	if ginadapter.Detect(app) {
		return ginadapter.RegisterRouteWithMeta(app, method, path, handler, a.Spec, meta)
	}
	return errors.New("handle: unsupported adapter")
}

// Run minimal CLI surface
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

// Meta (declarative) e tipos de amostra
type Meta struct {
	Input any
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

// Helpers para uso em handlers (pequeno wrapper)
func BindJSON(ctx any, dest any) error { // implementado por adapters via type switch
	// placeholder: adapters podem prover BindJSON mais preciso
	return nil
}

func RespondJSON(ctx any, status int, body any) error {
	return nil
}
```

### File: internal/adapters/gin/attach.go

```go
package ginadapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
	"github.com/Jeielsantosdev/autospec/internal/generators"
)

// Detect tenta identificar se `app` é um *gin.Engine
func Detect(app any) bool {
	_, ok := app.(*gin.Engine)
	return ok
}

// Attach registra rotas de docs e instala middleware de captura
func Attach(app any, spec *openapi.Spec) error {
	engine, ok := app.(*gin.Engine)
	if !ok {
		return nil
	}

	// registrar docs
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

	// instalar middleware de captura (coleta amostras para inferência)
	engine.Use(CaptureMiddleware(spec))

	return nil
}

// RegisterRouteWithMeta permite registro declarativo
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
```

### File: internal/adapters/gin/capture_middleware.go

```go
package ginadapter

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// CaptureMiddleware captura request e response e adiciona exemplo simples à spec
func CaptureMiddleware(spec *openapi.Spec) gin.HandlerFunc {
	return func(c *gin.Context) {
		// capturar request
		var reqBuf []byte
		if c.Request.Body != nil {
			b, _ := io.ReadAll(c.Request.Body)
			reqBuf = b
			c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
		}

		// capturar response
		bw := &bodyWriter{ResponseWriter: c.Writer, body: bytes.NewBuffer(nil)}
		c.Writer = bw

		c.Next()

		status := c.Writer.Status()
		path := c.FullPath()
		method := c.Request.Method

		// anexar exemplo simples na spec: usamos example sem schema real
		// para uma evolução futura, integrar reflect/schema_builder
		if path != "" {
			example := map[string]any{"request": string(reqBuf), "response": bw.body.String()}
			schema := openapi.Schema{Type: "object", Example: example}
			spec.AddJSONResponse(path, method, http.StatusText(status), "captured example", schema)
		}
	}
}
```

### File: internal/reflect/schema_builder.go (esqueleto)

```go
package reflect

// Implementação futura: reflect.Type -> JSON Schema
// Forneça utilitários para transformar structs em openapi.Schema
```

### File: internal/inspector/inspector.go (esqueleto)

```go
package inspector

// Aqui ficará a lógica de amostragem, merge de exemplos e deduplicação.
// Para a v0 inicial usamos a CaptureMiddleware que registra exemplos diretamente na spec.
```

------

## Fluxo interno resumido

- `autospec.New()` cria um `openapi.Spec` mínimo.
- `as.AttachTo(app)` detecta framework via adapters e chama `Attach` do adapter.
- Adapter registra `/openapi.json` e `/docs` e instala um `CaptureMiddleware` (coleta amostras sem alterar handlers).
- Middleware captura request/response e insere exemplos em `spec`.
- Opcional: `autospec.Handle(...)` registra rotas declarativamente e adiciona metadados precisos.

------

## Roadmap técnico curto

- v0.1: Core + Gin adapter + capture middleware + Swagger UI.
- v0.2: Schema builder reflect + sample merge + masking de campos sensíveis.
- v0.3: Plugin system, AuthDetector plugin, middleware detector.
- v0.4: Adapters: Echo, Fiber, Chi, net/http.

------

## Problemas técnicos e recomendações

- Inferência apenas via amostras é heurística; oferecer combinador declarativo (Meta) para precisão.
- Limitar tamanho das amostras por padrão; mascarar campos sensíveis.
- Fornecer modo "offline" (análise estática com go/packages) para projetos que preferem zero-runtime-overhead.

------

Se quiser, implemento agora esses arquivos no repositório do workspace (criar/atualizar arquivos), e então executo um build Go rápido para validar. Quer que eu faça isso agora?