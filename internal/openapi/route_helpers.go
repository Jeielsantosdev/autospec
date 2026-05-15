package openapi

import (
	"path"
	"regexp"
	"strings"
	"unicode"
)

var (
	ginPathParamPattern     = regexp.MustCompile(`:([^/]+)`)
	openAPIPathParamPattern = regexp.MustCompile(`\{([^/{}]+)\}`)
	camelWordBoundary       = regexp.MustCompile(`([a-z0-9])([A-Z])`)
	acronymBoundary         = regexp.MustCompile(`([A-Z]+)([A-Z][a-z])`)
)

func normalizePath(routePath string) string {
	return ginPathParamPattern.ReplaceAllString(routePath, `{$1}`)
}

func extractPathParameters(routePath string) []Parameter {
	matches := openAPIPathParamPattern.FindAllStringSubmatch(routePath, -1)
	if len(matches) == 0 {
		return nil
	}

	parameters := make([]Parameter, 0, len(matches))
	seen := make(map[string]struct{}, len(matches))

	for _, match := range matches {
		if len(match) < 2 {
			continue
		}

		name := strings.TrimSpace(match[1])
		if name == "" {
			continue
		}

		if _, exists := seen[name]; exists {
			continue
		}

		seen[name] = struct{}{}
		parameters = append(parameters, Parameter{
			Name:     name,
			In:       "path",
			Required: true,
			Schema:   Schema{Type: "string"},
		})
	}

	return parameters
}

func mergePathParameters(existing []Parameter, generated []Parameter) []Parameter {
	if len(existing) == 0 {
		return generated
	}

	merged := make([]Parameter, len(existing))
	copy(merged, existing)
	index := make(map[string]int, len(existing))

	for i, parameter := range merged {
		index[parameterKey(parameter)] = i
	}

	for _, parameter := range generated {
		key := parameterKey(parameter)
		if position, exists := index[key]; exists {
			merged[position].In = "path"
			merged[position].Required = true
			if merged[position].Schema.Type == "" {
				merged[position].Schema.Type = "string"
			}
			continue
		}

		merged = append(merged, parameter)
		index[key] = len(merged) - 1
	}

	return merged
}

func parameterKey(parameter Parameter) string {
	return parameter.In + ":" + parameter.Name
}

func looksLikeHandlerName(handler string) bool {
	return strings.Contains(handler, ".") || strings.Contains(handler, "/") || strings.Contains(handler, "func")
}

func formatOperationName(handler string) string {
	handler = strings.TrimSpace(handler)
	if handler == "" {
		return ""
	}

	handler = strings.TrimSuffix(handler, "-fm")
	handler = path.Base(handler)

	segments := strings.Split(handler, ".")
	candidate := ""
	for i := len(segments) - 1; i >= 0; i-- {
		segment := strings.TrimSpace(segments[i])
		if segment == "" || isAnonymousFunction(segment) {
			continue
		}

		candidate = segment
		break
	}

	if candidate == "" {
		candidate = segments[len(segments)-1]
	}

	candidate = strings.Trim(candidate, "()")
	candidate = strings.TrimPrefix(candidate, "*")
	candidate = strings.ReplaceAll(candidate, "_", " ")
	candidate = acronymBoundary.ReplaceAllString(candidate, "$1 $2")
	candidate = camelWordBoundary.ReplaceAllString(candidate, "$1 $2")
	words := strings.Fields(candidate)
	for i, word := range words {
		words[i] = titleWord(word)
	}

	return strings.Join(words, " ")
}

func isAnonymousFunction(segment string) bool {
	if !strings.HasPrefix(segment, "func") {
		return false
	}

	for _, r := range segment[4:] {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return len(segment) > 4
}

func titleWord(word string) string {
	if word == "" {
		return word
	}

	runes := []rune(word)
	if unicode.IsUpper(runes[0]) {
		return word
	}

	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
