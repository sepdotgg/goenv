package goenv

import (
	"fmt"
	"os"
)

type defaultEnvironment struct {
	env func(string) string
}

// Environment is an interface that provides safe access to environment variables.
//
// It is useful for testing and for providing default values for environment variables.
// For example, you can provide your own mocked [os.Getenv] implementation:
//
//	testEnv := goenv.Environment{
//		Get: func(key string) string {
//			if key == "MY_ENV_VAR" {
//				return "my value"
//			}
//			return ""
//		},
//	}
type Environment interface {
	// Get returns the value of the environment variable with the given key, or an error if the variable is not set or is
	// an empty string.
	Get(key string) (string, error)

	// MustGet returns the value of the environment variable with the given key, or panics if the variable is not set or is
	// an empty string.
	MustGet(key string) string

	// GetOrDefault returns the value of the environment variable with the given key, or the default value if the variable is
	GetOrDefault(key string, def string) string
}

// NewDefaultEnvironment returns a new Environment that reads environment variables from [os.Getenv].
func NewDefaultEnvironment() Environment {
	return &defaultEnvironment{env: os.Getenv}
}

func (de defaultEnvironment) Get(key string) (string, error) {
	if val := de.env(key); val == "" {
		return "", fmt.Errorf("missing required environment variable: %s", key)
	} else {
		return val, nil
	}
}

func (de defaultEnvironment) MustGet(key string) string {
	if val, err := de.Get(key); err != nil {
		panic(err)
	} else {
		return val
	}
}

func (de defaultEnvironment) GetOrDefault(key string, def string) string {
	if val := de.env(key); val == "" {
		return def
	} else {
		return val
	}
}
