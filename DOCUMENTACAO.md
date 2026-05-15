# Documentação do Projeto autospec

## Visão Geral

O `autospec` será um SDK para Go voltado à geração automática de documentação OpenAPI/Swagger, com foco em uma experiência semelhante à do FastAPI.

A proposta central é simples:

- sem annotations
- sem comentários Swagger
- sem boilerplate
- sem duplicação de documentação

O código da aplicação será a única fonte de verdade.

## Problema que queremos resolver

Hoje, documentar APIs em Go ainda depende muito de comentários, tags e annotations manuais. Isso cria:

- documentação duplicada
- alto custo de manutenção
- risco de Swagger desatualizado
- perda de produtividade
- baixa consistência entre código e documentação

## Objetivo do Projeto

Construir uma solução que consiga observar a aplicação Go e gerar documentação de API automaticamente, com pouco ou nenhum esforço extra do desenvolvedor.

O resultado esperado é algo como:

```go
app.POST("/users", CreateUser)
```

E, a partir disso, gerar automaticamente:

- OpenAPI
- Swagger UI
- request schemas
- response schemas
- validações
- exemplos
- metadados de rota
- suporte a autenticação

## Princípios do autospec

### 1. Code First

O código da aplicação define o contrato. Não haverá necessidade de escrever documentação paralela.

### 2. Zero Config

O uso deve ser simples e sem configuração obrigatória.

### 3. Introspecção Automática

O SDK deve analisar código, tipos e execução real da aplicação para inferir o máximo de informação possível.

### 4. Plug and Play

A solução deve funcionar sobre frameworks existentes, sem exigir migração total de stack.

## Como o autospec deve funcionar

O projeto vai combinar três estratégias principais:

- AST Parsing para analisar handlers, structs, rotas, requests e responses
- Reflection para inferir tipos e estruturas em tempo de execução
- Runtime Inspection para identificar middlewares, grupos de rotas, autenticação e metadata real

## Estrutura proposta

```txt
autospec/
├── core/
├── adapters/
├── runtime/
├── generators/
├── plugins/
└── cmd/
```

### core

Responsável pelo núcleo da inteligência do SDK:

- parsing de AST
- reflection
- schemas
- geração OpenAPI

### adapters

Integração com frameworks web como:

- Gin
- Fiber
- Echo
- Chi
- net/http

### runtime

Camada responsável por observar a aplicação em execução e detectar:

- middlewares
- autenticação
- grupos de rotas
- metadata dinâmica
- rotas reais registradas

### generators

Geração dos artefatos finais:

- OpenAPI JSON/YAML
- Swagger UI
- exemplos
- schemas

### plugins

Sistema extensível para suportar:

- validação
- JWT
- OAuth
- exemplos automáticos
- observabilidade

### cmd

Ferramentas de linha de comando, como:

- `autospec dev`
- `autospec watch`

## Escopo Inicial

A primeira versão deve focar em:

- suporte ao Gin
- detecção de rotas
- detecção de request e response
- geração básica de OpenAPI
- integração inicial com Swagger UI

## Roadmap

### Fase 1

- suporte ao Gin
- request/response detection
- OpenAPI básico

### Fase 2

- reflection avançado
- inferência de tipos melhorada

### Fase 3

- runtime inspection
- auth detection
- middleware detection

### Fase 4

- hot reload
- docs live

### Fase 5

- VSCode extension
- geração de SDKs
- API diff
- integração com CI/CD

## Entregáveis Esperados

Ao final da implementação inicial, esperamos ter:

- uma base funcional do SDK
- integração com pelo menos um framework Go
- geração automática de spec OpenAPI
- documentação navegável via Swagger UI
- arquitetura preparada para expansão

## Diferencial Esperado

O autospec não deve ser apenas mais uma ferramenta de documentação. A proposta é entregar uma experiência moderna de DX, com:

- automação real
- compatibilidade com projetos existentes
- baixo acoplamento
- evolução para plataforma de DevTools no futuro

## Próximos Passos

1. Definir o MVP técnico da Fase 1.
2. Escolher a estrutura inicial do repositório.
3. Criar a primeira implementação para Gin.
4. Validar a geração básica de OpenAPI.
5. Evoluir para Swagger UI e hot reload.

## Resumo

O que iremos fazer é construir o `autospec` como um SDK inteligente para Go que lê a aplicação, entende suas rotas e tipos, e gera documentação OpenAPI automaticamente, sem annotations e sem boilerplate.
