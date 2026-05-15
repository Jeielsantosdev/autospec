# Autospec SDK (Go)

Autospec é um SDK para geração automática de OpenAPI em aplicações Go, com foco em excelente Developer Experience (DX). O objetivo é permitir "plug-and-play" em projetos existentes sem exigir anotações manuais.

Requisitos
- Go 1.18+

Instalação
1. Clone o repositório.
2. Baixe dependências:

```bash
go mod download
```

Como usar (nova API)

Exemplo mínimo — Zero-config (Gin):

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Jeielsantosdev/autospec/internal/autospec"
)

func main() {
	r := gin.Default()
	as := autospec.New()    // cria Spec e infra mínima
	_ = as.AttachTo(r)      // autodetecta Gin, registra /openapi.json e /docs e instala capture middleware

	r.POST("/users", CreateUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)
	r.GET("/users/:id", GetUser)

	r.Run(":8080")
}
```

Declarative (opcional):

```go
as.Handle(r, "POST", "/users", CreateUser, autospec.Meta{
	Input:  autospec.TypeOf[CreateUserRequest](),
	Output: autospec.TypeOf[User](),
})
```

Helpers (opcionais):

```go
func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	autospec.BindJSON(c, &req)          // registra tipo/amostra para melhor inferência
	user := doCreate(req)
	autospec.RespondJSON(c, 201, user)  // registra amostra de resposta
}
```

Rotas expostas automaticamente
- `/openapi.json` — JSON gerado da spec
- `/docs` — Swagger UI apontando para `/openapi.json`

Testes de integração

Um teste de integração de exemplo foi criado em `internal/examples/autospec_integration_test.go`. Para rodar somente esse teste:

```bash
go test ./internal/examples -run TestAutospecGeneratesOpenAPISpecForBasicRoutes
```

Comandos úteis

```bash
# Build
go build ./...

# Run tests
go test ./...
```

Estrutura principal

- `cmd/autospec` — CLI / ponto de entrada
- `internal/autospec` — core público (API: `New`, `AttachTo`, `Handle`)
- `internal/adapters/ginadapter` — adaptador Gin com middleware de captura
- `internal/openapi` — modelos e geração de spec
- `internal/generators` — Swagger UI page

Notas e recomendações
- A inferência via amostragem é heurística: oferecemos API declarativa (Meta) para precisão quando necessário.
- Por padrão, amostras são limitadas e campos sensíveis devem ser mascarados; implemente masking antes de usar em produção.
- Para ambientes sem overhead em runtime, utilize análise estática futura (go/packages) para gerar spec em build-time.

Próximos passos sugeridos
- Implementar `internal/reflect/schema_builder.go` para conversão `reflect.Type -> OpenAPI schema`.
- Melhorar o inspector para merge de amostras e detecção de autenticação.
- Adicionar adapters para Fiber, Echo, Chi e `net/http`.

Contribuições
- Abra PRs com pequenas mudanças; testes e exemplos são bem-vindos.

----

Se quiser, eu atualizo o README com comandos de CI ou adiciono um exemplo runnable em `examples/`.
