package main

/*
	github: https://github.com/patrickmn/go-cache
*/
import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

type cacheValue struct {
	ctr       float64
	timestamp int64
}

func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(1*time.Minute, 1*time.Minute)
	c1 := cache.New(1*time.Minute, 1*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)
	c1.Set("foo", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)
	c1.Set("baz", 42, cache.NoExpiration)

	// Get the string associated with the key "foo" from the cache
	foo1, found := c.Get("foo")
	if found {
		fmt.Println(foo1)
	}

	// Since Go is statically typed, and cache values can be anything, type
	// assertion is needed when values are being passed to functions that don't
	// take arbitrary types, (i.e. interface{}). The simplest way to do this for
	// values which will only be used once--e.g. for passing to another
	// function--is:
	foo2, found := c.Get("foo")
	if found {
		fmt.Println(foo2)
		//MyFunction(foo2.(string))
	}

	// This gets tedious if the value is used several times in the same function.
	// You might do either of the following instead:
	if x, found := c.Get("foo"); found {
		foo := x.(string)
		fmt.Println(foo)
		// ...
	}
	// or
	var foo string
	if x, found := c.Get("foo"); found {
		foo = x.(string)
		fmt.Println(foo)
	}
	// ...
	// foo can then be passed around freely as a string

	// Want performance? Store pointers!
	t := time.Now()
	timestamp := t.UTC().UnixNano()
	var structValue = cacheValue{0.8, timestamp}
	c.Set("foo", &structValue, cache.DefaultExpiration)
	if x, found := c.Get("foo"); found {
		foo := &(x.(*cacheValue).ctr)
		fmt.Println(*foo)
		// ...
	}

	if x1, found := c1.Get("foo"); found {
		fmt.Println(x1)
		// ...
	}
}
