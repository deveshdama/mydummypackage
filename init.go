package mydummypackage

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

//go:embed data/data.yaml
var data []byte

type Data struct {
	Foo string
	Bar string
	Baz string
}

func LoadData() (*Data, error) {
	var d Data
	if err := yaml.Unmarshal(data, &d); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %v", err)
	}
	return &d, nil
}
