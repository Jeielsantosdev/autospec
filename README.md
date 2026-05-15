# SDK Go — Autospec

Projeto em Go para geração/inspeção de especificações OpenAPI e utilitários relacionados.

**Requisitos**

- Go instalado (1.18+)

**Instalação**

1. Clone o repositório.
2. No diretório do projeto, baixe dependências (Go modules):

```bash
go mod download
```

**Uso**

- Compilar tudo:

```bash
go build ./...
```

- Rodar o binário principal (exemplo de desenvolvimento):

```bash
go run ./cmd/autospec
```

- Executar testes:

```bash
go test ./...
```

**Estrutura principal**

- `cmd/autospec` — ponto de entrada da aplicação
- `internal/adapters/gin` — adaptador HTTP e rotas
- `internal/autospec` — lógica da aplicação
- `internal/openapi` — geradores e utilitários OpenAPI
- `internal/generators` — geradores auxiliares (ex.: Swagger UI)

**Notas rápidas**

- Ajuste variáveis de ambiente criando um `.env` na raiz se necessário.
- O projeto já contém geradores e comandos de desenvolvimento em `internal/commands`.

Se quiser, eu posso:

- Adicionar instruções de configuração mais detalhadas;
- Incluir exemplos de uso das APIs;
- Publicar no Git remoto (`git push`).

**Exemplo de Uso com Gin Gonic**

Um exemplo mínimo que inicializa um `gin` server, registra rotas, gera a especificação OpenAPI em runtime usando o adaptador do projeto e expõe as rotas de documentação (`/openapi.json` e `/docs`):

```go
package main

import (
	"net/http"

	ginadapter "github.com/Jeielsantosdev/autospec/internal/adapters/gin"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// Cria a Spec OpenAPI
	spec := openapi.NewSpec("Minha API", "v0.1.0", "Exemplo com Gin Gonic e autospec")
	spec.AddServer("http://localhost:8080", "Servidor local")

	// Registre suas rotas normalmente
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Olá, mundo"})
	})

	// Anexa o adaptador que varre as rotas em runtime e popula a Spec
	ginadapter.New(engine, spec).Attach()

	// Rotas para expor a documentação gerada
	ginadapter.RegisterDocsRoutes(engine, spec)

	// Inicia o servidor
	engine.Run(":8080")
}
```

Comandos úteis:

```bash
# Baixar dependências
go mod download

# Rodar o exemplo (a partir da raiz do repositório)
go run ./cmd/autospec  # ou crie um main a partir do snippet e execute: go run ./path/to/main.go
```

Esse exemplo usa as funções internas do projeto (`internal/adapters/gin` e `internal/openapi`). Ajuste os imports/paths conforme necessário se for extrair a parte de geração para outro módulo.
