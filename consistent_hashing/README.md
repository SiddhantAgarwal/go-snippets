# Consistent Hashing

Consistent Hashing is a specialized hashing strategy used in distributed systems to minimize data reorganization when nodes are added or removed.

Unlike traditional hashing (hash(key) mod n), where changing the number of nodes (n) forces a remapping of nearly all keys, Consistent Hashing ensures that only a small fraction of keys (1/n) need to be moved.

> [!NOTE]
> This just demonstrates how a nodes are placed on a ring using sha1 hashing as example, for production use cases
 the implementation is quite complex. 