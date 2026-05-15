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
