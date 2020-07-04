package providers

import (
	"os"

	"gopkg.in/yaml.v2"
)

// YamlProvider implements the Provider interface
// for reading YAML config files.
type YamlProvider struct {
	fileName string
	optional bool
}

// NewYamlProvider produces a new YamlProvider instance
// with the given fileName and optional flag.
func NewYamlProvider(fileName string, optional bool) *YamlProvider {
	return &YamlProvider{
		fileName: fileName,
		optional: optional,
	}
}

func (p *YamlProvider) GetMap() (map[string]interface{}, error) {
	_, err := os.Stat(p.fileName)
	if err != nil {
		if os.IsNotExist(err) && p.optional {
			return nil, nil
		}
		return nil, err
	}

	f, err := os.Open(p.fileName)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	dec := yaml.NewDecoder(f)
	err = dec.Decode(&m)

	return m, err
}
