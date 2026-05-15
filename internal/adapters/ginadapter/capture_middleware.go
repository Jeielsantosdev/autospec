package ginadapter

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/Jeielsantosdev/autospec/internal/openapi"
	"github.com/gin-gonic/gin"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	if w.body == nil {
		w.body = bytes.NewBuffer(nil)
	}
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func CaptureMiddleware(spec *openapi.Spec) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqbuf []byte
		if c.Request.Body != nil {
			b, _ := io.ReadAll(c.Request.Body)
			reqbuf = b
			c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
		}

		bw := &bodyWriter{ResponseWriter: c.Writer, body: bytes.NewBuffer(nil)}
		c.Writer = bw

		c.Next()

		status := c.Writer.Status()
		path := c.FullPath()
		method := c.Request.Method
		handlerName := c.HandlerName()

		if path != "" {
			spec.AddOperation(path, method, openapi.Operation{
				Summary:     handlerName,
				Description: "Route discovered at runtime",
				Tags:        []string{},
			})
			contentType := strings.ToLower(c.GetHeader("Content-Type"))
			if len(reqbuf) > 0 && strings.Contains(contentType, "application/json") {
				schema := openapi.InferJSONSchemaFromBytes(reqbuf)
				spec.AddJSONRequestBody(path, method, true, schema)
			}
			example := map[string]any{"request": string(reqbuf), "response": bw.body.String()}
			schema := openapi.Schema{Type: "object", Example: example}
			spec.AddJSONResponse(path, method, http.StatusText(status),
				"captured example", schema)
		}
	}

}
