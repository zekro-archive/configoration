package configoration

import (
	"path"
	"sync"

	"github.com/zekroTJA/configoration/providers"
)

// Builder provides functions to build a config
// with different source providers.
//
// Every function returns the Builder instance
// to be able to apply the builder pattern.
type Builder struct {
	provider []Provider

	basePath string
}

// NewBuilder returns a new instance of builder.
func NewBuilder() *Builder {
	return &Builder{
		provider: make([]Provider, 0),
	}
}

// SetBasePath sets the base path from which
// file providers are reading when a relative
// path is given.
func (b *Builder) SetBasePath(path string) *Builder {
	b.basePath = path
	return b
}

// AddJsonFile adds a JSON file provider which
// reads the passed fileName respecting the set
// base path. If optional is set, no error is
// returned when the file does not exist.
func (b *Builder) AddJsonFile(fileName string, optional bool) *Builder {
	p := providers.NewJsonProvider(path.Join(b.basePath, fileName), optional)
	return b.AddProvider(p)
}

// AddYamlFile adds a YAML file provider which
// reads the passed fileName respecting the set
// base path. If optional is set, no error is
// returned when the file does not exist.
func (b *Builder) AddYamlFile(fileName string, optional bool) *Builder {
	p := providers.NewYamlProvider(path.Join(b.basePath, fileName), optional)
	return b.AddProvider(p)
}

// AddProvider adds a generic Provider instance
// which must implememt the Provider interface.
func (b *Builder) AddProvider(p Provider) *Builder {
	b.provider = append(b.provider, p)
	return b
}

// Build esecutes all registered providers in
// the given order and builds the resulting
// config, which is returned.
//
// If a registered provider fails, the build
// stops and returns the error. The resulting
// Section will be nil.
func (b *Builder) Build() (Section, error) {
	res := make(ConfigMap)
	for _, prov := range b.provider {
		m, err := prov.GetMap()
		if err != nil {
			return nil, err
		}
		res.merge(m)
	}

	return &section{
		mtx: sync.Mutex{},
		m:   res,
	}, nil
}
