# go-map-from-lists

Google Cloud Datastore disallows using `map[string]string` on entities.

To get around this limitation, we implement a data structure that emulates a 
Go [`map`](https://blog.golang.org/go-maps-in-action) 
(see [associative array](https://en.wikipedia.org/wiki/Associative_array))
using only ["slices"](https://blog.golang.org/go-slices-usage-and-internals).