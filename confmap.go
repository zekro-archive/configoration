package configoration

import "fmt"

// ConfigMap extends map[string]interface{} with
// functionalities to merge two of them together.
type ConfigMap map[string]interface{}

// merge combines confMap with m by merging.
//
// Existing keys are owerwritten and non-
// existing keys are added to m.
func (m ConfigMap) merge(confMap ConfigMap) {
	for k, v := range confMap {
		if vm, ok := v.(map[interface{}]interface{}); ok {
			nm := make(map[string]interface{})
			for k, v := range vm {
				nm[fmt.Sprintf("%v", k)] = v
			}
			m.mergeInnerMap(ConfigMap(nm), k)
		} else if vm, ok := v.(map[string]interface{}); ok {
			m.mergeInnerMap(ConfigMap(vm), k)
		} else {
			m[k] = v
		}
	}
}

// mergeInnerMap merges confMap with an
// inner map of m with the specified
// innerKey.
//
// If the value of innerKey is not a
// ConfigMap, then the function returns.
func (m ConfigMap) mergeInnerMap(confMap ConfigMap, innerKey string) {
	if _, ok := m[innerKey]; !ok {
		m[innerKey] = make(ConfigMap)
	}

	_, ok := m[innerKey].(ConfigMap)
	if !ok {
		return
	}

	m[innerKey].(ConfigMap).merge(confMap)
}
