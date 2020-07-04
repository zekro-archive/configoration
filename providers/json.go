package providers

import (
	"encoding/json"
	"os"
)

// JsonProvider implements the Provider interface
// for reading JSON config files.
type JsonProvider struct {
	fileName string
	optional bool
}

// NewJsonProvider produces a new YamlProvider instance
// with the given fileName and optional flag.
func NewJsonProvider(fileName string, optional bool) *JsonProvider {
	return &JsonProvider{
		fileName: fileName,
		optional: optional,
	}
}

func (p *JsonProvider) GetMap() (map[string]interface{}, error) {
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
	dec := json.NewDecoder(f)
	err = dec.Decode(&m)

	return m, err
}
