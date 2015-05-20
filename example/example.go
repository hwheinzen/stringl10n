// example.go starts the creation of stringl10n_generated.go.
//
//go:generate stringl10n -json=example.json

// Package example does nothing.
// But 'go generate' will invoke 'stringl10n -json=example.json'.
package example
