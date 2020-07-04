package configoration

// Provider provides functionalities to get
// a configuration map from a desired source.
type Provider interface {
	// GetMap reads the config values from the
	// desired resource and returns them as
	// map[string]interface{}. If errors occur
	// during value collection, the error is
	// returned.
	GetMap() (map[string]interface{}, error)
}
