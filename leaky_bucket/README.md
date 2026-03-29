# Leaky Bucket

The Leaky Bucket algorithm is a traffic-shaping mechanism used to control data rate and smooth out bursty traffic. It ensures a consistent, 
predictable output rate regardless of the input speed.

> [!TIP]
> This show a leaky bucket on a node, to implement it in a true distributed environment use something like redis with 
lua script to achieve true distributed rate limiting. 
