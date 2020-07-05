package configoration

import (
	"testing"
)

func TestMerge(t *testing.T) {
	cm := make(ConfigMap)

	cm1 := ConfigMap{
		"a": 1,
		"b": map[string]interface{}{
			"b1": 1,
			"b2": 1,
		},
		"c": 1,
		"d": 1,
	}

	cm2 := ConfigMap{
		"a": 2,
		"b": map[string]interface{}{
			"b1": 2,
			"b3": 2,
		},
		"d": map[interface{}]interface{}{
			"d1": 2,
		},
	}

	cm.merge(cm1)
	cm.merge(cm2)

	assert(t, cm["a"], 2)
	assert(t, cm["b"].(ConfigMap)["b1"], 2)
	assert(t, cm["b"].(ConfigMap)["b2"], 1)
	assert(t, cm["b"].(ConfigMap)["b3"], 2)
	assert(t, cm["c"], 1)
	assert(t, cm["d"].(ConfigMap)["d1"], 2)
}

func TestMerheInnerMap(t *testing.T) {
	cm := make(ConfigMap)

	cm.merge(ConfigMap{
		"a": map[string]interface{}{
			"a1": 1,
			"a2": 1,
		},
	})

	innerCm := ConfigMap{
		"a1": 2,
		"a3": 2,
	}

	cm.mergeInnerMap(innerCm, "a")

	assert(t, cm["a"].(ConfigMap)["a1"], 2)
	assert(t, cm["a"].(ConfigMap)["a2"], 1)
	assert(t, cm["a"].(ConfigMap)["a3"], 2)
}

// --------------------------------------------------------------------------
// --- HELPERS

func assert(t *testing.T, val, expected interface{}) {
	if val != expected {
		t.Errorf("value (%+v) was not like expected (%+v)", val, expected)
	}
}
