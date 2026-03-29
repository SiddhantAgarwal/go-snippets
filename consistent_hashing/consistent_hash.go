package consistent_hashing

import (
	"crypto/sha1"
	"sort"
	"strconv"
)

type HashRing []uint32

func (h HashRing) Len() int           { return len(h) }
func (h HashRing) Less(i, j int) bool { return h[i] < h[j] }
func (h HashRing) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

type ConsistentHash struct {
	replicas int
	ring     HashRing
	nodes    map[uint32]string
}

func NewConsistentHash(replicas int) *ConsistentHash {
	return &ConsistentHash{
		replicas: replicas,
		nodes:    make(map[uint32]string),
	}
}

func (c *ConsistentHash) Hash(key string) uint32 {
	h := sha1.Sum([]byte(key))

	return uint32(h[0])<<24 | uint32(h[1])<<16 | uint32(h[2])<<8 | uint32(h[3])
}

func (c *ConsistentHash) Add(nodeName string) {
	for i := 0; i < c.replicas; i++ {
		vNodeName := nodeName + "#" + strconv.Itoa(i)
		h := c.Hash(vNodeName)
		c.ring = append(c.ring, h)
		c.nodes[h] = nodeName
	}

	sort.Sort(c.ring)
}

func (c *ConsistentHash) Get(key string) string {
	if len(c.ring) == 0 {
		return ""
	}

	h := c.Hash(key)

	idx := sort.Search(len(c.ring), func(i int) bool {
		return c.ring[i] >= h
	})

	if idx == len(c.ring) {
		idx = 0
	}

	return c.nodes[c.ring[idx]]
}
