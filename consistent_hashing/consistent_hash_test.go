package consistent_hashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsistentHash(t *testing.T) {
	ch := NewConsistentHash(3)
	ch.Add("Server-A")
	ch.Add("Server-B")
	ch.Add("Server-C")

	keys := []string{"key1", "key2", "key3", "key4", "key5", "key6", "key7"}

	keyMapping := make(map[string]string)
	for _, key := range keys {
		keyMapping[key] = ch.Get(key)
	}

	for _, key := range keys {
		cKey := ch.Get(key)
		assert.Equal(t, keyMapping[key], cKey)
	}
}
