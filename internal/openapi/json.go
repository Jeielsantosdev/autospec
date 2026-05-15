package openapi

import "encoding/json"
import "fmt"

func (s *Spec) JSON() ([]byte, error) {
	// debug: print paths before serialization
	fmt.Printf("[autospec] serializing spec.Paths: %+v\n", s.Paths)
	return json.MarshalIndent(s, "", " ")
}

func (s *Spec) YAMLPlaceholder() string {
	return "YAML generation will be added in a later step"
}
