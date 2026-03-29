# Token Bucket

The Token Bucket algorithm is a flexible rate-limiting mechanism that allows for bursty traffic while maintaining a 
strict average rate over time. Unlike the Leaky Bucket, it focuses on the availability of "tokens" to authorize requests.

> [!TIP]
> This shows a token bucket on a node, to implement it in a true distributed environment use something like redis with
lua script to achieve true distributed rate limiting. 
