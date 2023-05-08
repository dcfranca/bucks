# bucks
Simple token bucket (rate limiter) implementation

# Installation

```
go get -u github.com/dcfranca/bucks
```

# Usage

Create a new bucket share token, specifying its capacity and refill rate
```
tb := NewTokenBucket(30, 3)
```
This will initialize it with 30 tokens and refill with 3 more every second

To validate whether a request should be authorized you can call the TakeToken method with the number of tokens required for the request:
```
if !tb.TakeToken(1) {
    fmt.Println("Limit exceeded")
}
```

A token bucket can have a negative number of tokens, it is allow a last request, the request is only denied
when the current number of tokens is already less or equal zero
