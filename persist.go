// Package persist provides an abstraction
// over various means of storing key/value
// pairs.
//
// The primary driver here is speed
package persist

type Store interface {
	Get(key string) []byte
	GetWithMeta(key string) ([]byte, map[string]string)
	Set(key string, item []byte)
	SetWithMeta(key string, item []byte, meta map[string]string)
}
