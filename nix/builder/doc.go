// Package builder showcases use of bash script with nix
// expressions. Note that bash script in this example
// serves as a custom builder while building the
// derivation
//
// Usage: #1: when go is available
// $ go test
//
// Usage: #2: when nix is available
// $ nix-shell --run 'go test'
package builder
