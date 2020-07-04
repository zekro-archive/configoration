package providers

import (
	"fmt"
	"os"
	"strings"
)

// TODO: Docs

const (
	envDelimiter = "__"
)

type EnvProvider struct {
	prefix    string
	lowercase bool
}

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
		fmt.Println(kvSplit)
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

	fmt.Printf("%+v\n", env)
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
