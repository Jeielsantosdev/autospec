package openapi

import "strings"

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

func (s *Spec) AddSecurityScheme(name string, scheme SecurityScheme) {
	if s.Components.SecuritySchemes == nil {
		s.Components.SecuritySchemes = make(map[string]SecurityScheme)
	}

	s.Components.SecuritySchemes[name] = scheme
}

func (s *Spec) AddOperation(path string, method string, operation Operation) {
	normalizedPath := normalizePath(path)
	item := s.Paths[normalizedPath]
	operation = normalizeOperation(operation)
	operation.Parameters = mergePathParameters(operation.Parameters, extractPathParameters(normalizedPath))

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

	s.Paths[normalizedPath] = item
}

func (s *Spec) AddJSONResponse(path string, method string, status string, description string, schema Schema) {
	normalizedPath := normalizePath(path)
	item := s.Paths[normalizedPath]

	operation := operationForMethod(&item, method)
	if operation == nil {
		operation = &Operation{Responses: make(map[string]Response)}
	}

	if operation.Responses == nil {
		operation.Responses = make(map[string]Response)
	}

	operation.Parameters = mergePathParameters(operation.Parameters, extractPathParameters(normalizedPath))
	operation.Summary = normalizeSummary(operation.Summary)

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

	s.Paths[normalizedPath] = item
}

func (s *Spec) AddJSONRequestBody(path string, method string, required bool, schema Schema) {
	normalizedPath := normalizePath(path)
	item := s.Paths[normalizedPath]

	operation := operationForMethod(&item, method)
	if operation == nil {
		operation = &Operation{Responses: make(map[string]Response)}
	}

	operation.Parameters = mergePathParameters(operation.Parameters, extractPathParameters(normalizedPath))
	operation.Summary = normalizeSummary(operation.Summary)
	operation.RequestBody = &RequestBody{
		Required: required,
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

	s.Paths[normalizedPath] = item
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

func normalizeOperation(operation Operation) Operation {
	operation.Summary = normalizeSummary(operation.Summary)
	return operation
}

func normalizeSummary(summary string) string {
	summary = strings.TrimSpace(summary)
	if summary == "" || !looksLikeHandlerName(summary) {
		return summary
	}

	return formatOperationName(summary)
}
