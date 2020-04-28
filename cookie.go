package gsession

import (
	"sync"
)

// Set global cookie
// Concurrent security
var cookieSync sync.Map

// Define cookie structure
type cookie struct{}

// Get all the cookies in the session
func (c *cookie) GetMap() map[string]string {
	var keys []string
	j := make(map[string]string)

	// Get all keys of cookieSync
	f := func(k, v interface{}) bool {
		keys = append(keys, k.(string))
		return true
	}
	cookieSync.Range(f)

	// Assign value to j
	for i := 0; i < len(keys); i++ {
		v, _ := cookieSync.Load(keys[i])
		j[keys[i]] = v.(string)
	}
	return j
}

// Add cookies is equivalent to updating cookies
// sync.Map: Overwrite update
func (c *cookie) Add(cookie map[string]string) {
	for k, v := range cookie {
		cookieSync.Store(k, v)
	}
}

// update cookie
func (c *cookie) Update(cookie map[string]string) {
	for k, v := range cookie {
		cookieSync.Store(k, v)
	}
}

// delete one cookie
func (c *cookie) Delete(key string) {
	cookieSync.Delete(key)
}

// clear all cookie
func (c *cookie) Clear() {
	var keys []string
	f := func(k, v interface{}) bool {
		keys = append(keys, k.(string))
		return true
	}
	cookieSync.Range(f)

	for i := 0; i < len(keys); i++ {
		cookieSync.Delete(keys[i])
	}
}
