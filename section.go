package configoration

import (
	"strings"
	"sync"
)

// Section provides functionalities to access
// sections and values in a Section by a key.
//
// A key can be a value or section key itself
// like "webserver" or it can span over sections
// like "general:webserver". In this case, the
// value after the last delimiter is selected
// section or value.
type Section interface {
	// GetSection returns a section by key.
	// If the desired section is not existent,
	// the returned value will be nil.
	GetSection(key string) Section

	// GetValue returns an interface value by
	// key. If the desired value could not be
	// found, nil and ErrNil is returned.
	GetValue(key string) (interface{}, error)

	// GetString is shorthand for GetValue and
	// returns a string or an ErrNil if the
	// key was not found.
	//
	// If the value selected is not a string,
	// ErrInvalidType will be returned.
	GetString(key string) (string, error)

	// GetInt is shorthand for GetValue and
	// returns an int or an ErrNil if the
	// key was not found.
	//
	// If the value selected is not an int,
	// ErrInvalidType will be returned.
	GetInt(key string) (int, error)

	// GetBool is shorthand for GetValue and
	// returns a bool or an ErrNil if the
	// key was not found.
	//
	// If the value selected is not a bool,
	// ErrInvalidType will be returned.
	GetBool(key string) (bool, error)

	// GetFloat64 is shorthand for GetValue and
	// returns a float64 or an ErrNil if the
	// key was not found.
	//
	// If the value selected is not a float64,
	// ErrInvalidType will be returned.
	GetFloat64(key string) (float64, error)
}

// section is the default implementation of
// the Section interface.
type section struct {
	mtx sync.Mutex
	m   ConfigMap
}

func (s *section) GetSection(key string) Section {
	for _, nextSelector := range splitSections(key) {
		if s == nil {
			return nil
		}
		s = s.getSection(nextSelector)
	}
	return s
}

func (s *section) GetValue(key string) (interface{}, error) {
	if s == nil {
		return nil, ErrNil
	}

	selectors := splitSections(key)
	lenSelectors := len(selectors)
	if lenSelectors > 1 {
		for i := 0; i < lenSelectors-1; i++ {
			s = s.getSection(selectors[i])
			if s == nil {
				return nil, ErrNil
			}
		}
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	v, ok := s.m[selectors[lenSelectors-1]]
	if !ok {
		return nil, ErrNil
	}

	return v, nil
}

func (s *section) GetString(key string) (string, error) {
	v, err := s.GetValue(key)
	if err != nil {
		return "", err
	}

	vt, ok := v.(string)
	if !ok {
		return "", ErrInvalidType
	}

	return vt, nil
}

func (s *section) GetInt(key string) (int, error) {
	v, err := s.GetValue(key)
	if err != nil {
		return 0, err
	}

	vt, ok := v.(int)
	if !ok {
		return 0, ErrInvalidType
	}

	return vt, nil
}

func (s *section) GetBool(key string) (bool, error) {
	v, err := s.GetValue(key)
	if err != nil {
		return false, err
	}

	vt, ok := v.(bool)
	if !ok {
		return false, ErrInvalidType
	}

	return vt, nil
}

func (s *section) GetFloat64(key string) (float64, error) {
	v, err := s.GetValue(key)
	if err != nil {
		return 0, err
	}

	vt, ok := v.(float64)
	if !ok {
		return 0, ErrInvalidType
	}

	return vt, nil
}

// getSection returns the desired section
// or nil, if not found.
func (s *section) getSection(sec string) *section {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	v := s.m[sec]
	vc, ok := v.(ConfigMap)
	if !ok {
		return nil
	}

	return &section{
		mtx: sync.Mutex{},
		m:   vc,
	}
}

// splitSections splits the passed key by
// the Delimiter and returns the resulting
// array of strings.
func splitSections(key string) []string {
	return strings.Split(key, Delimiter)
}
