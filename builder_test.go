package configoration

import (
	"os"
	"testing"

	"github.com/zekroTJA/configoration/providers"
)

func TestNewBuilder(t *testing.T) {
	b := NewBuilder()
	if b == nil {
		t.Error("new builder instance was nil")
	}
	if b.provider == nil {
		t.Error("new builder provider array was nil")
	}
}

func TestSetBasePath(t *testing.T) {
	const testPath = "testpath"

	b := NewBuilder().
		SetBasePath(testPath)

	if b == nil {
		t.Error("returned builder instance was nil")
	}
	if b.basePath != testPath {
		t.Errorf("base path (%+v) was not set like expected (%+v)", b.basePath, testPath)
	}
}

func TestAddJsonFile(t *testing.T) {
	b := NewBuilder().
		AddJsonFile("file.json", false)

	if b == nil {
		t.Error("returned builder instance was nil")
	}
	if len(b.provider) != 1 {
		t.Error("providers array is empty")
	}
	v, ok := b.provider[0].(*providers.JsonProvider)
	if !ok || v == nil {
		t.Error("added provider is no JsonProvider")
	}
}

func TestAddYamlFile(t *testing.T) {
	b := NewBuilder().
		AddYamlFile("file.yaml", false)

	if b == nil {
		t.Error("returned builder instance was nil")
	}
	if len(b.provider) != 1 || b.provider[0] == nil {
		t.Error("providers array is empty")
	}
	v, ok := b.provider[0].(*providers.YamlProvider)
	if !ok || v == nil {
		t.Error("added provider is no YamlProvider")
	}
}

func TestAddEnvironmentVariables(t *testing.T) {
	b := NewBuilder().
		AddEnvironmentVariables("TEST_", false)

	if b == nil {
		t.Error("returned builder instance was nil")
	}
	if len(b.provider) != 1 || b.provider[0] == nil {
		t.Error("providers array is empty")
	}
	v, ok := b.provider[0].(*providers.EnvProvider)
	if !ok || v == nil {
		t.Error("added provider is no EnvProvider")
	}
}

func TestBuild(t *testing.T) {
	os.Setenv("TEST_b__f", "2")

	sec, err := NewBuilder().
		SetBasePath("./testdata").
		AddJsonFile("test1.json", false).
		AddJsonFile("test2.json", false).
		AddYamlFile("test3.yaml", false).
		AddEnvironmentVariables("TEST_", false).
		Build()

	if err != nil {
		t.Errorf("build failed: %s", err.Error())
	}

	{
		v, err := sec.GetString("a")
		assertVal(t, v, err, "test2")
	}
	{
		v, err := sec.GetInt("b:b")
		assertVal(t, v, err, 1)
	}
	{
		v, err := sec.GetInt("b:c")
		assertVal(t, v, err, 11112)
	}
	{
		v, err := sec.GetInt("b:e")
		assertVal(t, v, err, 3)
	}
	{
		v, err := sec.GetInt("b:f")
		assertVal(t, v, err, 2)
	}
	{
		v, err := sec.GetBool("g:e:f")
		assertVal(t, v, err, true)
	}
	{
		v, err := sec.GetInt("y:e")
		assertVal(t, v, err, 1123)
	}
}

// --------------------------------------------------------------------------
// --- HELPERS

func assertVal(t *testing.T, val interface{}, err error, expected interface{}) {
	if err != nil {
		t.Errorf("get value errored: %s", err.Error())
	} else if val != expected {
		t.Errorf("value (%+v) was not like expected (%+v)", val, expected)
	}
}
