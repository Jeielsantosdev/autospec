package openapi

import (
	"encoding/json"
	"testing"
)

func TestNormalizePath(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "single param", in: "/users/:id", want: "/users/{id}"},
		{name: "multiple params", in: "/teams/:teamId/users/:userId", want: "/teams/{teamId}/users/{userId}"},
		{name: "already normalized", in: "/users/{id}", want: "/users/{id}"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := normalizePath(tt.in); got != tt.want {
				t.Fatalf("normalizePath(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestExtractPathParameters(t *testing.T) {
	t.Parallel()

	parameters := extractPathParameters("/teams/{teamId}/users/{userId}/members/{userId}")
	if len(parameters) != 2 {
		t.Fatalf("expected 2 unique parameters, got %d: %+v", len(parameters), parameters)
	}

	if parameters[0].Name != "teamId" || parameters[0].In != "path" || !parameters[0].Required || parameters[0].Schema.Type != "string" {
		t.Fatalf("unexpected first parameter: %+v", parameters[0])
	}
	if parameters[1].Name != "userId" || parameters[1].In != "path" || !parameters[1].Required || parameters[1].Schema.Type != "string" {
		t.Fatalf("unexpected second parameter: %+v", parameters[1])
	}
}

func TestFormatOperationName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		handler string
		want    string
	}{
		{name: "package qualified", handler: "main.CreateUser", want: "Create User"},
		{name: "module path", handler: "github.com/acme/app/internal/users.(*Handler).ListUsers", want: "List Users"},
		{name: "anonymous function", handler: "main.main.func1", want: "Main"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := formatOperationName(tt.handler); got != tt.want {
				t.Fatalf("formatOperationName(%q) = %q, want %q", tt.handler, got, tt.want)
			}
		})
	}
}

func TestAddOperationAddsNormalizedPathAndParameters(t *testing.T) {
	t.Parallel()

	spec := NewSpec("test", "1.0.0", "")
	spec.AddOperation("/users/:id", "GET", Operation{Summary: "main.CreateUser"})

	item, ok := spec.Paths["/users/{id}"]
	if !ok {
		t.Fatalf("normalized path not found in spec paths: %+v", spec.Paths)
	}

	if item.Get == nil {
		t.Fatalf("expected GET operation on normalized path")
	}
	if item.Get.Summary != "Create User" {
		t.Fatalf("unexpected summary: %q", item.Get.Summary)
	}
	if len(item.Get.Parameters) != 1 {
		t.Fatalf("expected one parameter, got %+v", item.Get.Parameters)
	}
	param := item.Get.Parameters[0]
	if param.Name != "id" || param.In != "path" || !param.Required || param.Schema.Type != "string" {
		t.Fatalf("unexpected parameter: %+v", param)
	}
}

func TestAddJSONResponseAddsPathParameters(t *testing.T) {
	t.Parallel()

	spec := NewSpec("test", "1.0.0", "")
	spec.AddJSONResponse("/users/:id", "GET", "200", "OK", Schema{Type: "object"})

	item, ok := spec.Paths["/users/{id}"]
	if !ok {
		t.Fatalf("normalized path not found in spec paths: %+v", spec.Paths)
	}

	if item.Get == nil {
		t.Fatalf("expected GET operation on normalized path")
	}
	if len(item.Get.Parameters) != 1 {
		t.Fatalf("expected one path parameter, got %+v", item.Get.Parameters)
	}
	param := item.Get.Parameters[0]
	if param.Name != "id" || param.In != "path" || !param.Required || param.Schema.Type != "string" {
		t.Fatalf("unexpected parameter: %+v", param)
	}
}

func TestAddJSONRequestBodyAddsPayload(t *testing.T) {
	t.Parallel()

	spec := NewSpec("test", "1.0.0", "")
	spec.AddJSONRequestBody("/users", "POST", true, Schema{Type: "object", Example: map[string]any{"name": "john"}})

	item, ok := spec.Paths["/users"]
	if !ok {
		t.Fatalf("path not found in spec paths: %+v", spec.Paths)
	}

	if item.Post == nil {
		t.Fatalf("expected POST operation on /users")
	}
	if item.Post.RequestBody == nil {
		t.Fatalf("expected request body on POST /users")
	}
	if !item.Post.RequestBody.Required {
		t.Fatalf("expected request body to be required")
	}
	content := item.Post.RequestBody.Content["application/json"]
	if content.Schema.Type != "object" {
		t.Fatalf("unexpected request body schema: %+v", content.Schema)
	}
}

func TestInferJSONSchemaFromBytes(t *testing.T) {
	t.Parallel()

	schema := InferJSONSchemaFromBytes([]byte(`{"name":"john","age":30}`))
	if schema.Type != "object" {
		t.Fatalf("expected object schema, got %q", schema.Type)
	}

	encoded, err := json.Marshal(schema.Example)
	if err != nil {
		t.Fatalf("failed to marshal example: %v", err)
	}
	if string(encoded) != `{"age":30,"name":"john"}` && string(encoded) != `{"name":"john","age":30}` {
		t.Fatalf("unexpected example payload: %s", string(encoded))
	}
}
