package openapi

import "encoding/json"

func (s *Spec) JSON() ([]byte, error) {
	return json.MarshalIndent(s, "", " ")
}

func (s *Spec) YAMLPlaceholder() string {
	return "YAML generation will be added in a later step"
}
