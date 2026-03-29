package consistent_hashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsistentHash(t *testing.T) {
	t.Parallel()

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

func TestConsistentHash_Remove(t *testing.T) {
	t.Parallel()

	ch := NewConsistentHash(3)
	ch.Add("Server-A")
	ch.Add("Server-B")
	ch.Add("Server-C")

	keys := []string{"key1", "key2", "key3", "key4", "key5", "key6", "key7"}

	nodeToKeyMapping := make(map[string][]string)
	keyToNodeMapping := make(map[string]string)

	for _, key := range keys {
		nodeToKeyMapping[ch.Get(key)] = append(nodeToKeyMapping[ch.Get(key)], key)
		keyToNodeMapping[key] = ch.Get(key)
	}

	ch.Remove("Server-B")

	for _, key := range nodeToKeyMapping["Server-B"] {
		assert.Equal(t, "Server-C", ch.Get(key))
	}
}
