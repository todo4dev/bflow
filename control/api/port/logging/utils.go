// port/logging/utils.go
package logging

// String creates a string field
func String(key, value string) Field {
	return Field{Key: key, Value: value}
}

// Int creates an integer field
func Int(key string, value int) Field {
	return Field{Key: key, Value: value}
}

// Error creates an error field
func Error(err error) Field {
	return Field{Key: "error", Value: err.Error()}
}

// Any creates a field of any type
func Any(key string, value any) Field {
	return Field{Key: key, Value: value}
}
