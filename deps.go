// Package deps is a shared module that the snyk-demo-go repos depend on.
package deps

import "fmt"

// Version identifies this module.
func Version() string { return "deps-01 v0.1.0" }

// Greet is called by every dependent library.
func Greet(caller string) string { return fmt.Sprintf("%s uses deps-01", caller) }
