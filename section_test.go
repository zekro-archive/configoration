package configoration

import (
	"sync"
	"testing"
)

func TestGetSection(t *testing.T) {
	s := makeDefSection()

	rec1 := s.GetSection("a")
	if rec1.IsNil() {
		t.Error("recovered section was nil")
	}

	rec2 := s.GetSection("b")
	if !rec2.IsNil() {
		t.Error("non existent section was not nil")
	}
}

func TestGetValue(t *testing.T) {
	s := makeDefSection()

	rec, err := s.GetValue("a:i")
	if err != nil {
		t.Errorf("recovering returned error: %s", err.Error())
	}

	if rec.(int) != 1 {
		t.Errorf("recovered value (%+v) was not like expected (1)", rec)
	}
}

func TestGetString(t *testing.T) {
	s := makeDefSection()

	{
		rec, err := s.GetString("a:i")
		if err != nil {
			t.Errorf("recovering returned error: %s", err.Error())
		}
		if rec != "1" {
			t.Errorf(`recovered value (%+v) was not like expected ("1")`, rec)
		}
	}
	{
		rec, err := s.GetString("a:f")
		if err != nil {
			t.Errorf("recovering returned error: %s", err.Error())
		}
		if rec != "3.1415" {
			t.Errorf(`recovered value (%+v) was not like expected ("3.1415")`, rec)
		}
	}
	{
		rec, err := s.GetString("a:b")
		if err != nil {
			t.Errorf("recovering returned error: %s", err.Error())
		}
		if rec != "true" {
			t.Errorf(`recovered value (%+v) was not like expected ("true")`, rec)
		}
	}
	{
		rec, err := s.GetString("a:s")
		if err != nil {
			t.Errorf("recovering returned error: %s", err.Error())
		}
		if rec != "test123" {
			t.Errorf(`recovered value (%+v) was not like expected ("test123")`, rec)
		}
	}
	{
		_, err := s.GetString("a:none")
		if err == nil {
			t.Error("recovering returned no error")
		}
		if err != ErrNil {
			t.Error("recovering returned not the expected error ErrNil")
		}
	}
}

func makeSection(m ConfigMap) *section {
	return &section{
		mtx: sync.Mutex{},
		m:   m,
	}
}

func makeDefSection() *section {
	return makeSection(ConfigMap{
		"a": ConfigMap{
			"i": 1,
			"f": 3.1415,
			"b": true,
			"s": "test123",
		},
	})
}
