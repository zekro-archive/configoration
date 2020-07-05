package providers

import (
	"os"
	"strings"
)

const (
	envDelimiter = "__"
)

// EnvProvider implements the Provider interface for
// environment variables as configuration providers.
type EnvProvider struct {
	prefix    string
	lowercase bool
}

// NewEnvProvider returns a new instance of EnvProvider
// with the passed prefix and lowercase specification.
func NewEnvProvider(prefix string, lowercase bool) *EnvProvider {
	return &EnvProvider{
		prefix:    prefix,
		lowercase: lowercase,
	}
}

func (p *EnvProvider) GetMap() (map[string]interface{}, error) {
	environ := os.Environ()
	env := make(map[string]interface{})

	for _, e := range environ {
		if !strings.HasPrefix(e, p.prefix) {
			continue
		}

		e = e[len(p.prefix):]

		kvSplit := strings.SplitN(e, "=", 2)
		key := kvSplit[0]
		val := kvSplit[1]

		if p.lowercase {
			key = strings.ToLower(key)
		}

		sections := strings.Split(key, envDelimiter)
		if len(sections) > 1 {
			ensurePathAndSetValue(env, sections, val)
		} else {
			env[key] = val
		}
	}

	return env, nil
}

func ensurePathAndSetValue(m map[string]interface{}, sections []string, val interface{}) {
	for i := 0; i < len(sections)-1; i++ {
		sec := sections[i]
		if _, ok := m[sec]; !ok {
			m[sec] = make(map[string]interface{})
		}
		m = m[sec].(map[string]interface{})
	}

	m[sections[len(sections)-1]] = val
}
