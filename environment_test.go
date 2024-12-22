package goenv

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"unsafe"
)

func fakeEnv(key string, value string, err error) func(string) string {
	if err == nil {
		return func(s string) string {
			if s == key {
				return value
			} else {
				return ""
			}
		}
	} else {
		return func(_ string) string { return "" }
	}
}

func TestGet(t *testing.T) {
	var tests = []struct {
		name    string
		f       func(string) string
		wantVal string
		wantErr bool
	}{
		{"key exists", fakeEnv("MY_ENV_VAR", "foo", nil), "foo", false},
		{"key does not exist", fakeEnv("OTHER_ENV_VAR", "", fmt.Errorf("not found")), "", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			env := defaultEnvironment{env: test.f}
			val, err := env.Get("MY_ENV_VAR")
			if val != test.wantVal || (err != nil) != test.wantErr {
				t.Errorf("Want: (%v, %v), Got: (%v, %v)", test.wantVal, test.wantErr, val, err)
			}
		})
	}
}

func TestMustGet(t *testing.T) {
	var tests = []struct {
		name    string
		f       func(string) string
		wantVal string
		wantErr bool
	}{
		{"key exists", fakeEnv("MY_ENV_VAR", "foo", nil), "foo", false},
		{"key does not exist", fakeEnv("OTHER_ENV_VAR", "", fmt.Errorf("not found")), "", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.wantErr {
				defer func() { _ = recover() }()
			}

			env := defaultEnvironment{env: test.f}
			val := env.MustGet("MY_ENV_VAR")

			if test.wantErr {
				t.Error("Expected panic")
			}

			if val != test.wantVal {
				t.Errorf("Want: %v, Got: %v", test.wantVal, val)
			}
		})
	}
}

func TestGetOrDefault(t *testing.T) {
	var tests = []struct {
		name    string
		f       func(string) string
		def     string
		wantVal string
	}{
		{"key exists", fakeEnv("MY_ENV_VAR", "foo", nil), "bar", "foo"},
		{"key does not exist", fakeEnv("OTHER_ENV_VAR", "foo", nil), "bar", "bar"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			env := defaultEnvironment{env: test.f}
			val := env.GetOrDefault("MY_ENV_VAR", test.wantVal)
			if val != test.wantVal {
				t.Errorf("Want: (%v), Got: (%v)", test.wantVal, val)
			}
		})
	}
}

func TestNewDefaultEnvironment(t *testing.T) {
	env := NewDefaultEnvironment()

	envField := reflect.ValueOf(env).Elem().FieldByName("env")
	envFunc := reflect.NewAt(envField.Type(), unsafe.Pointer(envField.UnsafeAddr())).Elem().Interface()

	// Check if the env field points to os.Getenv
	if reflect.ValueOf(envFunc).Pointer() != reflect.ValueOf(os.Getenv).Pointer() {
		t.Errorf("Expected env to use os.Getenv, but it does not")
	}

}
