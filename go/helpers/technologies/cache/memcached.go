package main

import (
	"fmt"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
)

func NewMemcachedConnection() (*memcache.Client, error) {
	// Get the Memcached server address from environment variable
	server := os.Getenv("MEMCACHED_SERVER")
	if server == "" {
		return nil, fmt.Errorf("MEMCACHED_SERVER environment variable is not set")
	}

	// Create a new Memcached client
	client := memcache.New(server)

	// Optionally, you can set other client options here
	// client.Timeout = time.Second * 5

	return client, nil
}
